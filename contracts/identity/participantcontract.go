package identity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-traceability-go"
	"log"
)

// InitLedger adds a base set of data to the ledger
func (ic *ContractIdentity) InitLedger(ctx contractapi.TransactionContextInterface) error {
	log.Printf("[%s][InitLedger]", ctx.GetStub().GetChannelID())
	// check if client-node is connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return fmt.Errorf(err.Error())
	}
	accessFuelBatch := AccessCreateRequest{
		ContractName:      ic.Name,              // contract name
		ContractFunctions: ic.GetTransactions(), // functions name
	}

	// create fuelBatch access
	_, err := ic.CreateAccess(ctx, accessFuelBatch)
	if err != nil {
		return err
	}

	return nil
}

// OnlyDevParticipant [temporary] function to populate with test data
func (ic *ContractIdentity) OnlyDevParticipant(ctx contractapi.TransactionContextInterface) (string, error) {
	log.Printf("[%s][OnlyDevParticipant]", ctx.GetStub().GetChannelID())
	const b64UserWithAttrsCert = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlHRWpDQ0EvcWdBd0lCQWdJVWRyd2M0NFhvK2ZLQ1JzTXlIdHd6ZDBnQ01iVXdEUVlKS29aSWh2Y05BUUVODQpCUUF3Z2FVeEN6QUpCZ05WQkFZVEFrTlZNUkl3RUFZRFZRUUlEQWxNWVNCSVlXSmhibUV4RmpBVUJnTlZCQWNNDQpEVU5sYm5SeWJ5QklZV0poYm1FeEp6QWxCZ05WQkFvTUhrMXBibWx6ZEdWeWFXOGdaR1VnUlc1bGNtZkRyV0VnDQplU0JOYVc1aGN6RU9NQXdHQTFVRUN3d0ZRM1Z3WlhReE1UQXZCZ05WQkFNTUtFRjFkRzl5YVdSaFpDQmtaU0JEDQpaWEowYVdacFkyRmphY096YmlCVVpXTnViMjNEb1hScFkyRXdIaGNOTWpFd056TXdNVGN6TlRJM1doY05Nak13DQpOek13TVRjek5USTJXakNCcWpFYk1Ca0dDZ21TSm9tVDhpeGtBUUVNQ3pnM01EVXhNakV4TkRVM01SNHdIQVlEDQpWUVFEREJWWmFYTmxiQ0JCYzNScFlYcGhjbUZwYmlCRWFXNHhIVEFiQmdOVkJBd01GRVZ6Y0M0Z1FpQkRhV1Z1DQpZMmxoY3lCSmJtWXVNUlF3RWdZRFZRUUxEQXREZFhCbGRDMU5hVzVsYlRFVk1CTUdBMVVFQ2d3TVZHVmpibTl0DQp3NkYwYVdOaE1SSXdFQVlEVlFRSURBbE1ZU0JJWVdKaGJtRXhDekFKQmdOVkJBWVRBa05WTUZrd0V3WUhLb1pJDQp6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVDNVNwZTlRem12WldVcnBLMHo0bDJVYjVwcVczZEs4OXlzV1k3d0xHDQpUMldybjFwSHFLckpHM0NXdFl6QVllaW9aRlA1bENJTjdHUE5ycXNlWUY1S2thT0NBZnd3Z2dINE1Bd0dBMVVkDQpFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVU1JRcjNXaTNjSjNWdmhXS2tCSVk0VWNSSkprd1ZRWUlLd1lCDQpCUVVIQVFFRVNUQkhNQ01HQ0NzR0FRVUZCekFDaGhkb2RIUndjem92TDNCcmFTNWpkWEJsZEM1amRTOWpZVEFnDQpCZ2dyQmdFRkJRY3dBWVlVYUhSMGNEb3ZMMjlqYzNBdVkzVndaWFF1WTNVd0tRWURWUjB1QkNJd0lEQWVvQnlnDQpHb1lZYUhSMGNEb3ZMMlJsYkhSaFkzSnNMbU4xY0dWMExtTjFNRDRHQTFVZEpRUTNNRFVHQ0NzR0FRVUZCd01DDQpCZ2dyQmdFRkJRY0RBd1lJS3dZQkJRVUhBd1FHQ2lzR0FRUUJnamNLQXd3R0NTcUdTSWIzTHdFQkJUQ0IxUVlEDQpWUjBmQklITk1JSEtNSUhIb0JlZ0ZZWVRhSFIwY0RvdkwyTnliQzVqZFhCbGRDNWpkYUtCcTZTQnFEQ0JwVEV4DQpNQzhHQTFVRUF3d29RWFYwYjNKcFpHRmtJR1JsSUVObGNuUnBabWxqWVdOcHc3TnVJRlJsWTI1dmJjT2hkR2xqDQpZVEVPTUF3R0ExVUVDd3dGUTNWd1pYUXhKekFsQmdOVkJBb01IazFwYm1semRHVnlhVzhnWkdVZ1JXNWxjbWZEDQpyV0VnZVNCTmFXNWhjekVXTUJRR0ExVUVCd3dOUTJWdWRISnZJRWhoWW1GdVlURVNNQkFHQTFVRUNBd0pUR0VnDQpTR0ZpWVc1aE1Rc3dDUVlEVlFRR0V3SkRWVEFkQmdOVkhRNEVGZ1FVbkZXUFFRVkpDSGt3STNCUDVMSFVHMy9sDQozVjB3RGdZRFZSMFBBUUgvQkFRREFnWGdNQTBHQ1NxR1NJYjNEUUVCRFFVQUE0SUNBUUE5TTBGZlVMZVcySEkrDQpNNFZDNEZRczZybDZ0dmJ6NHFRalZ3MWZDbnNNVVNxZDBUejB0eTdWdGxCcGJteXFhNnRqSitFYmxncUVGOTBTDQpseHlnN1NyY1RkMGxWVFlIaExURUhhWFAzWENCVkhRUDVhVUIxQmZYR0pkNGJwaEZDUk5pQ1ExajhXa0xUTy8vDQo3bWNJSW9vSVh5Zk93K1N5d085VzFhZUt2amV5OEVrVTQ5bUtIakZXbC92Ritic0NPQUNwK1dCeGZQNFgrbm9yDQowdGVVR0MrZGZiYkY3OTM5c2tTdEU2SEwvTzRRT3RqeWZXWFZVLzhDNWRjMlpHMTJiSzZXOE9Hc294bTJyMFkrDQpSVkhOdWNPemhTTHJEL3B6ZDBLS1BVSE5ma1NmM3dBenlJYzM4amVreWx3ZFlHdnk5VWV0dER2NFJ2cFYrSE9BDQo2M1FUVDlFbmlvQkNQVHZTbTVMb2Q2eVpsd0xLbzBNakdJZXNmNG9tU3RsSWxIcUE0ZUppL2V5dzdxZmdZdTk4DQptY2t6dVNhVWlQVE1YaUgwWC84dDFlVEkzRHpjVEo0Uytmc251WEd0OGs5WEFWa2w4elNzM2xROERNVG1UVll3DQpTSG1ERm5hQ0psbDRlazJUNlpOWVB6dmFnOWdBVGxBQUMzSEw4ZGlycXdLN2FJeURveENPSHU3a0JJQU0xK202DQpDb0ZPcnVmalZvVVRtTzlUWUlocUlXZWJOMWY5Z3hXSkVBbnF6Zm4wUkVSa3NTMU0zV0JnaGZWY0xoeDY3WEd3DQpzTXltdHZLQ1hWRkpYSmNJaW5kRTUyYVAyQzVnQ1NVU1VyVVJOcTYvRjNqUVB1UE1ySW1ydHo1ZWoyV0tKVmhBDQpuS0hseFN2UkdYTXNuSDVoVmpibWZBU3B0cU9CQ3c9PQ0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ0K"

	// insert identityGetAccess
	hashPublicKey := sha256.Sum256([]byte("valid did"))
	did, _ := model.CreateDid(hex.EncodeToString(hashPublicKey[:]))

	roles, err := ic.GetRoles(ctx)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	} else if len(roles) < 1 {
		return "", fmt.Errorf("there is no role in the ledger")
	}

	identityRequest := model.ParticipantCreateRequest{
		Did:     did,
		CertPem: b64UserWithAttrsCert,
		Roles:   []string{roles[0].ID},
	}
	identity, err := ic.CreateParticipant(ctx, identityRequest)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return identity.Did, nil
}

