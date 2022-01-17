package identity

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-identity-go/model"
	modeltools "github.com/ic-matcom/model-identity-go/tools"
	"log"
)

// InitLedger adds a base set of data to the ledger
func (ci *ContractIdentity) InitLedger(ctx contractapi.TransactionContextInterface) error {
	log.Printf("[%s][InitLedger]", ctx.GetStub().GetChannelID())
	// check if client-node is connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return fmt.Errorf(err.Error())
	}
	accessIdentity := model.AccessCreateRequest{
		ContractName:      ci.Name,                        // contract name
		ContractFunctions: modeltools.GetTransactions(ci), // functions name
	}

	// create identity access
	_, err := ci.CreateAccess(ctx, accessIdentity)
	if err != nil {
		return err
	}

	return nil
}

// TODO: missing validates ca cert - participant cert
func (ci *ContractIdentity) CreateParticipant(ctx contractapi.TransactionContextInterface, request model.ParticipantCreateRequest) (*model.ParticipantResponse, error) {
	log.Printf("[%s][CreateParticipant]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Get MSP ID of the client
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf(lus.ErrorGetMSPID, err)
	}

	did, err := modeltools.CreateDid(request.PublicKey)
	if err != nil {
		return nil, err
	}

	exist, err := ci.ParticipantExits(ctx, model.ParticipantGetRequest{Did: did})
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf(lus.ErrorIdentityExists, did)
	}

	exist, err = lus.CertificateAlreadyExists(ctx, request.CertPem, ParticipantDocType, []string{})
	if err != nil {
		return nil, err
	} else if exist {
		return nil, fmt.Errorf("an identity with the same certificate already exists")
	}

	var issuedTime, expiresTime, attrs = "", "", model.Attrs{}

	if request.CertPem != "" {
		// validate cert
		certX509, err := lus.GetX509CertFromPem(request.CertPem)
		if err != nil {
			return nil, err
		}
		err = lus.HasExpired(certX509)
		if err != nil {
			return nil, err
		}
		// get dates
		dateCert := lus.GetDateCertificate(certX509)
		issuedTime = dateCert["issuedTime"]
		expiresTime = dateCert["expiresTime"]

		// get attrs
		attrs = modeltools.GetAttrsCert(certX509)
	}

	// creator := &privateIdentityResponse{}
	creatorID := ""
	if request.CreatorDid != "" {
		// get issuer Id from DID
		creator, err := ci.getParticipant(ctx, request.CreatorDid)
		if creator == nil || err != nil {
			return nil, fmt.Errorf("failed to get Creator identity: %v", err)
		}
		// TODO: debug then
		creatorID = creator.ID
	} else {
		creatorID = ""
	}
	// get default issuer
	if request.IssuerID == "" {
		issuerResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeIssuerByDefault, []string{})
		if err != nil {
			return nil, err
		}
		defer issuerResultsIterator.Close()
		if issuerResultsIterator.HasNext() {
			responseRange, err := issuerResultsIterator.Next()
			if responseRange == nil {
				return nil, err
			}

			_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
			if err != nil {
				return nil, err
			}
			if len(compositeKeyParts) > 0 {
				request.IssuerID = compositeKeyParts[0]
			}
		} else {
			return nil, fmt.Errorf("there is no default issuer, pass the issuer id as parameter")
		}
	} else {
		// get Issuer by ID
		issuer, err := ci.GetIssuer(ctx, model.GetRequest{ID: request.IssuerID})
		if err != nil {
			return nil, err
		} else if issuer == nil {
			return nil, fmt.Errorf("failed to get issuer ")
		}
	}

	// timestamp when the transaction was created, have the same value across all endorsers
	txTimestamp, err := lus.GetTxTimestampRFC3339(ctx.GetStub())
	if err != nil {
		return nil, err
	}

	if request.Roles == nil {
		// Use make to create an empty slice of string.
		request.Roles = make([]string, 0)
	}

	id := lus.GenerateUUIDFormatDate(ctx.GetStub())

	// Create Participant
	identity := model.Participant{
		DocType:     ParticipantDocType,
		ID:          id,
		Did:         did,
		PublicKey:   request.PublicKey,
		IssuerID:    request.IssuerID,
		Creator:     creatorID,
		Roles:       request.Roles,
		Attrs:       attrs,
		AttrsExtras: make(map[string]string),
		Time:        txTimestamp,
		IssuedTime:  issuedTime,
		ExpiresTime: expiresTime,
		Active:      true,
		MspID:       clientMSPID,
	}
	// compositeKey ID
	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(ParticipantDocType, []string{identity.ID})
	if err != nil {
		return nil, err
	}

	// JSON encoding of identity
	identityJE, _ := json.Marshal(identity)
	if err := ctx.GetStub().PutState(compositeKeyID, identityJE); err != nil {
		return nil, fmt.Errorf("failed to create identity: %v", err)
	}

	if err := lus.CreateIndex(ctx.GetStub(), ObjectTypeParticipantByDidUUID, []string{identity.Did, identity.ID}); err != nil {
		return nil, fmt.Errorf("could not create identity %v: %v", identity.ID, err)
	}

	return &model.ParticipantResponse{
		Did:     identity.Did,
		ID:      id,
		Roles:   identity.Roles,
		Creator: nil,
	}, nil
}

