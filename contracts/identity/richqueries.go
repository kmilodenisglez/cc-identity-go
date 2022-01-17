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
func (ci *ContractIdentity) QueryAssetsWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*model.PaginatedQueryResponse, error) {
	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*model.PaginatedQueryResponse, error) {
	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	assets, err := libUtils.ConstructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &model.PaginatedQueryResponse{
		Records:             assets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}
