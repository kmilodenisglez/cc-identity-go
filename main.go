package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/ic-matcom/cc-identity-go/contracts/identity"
	"github.com/ic-matcom/cc-identity-go/hooks"
	model "github.com/ic-matcom/model-traceability-go"
	"log"
)

func main() {
	// *** This smart-constract later becomes a chaincode  ***
	contractIdentity := new(identity.ContractIdentity)
	contractIdentity.Name = model.ContractNameIdentity
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
}