// DeleteParticipant
// TODO: debug
func (ci *ContractIdentity) DeleteParticipant(ctx contractapi.TransactionContextInterface, identityRequest model.ParticipantDeleteRequest) error {
	log.Printf("[%s][DeleteParticipant]", ctx.GetStub().GetChannelID())
	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return fmt.Errorf(err.Error())
	}

	// Get the MSP ID of submitting client identity
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf(lus.ErrorGetMSPID, err)
	}

	// get user
	userToRevoke, err := ci.getParticipant(ctx, identityRequest.UserDid)
	if err != nil {
		return fmt.Errorf("failed to get participant identity: %v", err)
	}
	if userToRevoke == nil {
		return fmt.Errorf(lus.ErrorDefaultNotExist, identityRequest.UserDid)
	}

	if userToRevoke.MspID != clientMSPID {
		return fmt.Errorf("client from org %v is not authorized to delete data from an identity generated by the org %v", clientMSPID, userToRevoke.MspID)
	}

	var callerID = userToRevoke.ID
	// if not the participant himself then we get his ID
	if identityRequest.UserDid != identityRequest.CallerDid {
		// get caller
		callerParticipant, err := ci.getParticipant(ctx, identityRequest.CallerDid)
		if err != nil {
			return fmt.Errorf("failed to get caller identity: %v", err)
		} else if callerParticipant == nil {
			return fmt.Errorf(lus.ErrorDefaultNotExist, identityRequest.CallerDid)
		}
		callerID = callerParticipant.ID
	}

	// participant composite KEY
	participantKey, err := ctx.GetStub().CreateCompositeKey(ParticipantDocType, []string{userToRevoke.ID})
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(participantKey)
	if err != nil {
		return fmt.Errorf("failed to delete identity %s: %v", userToRevoke.Did, err)
	}

	// preparing the composite key to prune all in worldState
	if err = lus.DeleteIndex(ctx.GetStub(), ObjectTypeParticipantByDidUUID, []string{userToRevoke.Did, userToRevoke.ID}, true); err != nil {
		return err
	}

	// TODO: cambiar orden y agregar un index using couchdb
	// index
	deletedKey, err := ctx.GetStub().CreateCompositeKey(ObjectTypeParticipantDeleted, []string{Deleted, userToRevoke.ID, userToRevoke.Did})
	if err != nil {
		return err
	}

	// timestamp when the transaction was created, have the same value across all endorsers
	txTimestamp, err := lus.GetTxTimestampRFC3339(ctx.GetStub())
	if err != nil {
		return err
	}

	payload := &model.ParticipantDeletedPayload{
		MspID:    clientMSPID,
		Time:     txTimestamp,
		CallerID: callerID, // keep a record of the user who deleted the identity
	}
	// JSON encoding of payload
	payloadJE, _ := json.Marshal(payload)

	if err := ctx.GetStub().PutState(deletedKey, payloadJE); err != nil {
		return fmt.Errorf("could not create deleted index %v: %v", deletedKey, err)
	}
	return err
}

func (ci *ContractIdentity) DisarmParticipant(ctx contractapi.TransactionContextInterface, identityRequest model.ParticipantDeleteRequest) (bool, error) {
	log.Printf("[%s][DisarmParticipant]", ctx.GetStub().GetChannelID())
	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return false, err
	}

	// get user
	userToRevoke, err := ci.getParticipant(ctx, identityRequest.UserDid)
	if err != nil {
		return false, err
	}
	if userToRevoke == nil {
		return false, fmt.Errorf(lus.ErrorDefaultNotExist, identityRequest.UserDid)
	}

	userToRevoke.Active = false
	identityJE, _ := json.Marshal(&userToRevoke.identityAlias)
	// Put index entry
	if err = ctx.GetStub().PutState(userToRevoke.ID, identityJE); err != nil {
		return false, err
	}

	return true, nil
}

// ParticipantRenewRequest
type ParticipantRenewRequest struct {
	Did       string   `json:"did"`
	CertPem   string   `json:"certPem"`
	Signature string   `json:"signature"`
	Roles     []string `json:"roles,omitempty" metadata:",optional"` // role id list
}

func (ci *ContractIdentity) RenewParticipant(ctx contractapi.TransactionContextInterface, request ParticipantRenewRequest) error {
	log.Printf("[%s][RenewParticipant]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return err
	}

	// Get the MSP ID of submitting client identity
	_, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf(lus.ErrorGetMSPID, err)
	}

	did := request.Did
	// check the did format
	if err := modeltools.MatchDidFormat(did); err != nil {
		return err
	}

	exist, err := ci.ParticipantExits(ctx, model.ParticipantGetRequest{Did: did})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf(lus.ErrorIdentityExists, did)
	}

	exist, err = lus.CertificateAlreadyExists(ctx, request.CertPem, ParticipantDocType, []string{})
	if err != nil {
		return err
	} else if exist {
		return fmt.Errorf("an identity with the same certificate already exists")
	}

	return nil
}

