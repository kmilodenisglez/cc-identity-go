package identity

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-identity-go/model"
	modeltools "github.com/ic-matcom/model-identity-go/tools"
	"log"
)

// CreateIssuer create in the ledger the issuer's certificate with its attributes
//
// Arguments:
//		0: IssuerCreateRequest
// Returns:
//		0: Issuer
//		1: error
func (ci *ContractIdentity) CreateIssuer(ctx contractapi.TransactionContextInterface, issuerRequest model.IssuerCreateRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][CreateIssuer]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	exist, err := lus.CertificateAlreadyExists(ctx, issuerRequest.CertPem, IssuerDocType, []string{})
	if err != nil {
		return nil, err
	} else if exist {
		return nil, fmt.Errorf("an issuer with the same certificate already exists")
	}

	issuerID := lus.GenerateUUIDStr()

	// If the new issuer is to be the default, we must check if there is
	// another issuer registered and marked "by default"
	if err := ci.issuerMarkAsDefault(ctx, issuerID, &issuerRequest); err != nil {
		return nil, err
	}

	// decode certificate
	certX509, err := lus.GetX509CertFromPem(issuerRequest.CertPem)
	if err != nil {
		return nil, err
	}
	// we insert commonName if the Name field is empty
	commonName := certX509.Subject.CommonName
	if issuerRequest.Name == "" {
		issuerRequest.Name = commonName
	}

	attrs := modeltools.GetAttrsCert(certX509)
	// get dates
	dateCert := lus.GetDateCertificate(certX509)
	// begin: get publicKey
	parsedKey, ok := certX509.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("wanted an ECDSA public key but found: %#v", parsedKey)
	}
	parsedPKBytes, err := x509.MarshalPKIXPublicKey(parsedKey)
	if err != nil {
		panic(err)
	}
	publicKey := base64.StdEncoding.EncodeToString(parsedPKBytes)
	// end: publicKey

	// Create Issuer
	issuer := &model.Issuer{
		DocType:     IssuerDocType,
		ID:          issuerID,
		Name:        issuerRequest.Name,
		CertPem:     issuerRequest.CertPem,
		Attrs:       attrs,
		AttrsExtras: make(map[string]string),
		IssuedTime:  dateCert["issuedTime"],
		ExpiresTime: dateCert["expiresTime"],
		ByDefault:   issuerRequest.ByDefault,
	}

	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerID})
	if err != nil {
		return nil, err
	}

	// JSON encoding of issuer
	issuerJE, _ := json.Marshal(issuer)
	if err := ctx.GetStub().PutState(key, issuerJE); err != nil {
		return nil, fmt.Errorf("issuer %s could not be created: %v", issuerRequest.Name, err)
	}

	return &model.IssuerQueryResponse{
		ID:          issuer.ID,
		Name:        issuer.Name,
		PublicKey:   publicKey,
		Attrs:       issuer.Attrs,
		AttrsExtras: issuer.AttrsExtras,
		IssuedTime:  issuer.IssuedTime,
		ExpiresTime: issuer.ExpiresTime,
		Active:      issuer.Active,
		ByDefault:   issuer.ByDefault,
	}, nil
}

// IssuerUpdateRequest
type IssuerUpdateRequest struct {
	ID      string `json:"id"`
	CertPem string `json:"certPem"` // b64 certificate PEM active
}

// RenewIssuer update issuer certificate in the ledger
//
// Arguments:
//		0: IssuerUpdateRequest
// Returns:
//		0: Issuer
//		1: error
func (ci *ContractIdentity) RenewIssuer(ctx contractapi.TransactionContextInterface, issuerRequest IssuerUpdateRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][UpdateIssuer]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// get issuer
	issuerToUpdate, err := ci.GetIssuer(ctx, model.GetRequest{ID: issuerRequest.ID})
	if err != nil {
		return nil, fmt.Errorf("failed to get User identity: %v", err)
	}
	if issuerToUpdate == nil {
		return nil, fmt.Errorf("%s does not exist", issuerToUpdate.ID)
	}

	// decode certificate
	certX509, err := lus.GetX509CertFromPem(issuerRequest.CertPem)
	if err != nil {
		return nil, err
	}
	// we insert commonName if the Name field is empty
	commonName := certX509.Subject.CommonName

	attrs := modeltools.GetAttrsCert(certX509)
	// get dates
	dateCert := lus.GetDateCertificate(certX509)

	// Create Issuer
	issuerToUpdate.Name = commonName
	issuerToUpdate.PublicKey = issuerRequest.CertPem
	issuerToUpdate.Attrs = model.Attrs(attrs)
	issuerToUpdate.IssuedTime = dateCert["issuedTime"]
	issuerToUpdate.ExpiresTime = dateCert["expiresTime"]

	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerToUpdate.ID})
	if err != nil {
		return nil, err
	}

	// JSON encoding of issuer
	issuerJE, _ := json.Marshal(issuerToUpdate)
	if err := ctx.GetStub().PutState(key, issuerJE); err != nil {
		return nil, fmt.Errorf("issuer %s could not be created: %v", commonName, err)
	}
	return issuerToUpdate, nil
}

