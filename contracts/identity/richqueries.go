package identity

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	libUtils "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-identity-go/model"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// QueryAssetsBy uses a query string to perform a query for any identity contract asset
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// Param Ex: {"selector":{"docType":"did.participant","id":"myID"}}
//
// Arguments:
//		0: queryStruct map[string]interface{}
// Returns:
//		0: []string
func (ci *ContractIdentity) QueryAssetsBy(ctx contractapi.TransactionContextInterface, query map[string]interface{}) ([]interface{}, error) {
	queryString, err := json.MarshalToString(&query)
	if err != nil {
		return nil, err
	}

	res, err := libUtils.GetQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryAssetsWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (ci *ContractIdentity) QueryAssetsWithPagination(ctx contractapi.TransactionContextInterface, request model.RichQuerySelector) (*model.PaginatedQueryResponse, error) {
	queryString, err := json.MarshalToString(&request.QueryString)
	if err != nil {
		return nil, err
	}
	//TODO: add validation: len(request.QueryString)

	return libUtils.GetQueryResultForQueryStringWithPagination(ctx, queryString, int32(request.PageSize), request.Bookmark)
}