func (ci *ContractIdentity) GetParticipant(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) (*model.ParticipantResponse, error) {
	log.Printf("[%s][GetParticipant]", ctx.GetStub().GetChannelID())
	identity, err := ci.getParticipant(ctx, request.Did)
	if err != nil {
		return nil, err
	} else if identity == nil {
		return nil, fmt.Errorf(lus.ErrorDefaultNotExist, request.Did)
	}
	return &model.ParticipantResponse{
		Did:     identity.Did,
		Roles:   identity.Roles,
		Creator: identity.Creator,
	}, nil
}

func (ci *ContractIdentity) getParticipant(ctx contractapi.TransactionContextInterface, did string) (*privateIdentityResponse, error) {
	log.Printf("[%s][getParticipant]", ctx.GetStub().GetChannelID())
	identityResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeParticipantByDidUUID, []string{did})
	if err != nil {
		return nil, fmt.Errorf(lus.ErrorGetIdentity, did)
	}
	defer identityResultsIterator.Close()

	if identityResultsIterator.HasNext() {
		responseRange, err := identityResultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}
		if len(compositeKeyParts) > 1 {
			log.Printf("participant DID-ID: %v", compositeKeyParts)

			returnedIdentityID := compositeKeyParts[1]
			response, err := ci.readParticipant(ctx, returnedIdentityID)
			if err != nil {
				return nil, err
			}
			return response, nil
		}
	}

	return nil, nil
}

// GetParticipantHistory returns the chain of custody for a identity since issuance
//
// Arguments:
//		0: model_api.ParticipantGetRequest
// Returns:
//		0: []model_api.ParticipantHistoryQueryResponse
//		1: error
func (ci *ContractIdentity) GetParticipantHistory(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) ([]model.ParticipantHistoryQueryResponse, error) {
	log.Printf("GetParticipantHistory: ID %v", request.Did)

	identity, err := ci.getParticipant(ctx, request.Did)
	if err != nil {
		return nil, err
	} else if identity == nil {
		return nil, fmt.Errorf(lus.ErrorDefaultNotExist, request.Did)
	}

	// compositeKey ID
	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(ParticipantDocType, []string{identity.ID})
	if err != nil {
		return nil, err
	}
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(compositeKeyID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []model.ParticipantHistoryQueryResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		identity := model.ParticipantQueryResponse{}
		identity.ParticipantID = request.Did
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &identity)
			if err != nil {
				return nil, err
			}
		}

		timestamp := modeltools.GetTimestampRFC3339(response.Timestamp)

		record := model.ParticipantHistoryQueryResponse{
			TxID:     response.TxId,
			Time:     timestamp,
			Record:   &identity,
			IsDelete: response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

// GetParticipants get all identity
//
// Arguments:
//		0: none
// Returns:
//		0: []model_api.ParticipantResponse
//		1: error
func (ci *ContractIdentity) GetParticipants(ctx contractapi.TransactionContextInterface) ([]model.ParticipantResponse, error) {
	log.Printf("[%s][GetParticipants]", ctx.GetStub().GetChannelID())

	identitiesResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ParticipantDocType, []string{})
	if err != nil {
		return nil, err
	}
	defer identitiesResultsIterator.Close()

	var identities []model.ParticipantResponse
	if identitiesResultsIterator.HasNext() {
		responseRange, err := identitiesResultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}

		var identity model.ParticipantResponse
		err = json.Unmarshal(responseRange.Value, &identity)
		if err != nil {
			return nil, err
		}
		identities = append(identities, identity)
	}
	return identities, nil
}

// ParticipantExits returns true when identity with given key exists in the worldState.
func (ci *ContractIdentity) ParticipantExits(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) (bool, error) {
	log.Printf("[%s][ParticipantExits]", ctx.GetStub().GetChannelID())
	identityResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeParticipantByDidUUID, []string{request.Did})
	if err != nil {
		return false, fmt.Errorf("failed to read identity %s from world state. %v", request.Did, err)
	}
	defer identityResultsIterator.Close()

	return identityResultsIterator.HasNext(), nil
}

// readParticipant
func (ci *ContractIdentity) readParticipant(ctx contractapi.TransactionContextInterface, participantID string) (*privateIdentityResponse, error) {
	log.Printf("readParticipant ")
	// compositeKey ID
	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(ParticipantDocType, []string{participantID})
	if err != nil {
		return nil, err
	}

	identityBytes, err := ctx.GetStub().GetState(compositeKeyID)
	if err != nil {
		return nil, fmt.Errorf(lus.ErrorGetIdentity, participantID)
	}
	if identityBytes == nil {
		return nil, fmt.Errorf(lus.ErrorDefaultNotExist, participantID)
	}

	var identity model.Participant
	err = json.Unmarshal(identityBytes, &identity)
	if err != nil {
		return nil, err
	}

	// TODO: find and insert Creator
	jcp := privateIdentityResponse{identityAlias: (*identityAlias)(&identity)}

	return &jcp, nil
}
