package libutils

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/attrmgr"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	jose "gopkg.in/square/go-jose.v2"
	"log"
	"strings"
	"time"
)

// GetX509CertFromPemByte use to validate cert
func GetX509CertFromPemByte(certPem []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil, fmt.Errorf("failed to decode certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err.Error())
	} else if cert == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return cert, nil
}

// GetX509CertFromPem use to validate cert
func GetX509CertFromPem(certPemBase64 string) (*x509.Certificate, error) {
	// decode certificate
	certByte, err := base64.StdEncoding.DecodeString(certPemBase64)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(certByte)
	if block == nil {
		return nil, fmt.Errorf("failed to decode certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err.Error())
	} else if cert == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return cert, nil
}

func CompareCertsPemBase64(cert1, cert2 string) bool {
	// decode certificate
	certByte1, err := base64.StdEncoding.DecodeString(cert1)
	if err != nil {
		return false
	}

	certByte2, err := base64.StdEncoding.DecodeString(cert2)
	if err != nil {
		return false
	}

	return bytes.Compare(certByte1, certByte2) == 0
}

func GetAttrsNonStandardCert(certPem []byte) (map[string]string, error) {
	cert, err := GetX509CertFromPemByte(certPem)
	if err != nil {
		return nil, err
	}
	attrs, err := attrmgr.New().GetAttributesFromCert(cert)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return attrs.Attrs, nil
}

// CheckAttr return value for a certificate attribute or check the value of an attribute
func CheckAttr(name, value string, attrs *attrmgr.Attributes) (bool, error) {
	v, found, err := attrs.Value(name)
	if err != nil {
		return false, err
	}
	if !found {
		return found, fmt.Errorf("does not contain attribute '%s'", name)
	} else if value == "" && found {
		return found, nil
	} else if v == value && found {
		return found, nil
	}
	return false, fmt.Errorf("incorrect value for '%s'; expected '%s' but found '%s'", name, value, v)
}

func AssertAdmin(ctx contractapi.TransactionContextInterface) error {
	found, err := cid.HasOUValue(ctx.GetStub(), "admin")
	if err != nil {
		return fmt.Errorf("admin identity: %v", err)
	}
	if !found {
		return fmt.Errorf("is not operator admin identity")
	}
	return nil
}

// VerifyIssuedByRootCert Verify cert issued by Root Certificate.
func VerifyIssuedByRootCert(certRoot []byte, certIssue []byte) error {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(certRoot)
	if !ok {
		return fmt.Errorf("failed to parse root certificate")
	}
	cert, err := GetX509CertFromPemByte(certIssue)
	if err != nil {
		return err
	}

	opts := x509.VerifyOptions{
		//DNSName: "cupet.cu",
		Roots:     roots,
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}

	if _, err := cert.Verify(opts); err != nil {
		return fmt.Errorf("failed to verify certificate: %v", err.Error())
	}

	return nil
}

// CheckSignatureFrom verifies that the signature on certIssue is a valid signature from Root Cert.
func CheckSignatureFrom(certRoot []byte, certIssue []byte) error {
	certR, err := GetX509CertFromPemByte(certRoot)
	if err != nil {
		return err
	}

	certI, err := GetX509CertFromPemByte(certIssue)
	if err != nil {
		return err
	}

	err = certI.CheckSignatureFrom(certR)
	if err != nil {
		return err
	}
	return nil
}

func HasExpired(cert *x509.Certificate) error {
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return x509.CertificateInvalidError{
			Cert:   cert,
			Reason: x509.Expired,
			Detail: fmt.Sprintf("current time %s is before %s", now.Format(time.RFC3339), cert.NotBefore.Format(time.RFC3339)),
		}
	} else if now.After(cert.NotAfter) {
		return x509.CertificateInvalidError{
			Cert:   cert,
			Reason: x509.Expired,
			Detail: fmt.Sprintf("current time %s is after %s", now.Format(time.RFC3339), cert.NotAfter.Format(time.RFC3339)),
		}
	}

	return nil
}

// GetDateCertificate returns the issue and expiration date of a certificate
func GetDateCertificate(cert *x509.Certificate) map[string]string {
	issuedTime := cert.NotBefore.Format(time.RFC3339)
	expiresTime := cert.NotAfter.Format(time.RFC3339)

	res := map[string]string{"issuedTime": issuedTime, "expiresTime": expiresTime}

	return res
}

func DidFormat(hashPubKey string) string {
	return fmt.Sprintf("did:%s", hashPubKey)
}

func checkSignature(payload string, key string) (map[string]interface{}, error) {
	params := make(map[string]interface{})

	return params, nil
}

func parseMessage(message string) (*jose.JSONWebSignature, error) {
	jwsSignature, err := jose.ParseSigned(message)
	if err != nil {
		return nil, errors.New(ErrorParseJWS)
	}
	return jwsSignature, nil
}

func parsePublicKeyX509(publicKey string) (interface{}, error) {
	base64Data := []byte(publicKey)

	d := make([]byte, base64.StdEncoding.DecodedLen(len(base64Data)))
	n, err := base64.StdEncoding.Decode(d, base64Data)
	if err != nil {
		return nil, errors.New(ErrorBase64)
	}
	d = d[:n]

	publicKeyImported, err := x509.ParsePKIXPublicKey(d)
	if err != nil {
		return nil, errors.New(ErrorParseX509)
	}
	return publicKeyImported, nil
}

// CertificateAlreadyExists check if the certificate already exists in the world-state
//
// Arguments:
//		0: certPemBase64 - certificate to validate
//		1: indexName -
//		2: attributes -
// Returns:
//		0: bool
//		1: error
func CertificateAlreadyExists(ctx contractapi.TransactionContextInterface, certPemBase64 string, indexName string, attributes []string) (bool, error) {
	log.Printf("[inside][certificateAlreadyExists]")

	issuesResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(indexName, attributes)
	if err != nil {
		return false, err
	}
	defer issuesResultsIterator.Close()

	for issuesResultsIterator.HasNext() {
		responseRange, err := issuesResultsIterator.Next()
		if responseRange == nil {
			return false, err
		}

		var crt smallIssuer
		err = json.Unmarshal(responseRange.Value, &crt)
		if err != nil {
			return false, err
		}

		if CompareCertsPemBase64(crt.CertPem, certPemBase64) {
			return true, nil
		}
	}
	return false, nil
}

func parseKey(publicKey string) string {

	begin := "-----BEGIN PUBLIC KEY-----"
	end := "-----END PUBLIC KEY-----"

	// Replace all pairs.
	noBegin := strings.Split(publicKey, begin)
	parsed := strings.Split(noBegin[1], end)
	return parsed[0]
}

func VerifySignature(message string, key string) ([]byte, error) {
	msg, err := parseMessage(message)
	if err != nil {
		return nil, err
	}
	pbkey, err := parsePublicKeyX509(key)
	if err != nil {
		return nil, err
	}
	result, err := jose.JSONWebSignature.Verify(*msg, pbkey)
	if err != nil {
		return nil, errors.New(ErrorVerifying)
	}
	return result, nil
}
