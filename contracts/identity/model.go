package identity

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	model "github.com/ic-matcom/model-identity-go/model"
)

// ContractIdentity chaincode that defines the business logic for managing identity
type ContractIdentity struct {
	contractapi.Contract
}

type identityAlias model.Participant
type privateIdentityResponse struct {
	*identityAlias
	Creator *model.ParticipantCreateRequest `json:"issuer,omitempty" metadata:",optional"` // issuer
}
