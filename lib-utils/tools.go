package lib_utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io"
	"log"
	"reflect"
	"strings"
	"time"
	"unicode"
)

// GenerateBytesUUID returns a UUID based on RFC 4122 returning the generated bytes
func GenerateBytesUUID() []byte {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		panic(fmt.Sprintf("Error generating UUID: %s", err))
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return uuid
}

// GenerateUUID returns a UUID based on RFC 4122
func GenerateUUID() string {
	uuid := GenerateBytesUUID()
	return idBytesToStr(uuid)
}

func idBytesToStr(id []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
}

// keysPrealloc returns array with map key
func keysPrealloc(m map[string]string) []string {
	k := make([]string, len(m))
	var i uint64
	for key := range m {
		k[i] = key
		i++
	}

	return k
}

// ToChaincodeArgs converts string args to []byte args
func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// NormalizeString
func NormalizeString(text string) string {
	lower := strings.ToLower(text)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), runes.Remove(runes.In(unicode.Space)), norm.NFC) // Mn: nonspacing marks
	result, _, err := transform.String(t, lower)
	if err != nil {
		return lower
	}

	return result
}

// ConcatenateBytes is useful for combining multiple arrays of bytes, especially for
// signatures or digests over multiple fields
func ConcatenateBytes(data ...[]byte) []byte {
	finalLength := 0
	for _, slice := range data {
		finalLength += len(slice)
	}
	result := make([]byte, finalLength)
	last := 0
	for _, slice := range data {
		for i := range slice {
			result[i+last] = slice[i]
		}
		last += len(slice)
	}
	return result
}

func DeepCopy(v interface{}) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return nil, err
	}
	return vptr.Elem().Interface(), err
}

func GetTxTimestampRFC3339(stub shim.ChaincodeStubInterface) (string, error) {
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return "", err
	}
	tm := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	return tm.Format(time.RFC3339), nil
}

func ParseRFC3339toTime(tm string) (time.Time, error) {
	t1, err := time.Parse(time.RFC3339, tm)
	if err != nil {
		return time.Time{}, err
	}
	return t1, nil
}

func Contains(arr []string, elem string) bool {
	for _, e := range arr {
		if elem == e {
			return true
		}
	}
	return false
}

func GetFirstElem(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

type BeforeTransactionUnmarshalResponse struct {
	Id       string `json:"id"` // user id or did
	Function string `json:"function"`
}

// FunctionCompare compares one function name with another,
// the input parameter can be in the form "org.tecnomatica.participant:GetParticipant"
func FunctionCompare(f1, f2 string) (bool, error) {
	if s := strings.Split(f1, ":"); len(s) == 1 {
		log.Printf("1. FunctionCompare -- > %v - %v", f1, s[0])
		return s[0] == f2, nil
	} else if len(s) == 2 {
		log.Printf("2. FunctionCompare -- > %v", s[1])
		return s[1] == f2, nil
	}
	log.Printf("3. FunctionCompare -- > %v - %v", f1, f2)
	return false, fmt.Errorf("invalid transaction for function")
}

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func SliceToMap(slice []string, dMap map[string]string) {
	for _, data := range slice {
		if _, ok := dMap[data]; !ok {
			dMap[data] = ""
		}
	}
}
func MapToSlice(dMap map[string]string) []string {
	// Convert map to slice of keys.
	var slice []string
	for key, _ := range dMap {
		slice = append(slice, key)
	}
	return slice
}

// CreateIndex create search index for ledger
//
// Arguments:
//		0: indexName -
//		1: attributes -
// Returns:
//		0: error
func CreateIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	log.Printf("[start][CreateIndex]")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of object.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	if err = stub.PutState(indexKey, value); err != nil {
		return err
	}

	log.Printf("[end][CreateIndex]")
	return nil
}

// DeleteIndex remove search index for ledger
//
// Arguments:
//		0: indexName -
//		1: attributes -
// Returns:
//		0: error
func DeleteIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string, validateDelState bool) error {
	log.Printf("[start][DeleteIndex]")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Delete index by key
	if err = stub.DelState(indexKey); err != nil && validateDelState {
		return err
	}

	log.Printf("[end][DeleteIndex]")
	return nil
}

type smallIssuer struct {
	CertPem string `json:"certPem"` // cert PEM active
}

// validateArgsLen ensures `args` has at least size `length`.
func validateArgsLen(args []string, length int) error {
	if len(args) < length {
		return fmt.Errorf("invalid arguments length %d expected: %d", len(args), length)
	}
	return nil
}
