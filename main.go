package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/kmilodenisglez/cc-identity-go/contracts/identity"
	"github.com/kmilodenisglez/cc-identity-go/hooks"
	lus "github.com/kmilodenisglez/cc-identity-go/lib-utils"
	modelapi "github.com/kmilodenisglez/model-identity-go/model"
	"gopkg.in/square/go-jose.v2"
	"io/ioutil"
	"log"
	"os"
)

type serverConfig struct {
	CCID    string
	Address string
}

func main() {
	// See chaincode.env
	config := serverConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	// *** This smart-constract later becomes a chaincode  ***
	contractIdentity := new(identity.ContractIdentity)
	contractIdentity.Name = modelapi.ContractNameIdentity
	contractIdentity.Info.Version = "0.2.1"
	contractIdentity.UnknownTransaction = hooks.UnknownTransactionHandler // Is only called if a request is made to invoke a transaction not defined in the smart contract
	chaincode, err := contractapi.NewChaincode(contractIdentity)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "IdentityChaincode"
	chaincode.Info.Version = "0.0.2"
	chaincode.DefaultContract = contractIdentity.GetName() // default contract

	server := &shim.ChaincodeServer{
		CCID:     config.CCID,
		Address:  config.Address,
		CC:       chaincode,
		TLSProps: getTLSProperties(),
	}

	if err := server.Start(); err != nil {
		log.Panicf("Error starting %s %s %v", chaincode.Info.Title, chaincode.Info.Version, err)
	}

	//_, err = testcerts.Certificates[2].PrivateBytes()
	//certByte, err := testcerts.Certificates[2].CertBytes()
	//if err != nil {
	//	fmt.Printf("1-->> %v", err)
	//}
	//cert, err := lus.GetX509CertFromPemByte(certByte)
	////public := cert.PublicKey
	//
	//parsedKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	fmt.Errorf("wanted an ECDSA public key but found: %#v", parsedKey)
	//}
	//
	//
	//parsedPKBytes, err := x509.MarshalPKIXPublicKey(parsedKey)
	//if err != nil {
	//	panic(err)
	//}
	//
	////key, err := x509.ParsePKIXPublicKey(parsedPKBytes)
	////if err != nil {
	////	panic(err)
	////}
	//
	//
	//fmt.Printf("%v : %v", string(parsedPKBytes), base64.StdEncoding.EncodeToString(parsedPKBytes))
	//lop := modelapi.RoleCreateRequest{
	//	Name:              "name 1",
	//	ContractFunctions: nil,
	//}
	//
	//lopByte,_ := json.Marshal(lop)
	//
	//signingKey, err := LoadPrivateKey(privateCertByte)
	////fmt.Printf("signingKey-->  %v", signingKey)
	//
	//alg := jose.SignatureAlgorithm("ES256")
	//signer, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: signingKey}, nil)
	//if err != nil {
	//	fmt.Printf("err-->  %v", err)
	//}
	//
	//
	//obj, err := signer.Sign(lopByte)
	//if err != nil {
	//	fmt.Printf("err-->  %v", err)
	//}
	//
	//res := obj.FullSerialize()
	//
	//fmt.Println("signed: ", res)
	//verify(`{"payload":"eyJuYW1lIjoibmFtZSAxIn0","protected":"eyJhbGciOiJFUzI1NiJ9","signature":"lMFcByiTjr4kd0C9qWVPCdpLCcsgpBuoUSh73mdm_xstKNxoKWXwaO0GkBF_55T5bJSaaSphUfZhu-9XmNABQw"}`, parsedPKBytes)

}

func getTLSProperties() shim.TLSProperties {
	// Check if chaincode is TLS enabled
	tlsDisabledStr := lus.GetEnvOrDefault("CHAINCODE_TLS_DISABLED", "true")
	key := lus.GetEnvOrDefault("CHAINCODE_TLS_KEY", "")
	cert := lus.GetEnvOrDefault("CHAINCODE_TLS_CERT", "")
	clientCACert := lus.GetEnvOrDefault("CHAINCODE_CLIENT_CA_CERT", "")

	// convert tlsDisabledStr to boolean
	tlsDisabled := lus.GetBoolOrDefault(tlsDisabledStr, false)
	var keyBytes, certBytes, clientCACertBytes []byte
	var err error

	if !tlsDisabled {
		keyBytes, err = ioutil.ReadFile(key)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
		certBytes, err = ioutil.ReadFile(cert)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
	}
	// Did not request for the peer cert verification
	if clientCACert != "" {
		clientCACertBytes, err = ioutil.ReadFile(clientCACert)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
	}

	return shim.TLSProperties{
		Disabled:      tlsDisabled,
		Key:           keyBytes,
		Cert:          certBytes,
		ClientCACerts: clientCACertBytes,
	}
}

func verify(input string, pubKey []byte) {
	verificationKey, _ := LoadPublicKey(pubKey)

	obj, _ := jose.ParseSigned(input)

	plaintext, _ := obj.Verify(verificationKey)

	//fmt.Println("verify: ", plaintext)

	var lop = modelapi.RoleCreateRequest{}
	_ = json.Unmarshal(plaintext, &lop)

	fmt.Println("verify: ", lop)
}

func LoadJSONWebKey(json []byte, pub bool) (*jose.JSONWebKey, error) {
	var jwk jose.JSONWebKey
	err := jwk.UnmarshalJSON(json)
	if err != nil {
		return nil, err
	}
	if !jwk.Valid() {
		return nil, errors.New("invalid JWK key")
	}
	if jwk.IsPublic() != pub {
		return nil, errors.New("priv/pub JWK key mismatch")
	}
	return &jwk, nil
}

// LoadPublicKey loads a public key from PEM/DER/JWK-encoded data.
func LoadPublicKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	// Try to load SubjectPublicKeyInfo
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
		return pub, nil
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	jwk, err2 := LoadJSONWebKey(data, true)
	if err2 == nil {
		return jwk, nil
	}

	return nil, errors.New("parse error, invalid public key")
}

// LoadPrivateKey loads a private key from PEM/DER/JWK-encoded data.
func LoadPrivateKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	var priv interface{}
	priv, err0 := x509.ParsePKCS1PrivateKey(input)
	if err0 == nil {
		return priv, nil
	}

	priv, err1 := x509.ParsePKCS8PrivateKey(input)
	if err1 == nil {
		return priv, nil
	}

	priv, err2 := x509.ParseECPrivateKey(input)
	if err2 == nil {
		return priv, nil
	}

	jwk, err3 := LoadJSONWebKey(input, false)
	if err3 == nil {
		return jwk, nil
	}

	return nil, errors.New("parse error, invalid private key")
}
