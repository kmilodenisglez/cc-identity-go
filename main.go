package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/ic-matcom/cc-identity-go/contracts/identity"
	"github.com/ic-matcom/cc-identity-go/hooks"
	modelapi "github.com/ic-matcom/model-identity-go/model"
	"log"
)

func main() {
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

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode %v", err)
	}

	//privateCertByte, err := testcerts.Certificates[2].PrivateBytes()
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
	//parsedPKBytes, _ := x509.MarshalPKIXPublicKey(parsedKey)
	//
	////fmt.Printf("%v", string(parsedPKBytes))
	//
	//
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
