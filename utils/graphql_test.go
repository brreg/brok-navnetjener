package utils_test

import (
	"testing"
)

// ----------------------------------------
// This test is disabled because it requires a running Graph Node instance
// The node can be started from brok-monorepo
// ----------------------------------------
func TestShouldGetFirstCaptable(t *testing.T) {

	// // load env
	// t.Setenv("GRAPHQL_URL", "http://localhost:8000/subgraphs/name/brok/captable")
	// logrus.SetLevel(logrus.DebugLevel)

	// query := `
	// query getCaptable($num: Int!) {
	// 	capTables(first: $num) {
	// 		id
	// 		name
	// 		symbol
	// 		partitions
	// 		status
	// 		registry {
	// 			id
	// 		}
	// 		tokenHolders {
	// 			address
	// 			balances {
	// 				amount
	// 				partition
	// 			}
	// 		}
	// 		totalSupply
	// 		owner
	// 		minter
	// 		controllers
	// 		orgnr
	// 	}
	// }`

	// var response struct {
	// 	CapTables []model.CapTable `json:"capTables"`
	// }

	// // execute query

	// err := utils.ExecuteGraphQLQuery(query, map[string]interface{}{"num": 1}, &response)
	// if err != nil {
	// 	t.Errorf("Error: %v", err)
	// }

	// // log output
	// log.Printf("GraphQL response: %v", response)
}