// GetIssuer get an issuer from the ledger
//
// Arguments:
//		0: GetRequest
// Returns:
//		0: Issuer
//		1: error
func (ci *ContractIdentity) GetIssuer(ctx contractapi.TransactionContextInterface, request model.GetRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][GetIssuer]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{request.ID})
	if err != nil {
		return nil, fmt.Errorf("error happened creating key for: %v", err)
	} else if key == "" {
		return nil, fmt.Errorf("no state found for %s", key)
	}
	issuer, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %v", err)
	} else if issuer == nil {
		return nil, fmt.Errorf("no state found for %s", key)
	}
	var issuerJD model.Issuer
	err = json.Unmarshal(issuer, &issuerJD)
	if err != nil {
		return nil, err
	}

	return &model.IssuerQueryResponse{
		ID:          issuerJD.ID,
		Name:        issuerJD.Name,
		PublicKey:   issuerJD.CertPem,
		Attrs:       issuerJD.Attrs,
		AttrsExtras: issuerJD.AttrsExtras,
		IssuedTime:  issuerJD.IssuedTime,
		ExpiresTime: issuerJD.ExpiresTime,
		Active:      issuerJD.Active,
		ByDefault:   issuerJD.ByDefault,
	}, nil
}

// DeleteIssuer delete an issuer from the ledger
//
// Arguments:
//		0: GetRequest
// Returns:
//		0: Issuer
//		1: error
func (ci *ContractIdentity) DeleteIssuer(ctx contractapi.TransactionContextInterface, issuerRequest model.GetRequest) error {
	log.Printf("[%s][DeleteIssuer]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerRequest.ID})
	if err != nil {
		return fmt.Errorf("error happened creating key: %v", err)
	} else if key == "" {
		return fmt.Errorf("no state found for %s", key)
	}
	err = ctx.GetStub().DelState(key)
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}

	// validateDelState = false
	// We are not interested in validating error, it may be the case that the issuer does not have an index
	if err = lus.DeleteIndex(ctx.GetStub(), ObjectTypeIssuerByDefault, []string{issuerRequest.ID}, false); err != nil {
		return err
	}

	return nil
}

// GetIssuers get all issuer
//
// Arguments:
//		0: none
// Returns:
//		0: []Issuer
//		1: error
func (ci *ContractIdentity) GetIssuers(ctx contractapi.TransactionContextInterface) ([]model.IssuerQueryResponse, error) {
	log.Printf("[%s][GetIssuers]", ctx.GetStub().GetChannelID())

	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IssuerDocType, []string{})
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var items []model.IssuerQueryResponse
	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}

		var item model.IssuerQueryResponse
		err = json.Unmarshal(responseRange.Value, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil

	//return  (*model.IssuerQueryResponse)(unsafe.Pointer(&items)), nil
}

// GetIssuerHistory returns the chain of custody for a issuer since issuance
//
// Arguments:
//		0: model.GetRequest
// Returns:
//		0: []model.IssuerHistoryQueryResponse
//		1: error
func (ci *ContractIdentity) GetIssuerHistory(ctx contractapi.TransactionContextInterface, issuerRequest model.GetRequest) ([]model.IssuerHistoryQueryResponse, error) {
	log.Printf("GetIssuerHistory: ID %v", issuerRequest.ID)

	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerRequest.ID})
	if err != nil {
		return nil, fmt.Errorf("error happened creating composite key for issuer: %v", err)
	} else if compositeKeyID == "" {
		return nil, fmt.Errorf("no state found for %s", compositeKeyID)
	}
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(compositeKeyID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []model.IssuerHistoryQueryResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		issuer := &model.IssuerQueryResponse{}
		issuer.ID = issuerRequest.ID
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &issuer)
			if err != nil {
				return nil, err
			}
		}

		timestamp := modeltools.GetTimestampRFC3339(response.Timestamp)

		record := model.IssuerHistoryQueryResponse{
			TxID:     response.TxId,
			Time:     timestamp,
			Record:   issuer,
			IsDelete: response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

// issuerMarkAsDefault modifies issuer  "byDefault" status to true only if request ByDefault field
// is true or if there are no issuers
func (ci *ContractIdentity) issuerMarkAsDefault(ctx contractapi.TransactionContextInterface, issuerID string, issuerRequest *model.IssuerCreateRequest) error {
	log.Printf("[inside][issuerMarkAsDefault]")

	issuerResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeIssuerByDefault, []string{})
	if err != nil {
		return err
	}
	defer issuerResultsIterator.Close()
	thereIsAtLeastOne := issuerResultsIterator.HasNext()

	if thereIsAtLeastOne && issuerRequest.ByDefault {
		responseRange, err := issuerResultsIterator.Next()
		if responseRange == nil {
			return err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return err
		}

		if len(compositeKeyParts) == 1 {
			returnedIssuerID := compositeKeyParts[0]
			response, err := ci.GetIssuer(ctx, model.GetRequest{ID: returnedIssuerID})
			if err != nil {
				return err
			}
			response.ByDefault = false

			// struct to JSON
			responseJSON, err := json.Marshal(response)
			if err != nil {
				return fmt.Errorf("error happened marshalling the issuer: %v", err)
			}

			ledgerKey, _ := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{returnedIssuerID})
			// Put issuer to StateDB
			err = ctx.GetStub().PutState(ledgerKey, responseJSON)
			if err != nil {
				return fmt.Errorf("error happened updating the issuer: %v", err)
			}

			// remove index ObjectTypeIssuerByDefault
			err = ctx.GetStub().DelState(responseRange.Key)
			if err != nil {
				return fmt.Errorf("error happened delete the  old issuer index : %v", err)
			}
		}
	} else if thereIsAtLeastOne && !issuerRequest.ByDefault {
		return nil
	}

	if !thereIsAtLeastOne && !issuerRequest.ByDefault {
		issuerRequest.ByDefault = true
	}

	if err := lus.CreateIndex(ctx.GetStub(), ObjectTypeIssuerByDefault, []string{issuerID}); err != nil {
		return fmt.Errorf("could not create the byDefault index for the issuer %v: %v", issuerID, err)
	}

	return nil
}
