package identity

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	modelapi "github.com/ic-matcom/model-identity-go/model"
	"log"
)

// TODO: remove model-traceability-go dependence
// CreateRole
func (ic *ContractIdentity) CreateRole(ctx contractapi.TransactionContextInterface, request modelapi.RoleCreateRequest) (*modelapi.RoleResponse, error) {
	log.Printf("[%s][CreateRole]", ctx.GetStub().GetChannelID())

	id := lus.GenerateUUID()
	key, err := ctx.GetStub().CreateCompositeKey(RoleDocType, []string{id})
	if err != nil {
		return nil, err
	}

	cFunctions := make(map[string]string, 0)
	// using map, because it is very fast
	lus.SliceToMap(request.ContractFunctions, cFunctions)
	// Create Role
	role := &Role{
		DocType:           RoleDocType,
		ID:                id,
		Name:              request.Name,
		ContractFunctions: cFunctions,
	}
	// JSON encoding
	roleJE, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(key, roleJE); err != nil {
		return nil, fmt.Errorf("role %s could not be created: %v", request.Name, err)
	}
	return &modelapi.RoleResponse{
		DocType:           role.DocType,
		ID:                role.ID,
		Name:              role.Name,
		ContractFunctions: request.ContractFunctions,
	}, nil
}

// GetRole
func (ic *ContractIdentity) GetRole(ctx contractapi.TransactionContextInterface, request modelapi.GetRequest) (*modelapi.RoleResponse, error) {
	log.Printf("[%s][GetRole]", ctx.GetStub().GetChannelID())

	key, err := ctx.GetStub().CreateCompositeKey(RoleDocType, []string{request.ID})
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

	var itemJD Role
	err = json.Unmarshal(item, &itemJD)
	if err != nil {
		return nil, err
	}

	return &modelapi.RoleResponse{
		DocType:           itemJD.DocType,
		ID:                itemJD.ID,
		Name:              itemJD.Name,
		ContractFunctions: lus.MapToSlice(itemJD.ContractFunctions),
	}, nil
}

// GetRoles get all role
func (ic *ContractIdentity) GetRoles(ctx contractapi.TransactionContextInterface) ([]modelapi.RoleResponse, error) {
	log.Printf("[%s][GetRoles]", ctx.GetStub().GetChannelID())

	rolesResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(RoleDocType, []string{})
	if err != nil {
		return nil, err
	}
	defer rolesResultsIterator.Close()

	var items []modelapi.RoleResponse
	if rolesResultsIterator.HasNext() {
		responseRange, err := rolesResultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}

		var role Role
		err = json.Unmarshal(responseRange.Value, &role)
		if err != nil {
			return nil, err
		}
		items = append(items, modelapi.RoleResponse{
			DocType:           role.DocType,
			ID:                role.ID,
			Name:              role.Name,
			ContractFunctions: lus.MapToSlice(role.ContractFunctions),
		})
	}
	return items, nil
}

// UpdateRole
func (ic *ContractIdentity) UpdateRole(ctx contractapi.TransactionContextInterface, request modelapi.RoleUpdateRequest) error {
	log.Printf("[%s][UpdateRole]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(RoleDocType, []string{request.ID})
	if err != nil {
		return err
	}

	role, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("failed to get a role: %v", err)
	} else if role == nil {
		return fmt.Errorf("no state found for %s", key)
	}

	var roleJD Role
	err = json.Unmarshal(role, &roleJD)
	if err != nil {
		return err
	}

	// using map, because it is very fast
	lus.SliceToMap(request.ContractFunctions, roleJD.ContractFunctions)

	// JSON encoding
	roleJE, err := json.Marshal(roleJD)
	if err != nil {
		return err
	}
	if err := ctx.GetStub().PutState(key, roleJE); err != nil {
		return fmt.Errorf("role %s could not be updated: %v", roleJD.ID, err)
	}

	return nil
}

// DeleteRole
func (ic *ContractIdentity) DeleteRole(ctx contractapi.TransactionContextInterface, request modelapi.GetRequest) error {
	log.Printf("[%s][DeleteRole]", ctx.GetStub().GetChannelID())
	if err := lus.DeleteIndex(ctx.GetStub(), RoleDocType, []string{request.ID}, true); err != nil {
		return err
	}

	return nil
}
