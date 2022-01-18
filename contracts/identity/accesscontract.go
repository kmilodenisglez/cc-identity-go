package identity

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-identity-go/model"
	"log"
)

// CreateAccess
//
// Arguments:
//		0: AccessCreateRequest -
// Returns:
//		0: AccessResponse
//		1: error
func (ci *ContractIdentity) CreateAccess(ctx contractapi.TransactionContextInterface, request model.AccessCreateRequest) (*model.AccessResponse, error) {
	log.Printf("[%s][CreateAccess]", ctx.GetStub().GetChannelID())
	lowerNonSpace := lus.NormalizeString(request.ContractName)

	key, err := ctx.GetStub().CreateCompositeKey(AccessDocType, []string{lowerNonSpace})
	if err != nil {
		return nil, err
	}

	cFunctions := make(map[string]string)
	// using map, because it is very fast
	lus.SliceToMap(request.ContractFunctions, cFunctions)

	// Create Access
	access := &model.Access{
		DocType:           AccessDocType,
		ID:                lowerNonSpace,
		ContractFunctions: cFunctions,
	}
	// JSON encoding
	accessJE, _ := json.Marshal(access)

	if err := ctx.GetStub().PutState(key, accessJE); err != nil {
		return nil, fmt.Errorf("access %s could not be created: %v", request.ContractName, err)
	}
	return &model.AccessResponse{
		DocType:           access.DocType,
		ID:                access.ID,
		ContractFunctions: request.ContractFunctions,
	}, nil
}

// GetAccess get an access
//
// Arguments:
//		0: GetRequest
// Returns:
//		0: AccessResponse
//		1: error
func (ci *ContractIdentity) GetAccess(ctx contractapi.TransactionContextInterface, request model.GetRequest) (*model.AccessResponse, error) {
	log.Printf("[%s][GetAccess]", ctx.GetStub().GetChannelID())

	key, err := ctx.GetStub().CreateCompositeKey(AccessDocType, []string{request.ID})
	if err != nil {
		return nil, fmt.Errorf("error happened creating composite key: %v", err)
	} else if key == "" {
		return nil, fmt.Errorf("no state found for %s", request.ID)
	}
	item, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %v", err)
	} else if item == nil {
		return nil, fmt.Errorf("no state found for %s", key)
	}
	var itemJD model.Access
	err = json.Unmarshal(item, &itemJD)
	if err != nil {
		return nil, err
	}
	return &model.AccessResponse{
		DocType:           itemJD.DocType,
		ID:                itemJD.ID,
		ContractFunctions: lus.MapToSlice(itemJD.ContractFunctions),
	}, nil
}

// GetAccesses get all access
//
// Arguments:
//		0: none
// Returns:
//		0: []model.AccessResponse
//		1: error
func (ci *ContractIdentity) GetAccesses(ctx contractapi.TransactionContextInterface, request model.RichQuerySelector) (*model.PaginatedQueryResponse, error) {
	log.Printf("[%s][GetAccesses]", ctx.GetStub().GetChannelID())

	resultsIterator, responseMetadata, err := ctx.GetStub().GetStateByPartialCompositeKeyWithPagination(AccessDocType, []string{}, int32(request.PageSize), request.Bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var items = make([]interface{}, 0)
	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}
		var item model.Access
		err = json.Unmarshal(responseRange.Value, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, model.AccessResponse{
			DocType:           item.DocType,
			ID:                item.ID,
			Description:       item.Description,
			ContractFunctions: lus.MapToSlice(item.ContractFunctions),
		})
	}

	return &model.PaginatedQueryResponse{
		Records:             items,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

// updateAccess
func (ci *ContractIdentity) updateAccess(ctx contractapi.TransactionContextInterface, request model.AccessUpdateRequest) error {
	log.Printf("[%s][updateAccess]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(AccessDocType, []string{request.ID})
	if err != nil {
		return err
	}

	role, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("failed to get a role: %v", err)
	} else if role == nil {
		return fmt.Errorf("no state found for %s", key)
	}

	var accessJD model.Access
	err = json.Unmarshal(role, &accessJD)
	if err != nil {
		return err
	}

	// using map, because it is very fast
	lus.SliceToMap(request.ContractFunctions, accessJD.ContractFunctions)

	// JSON encoding
	accessJE, err := json.Marshal(accessJD)
	if err != nil {
		return err
	}
	if err := ctx.GetStub().PutState(key, accessJE); err != nil {
		return fmt.Errorf("role %s could not be updated: %v", accessJD.ID, err)
	}

	return nil
}

// deleteAccess
func (ci *ContractIdentity) deleteAccess(ctx contractapi.TransactionContextInterface, request model.GetRequest) error {
	log.Printf("[%s][deleteAccess]", ctx.GetStub().GetChannelID())
	if err := lus.DeleteIndex(ctx.GetStub(), AccessDocType, []string{request.ID}, true); err != nil {
		return err
	}

	return nil
}