// TODO: missing validates ca cert - participant cer
func (ic *ContractIdentity) CreateParticipant(ctx contractapi.TransactionContextInterface, identityRequest model.ParticipantCreateRequest) (*Participant, error) {
	log.Printf("[%s][CreateParticipant]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Get the MSP ID of submitting client identity
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("failed to get verified MSPID: %v", err)
	}

	did := identityRequest.Did
	// check the did format
	if err := model.CheckDid(did); err != nil {
		return nil, err
	}

	exist, err := ic.ParticipantExits(ctx, model.ParticipantGetRequest{Did: did})
	if err != nil {
		return nil, fmt.Errorf("failed to get identity: %v", err)
	}
	if exist {
		return nil, fmt.Errorf("%s identity already exists", did)
	}

	exist, err = lus.CertificateAlreadyExists(ctx, identityRequest.CertPem, ParticipantDocType, []string{})
	if err != nil {
		return nil, err
	} else if exist {
		return nil, fmt.Errorf("an identity with the same certificate already exists")
	}

	// validate cert
	certX509, err := lus.GetX509CertFromPem(identityRequest.CertPem)
	if err != nil {
		return nil, err
	}
	err = lus.HasExpired(certX509)
	if err != nil {
		return nil, err
	}

	// creator := &privateIdentityResponse{}
	creatorID := ""
	if identityRequest.CreatorDid != "" {
		// get issuer Id from DID
		creator, err := ic.getParticipant(ctx, identityRequest.CreatorDid)
		if creator == nil || err != nil {
			return nil, fmt.Errorf("failed to get Creator identity: %v", err)
		}
		// TODO: debug then
		creatorID = creator.ID
	} else {
		creatorID = ""
	}
	// get default issuer
	if identityRequest.IssuerID == "" {
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
				identityRequest.IssuerID = compositeKeyParts[0]
			}
		} else {
			return nil, fmt.Errorf("there is no default issuer, pass the issuer id as parameter")
		}
	} else {
		// get Issuer by ID
		issuer, err := ic.GetIssuer(ctx, model.GetRequest{ID: identityRequest.IssuerID})
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

	if identityRequest.Roles == nil {
		// Use make to create an empty slice of string.
		identityRequest.Roles = make([]string, 0)
	}

	// get dates
	dateCert := lus.GetDateCertificate(certX509)

	// get attrs
	attrs := lus.GetAttrsCert(certX509)
	// Create Participant
	identity := Participant{
		DocType:     ParticipantDocType,
		ID:          lus.GenerateUUID(),
		Did:         did,
		CertPem:     identityRequest.CertPem,
		IssuerID:    identityRequest.IssuerID,
		Creator:     creatorID,
		Roles:       identityRequest.Roles,
		Attrs:       attrs,
		AttrsExtras: make(map[string]string, 0),
		Time:        txTimestamp,
		IssuedTime:  dateCert["issuedTime"],
		ExpiresTime: dateCert["expiresTime"],
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

	if err := lus.CreateIndex(ctx.GetStub(), ObjectTypeParticipantByDidUuid, []string{identity.Did, identity.ID}); err != nil {
		return nil, fmt.Errorf("could not create identity %v: %v", identity.ID, err)
	}

	return &identity, nil
}

// TODO: debug
func (ic *ContractIdentity) DeleteParticipant(ctx contractapi.TransactionContextInterface, identityRequest model.ParticipantDeleteRequest) error {
	log.Printf("[%s][DeleteParticipant]", ctx.GetStub().GetChannelID())
	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return fmt.Errorf(err.Error())
	}

	// Get the MSP ID of submitting client identity
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get verified MSPID: %v", err)
	}

	// get user
	userToRevoke, err := ic.getParticipant(ctx, identityRequest.UserDid)
	if err != nil {
		return fmt.Errorf("failed to get participant identity: %v", err)
	}
	if userToRevoke == nil {
		return fmt.Errorf("%s does not exist", identityRequest.UserDid)
	}

	if userToRevoke.MspID != clientMSPID {
		return fmt.Errorf("client from org %v is not authorized to delete data from an identity generated by the org %v", clientMSPID, userToRevoke.MspID)
	}

	var callerID = userToRevoke.ID
	// if not the participant himself then we get his ID
	if identityRequest.UserDid != identityRequest.CallerDid {
		// get caller
		callerParticipant, err := ic.getParticipant(ctx, identityRequest.CallerDid)
		if err != nil {
			return fmt.Errorf("failed to get caller identity: %v", err)
		} else if callerParticipant == nil {
			return fmt.Errorf("%s does not exist", identityRequest.CallerDid)
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
	if err = lus.DeleteIndex(ctx.GetStub(), ObjectTypeParticipantByDidUuid, []string{userToRevoke.Did, userToRevoke.ID}, true); err != nil {
		return err
	}

	// index
	deletedKey, err := ctx.GetStub().CreateCompositeKey(ObjectTypeParticipantDeleted, []string{Deleted, userToRevoke.Did, userToRevoke.ID})
	if err != nil {
		return err
	}

	// timestamp when the transaction was created, have the same value across all endorsers
	txTimestamp, err := lus.GetTxTimestampRFC3339(ctx.GetStub())
	if err != nil {
		return err
	}

	payload := &model.ParticipantDeletedPayload{
		MspId:    clientMSPID,
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

func (ic *ContractIdentity) DisarmParticipant(ctx contractapi.TransactionContextInterface, identityRequest model.ParticipantDeleteRequest) (bool, error) {
	log.Printf("[%s][DisarmParticipant]", ctx.GetStub().GetChannelID())
	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return false, fmt.Errorf(err.Error())
	}

	// get user
	userToRevoke, err := ic.getParticipant(ctx, identityRequest.UserDid)
	if err != nil {
		return false, fmt.Errorf("failed to get User identity: %v", err)
	}
	if userToRevoke == nil {
		return false, fmt.Errorf("%s does not exist", identityRequest.UserDid)
	}

	userToRevoke.Active = false
	identityJE, _ := json.Marshal(&userToRevoke.identityAlias)
	// Put index entry
	if err = ctx.GetStub().PutState(userToRevoke.ID, identityJE); err != nil {
		return false, err
	}

	return true, nil
}

func (ic *ContractIdentity) GetParticipant(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) (*model.ParticipantResponse, error) {
	log.Printf("[%s][GetParticipant]", ctx.GetStub().GetChannelID())
	identity, err := ic.getParticipant(ctx, request.Did)
	if err != nil {
		return nil, fmt.Errorf("GetParticipant: failed to get identity: %v", err)
	} else if identity == nil {
		return nil, fmt.Errorf("GetParticipant: no state found for %s", request.Did)
	}
	return &model.ParticipantResponse{
		Did:     identity.Did,
		Roles:   identity.Roles,
		Creator: identity.Creator,
	}, nil
}

func (ic *ContractIdentity) getParticipant(ctx contractapi.TransactionContextInterface, did string) (*privateIdentityResponse, error) {
	log.Printf("[%s][getParticipant]", ctx.GetStub().GetChannelID())
	identityResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeParticipantByDidUuid, []string{did})
	if err != nil {
		return nil, err
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
			response, err := ic.readParticipant(ctx, returnedIdentityID)
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
//		0: model.ParticipantGetRequest
// Returns:
//		0: []model.ParticipantHistoryQueryResponse
//		1: error
func (ic *ContractIdentity) GetParticipantHistory(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) ([]model.ParticipantHistoryQueryResponse, error) {
	log.Printf("GetParticipantHistory: ID %v", request.Did)

	identity, err := ic.getParticipant(ctx, request.Did)
	if err != nil {
		return nil, err
	} else if identity == nil {
		return nil, fmt.Errorf("no state found for %s", request.Did)
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

		timestamp := model.GetTimestampRFC3339(response.Timestamp)

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
//		0: []model.ParticipantResponse
//		1: error
func (ic *ContractIdentity) GetParticipants(ctx contractapi.TransactionContextInterface) ([]model.ParticipantResponse, error) {
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
func (ic *ContractIdentity) ParticipantExits(ctx contractapi.TransactionContextInterface, request model.ParticipantGetRequest) (bool, error) {
	log.Printf("[%s][ParticipantExits]", ctx.GetStub().GetChannelID())
	identityResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeParticipantByDidUuid, []string{request.Did})
	if err != nil {
		return false, fmt.Errorf("failed to read identity %s from world state. %v", request.Did, err)
	}
	defer identityResultsIterator.Close()

	return identityResultsIterator.HasNext(), nil
}

// readParticipant
func (ic *ContractIdentity) readParticipant(ctx contractapi.TransactionContextInterface, participantID string) (*privateIdentityResponse, error) {
	log.Printf("readParticipant ")
	// compositeKey ID
	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(ParticipantDocType, []string{participantID})
	if err != nil {
		return nil, err
	}

	identityBytes, err := ctx.GetStub().GetState(compositeKeyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity %s: %v", participantID, err)
	}
	if identityBytes == nil {
		return nil, fmt.Errorf("%s does not exist", participantID)
	}

	var identity Participant
	err = json.Unmarshal(identityBytes, &identity)
	if err != nil {
		return nil, err
	}

	// TODO: find and insert Creator
	jcp := privateIdentityResponse{identityAlias: (*identityAlias)(&identity)}

	return &jcp, nil
}

// GetTransactions returns callables functions of Contract
// use to populate ContractFunctions field in identity:Access
func (ic *ContractIdentity) GetTransactions() []string {
	return []string{"CreateParticipant", "DeleteParticipant",
		"GetParticipant", "DisarmParticipant", "GetParticipants",
		"GetParticipantHistory", "CreateRole", "GetRole",
		"DeleteRole", "GetRoles", "UpdateRole", "UpdateAccess",
		"GetAccess", "GetAccesses", "CreateAccess"}
}
