package model

import (
	"brok/navnetjener/utils"

	"github.com/sirupsen/logrus"
)

type GraphQLResponse struct {
	Data struct {
		CapTables []CapTable `json:"capTables"`
	} `json:"data"`
}

type CapTable struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Orgnr        string        `json:"orgnr"`
	Owner        string        `json:"owner"`
	Partitions   []string      `json:"partitions"`
	Status       string        `json:"status"`
	TokenHolders []TokenHolder `json:"tokenHolders"`
	TotalSupply  string        `json:"totalSupply"`
	Minter       string        `json:"minter"`
}

type TokenHolder struct {
	Address  string    `json:"address"`
	Balances []Balance `json:"balances"`
	Owner    Owner     `json:"owner"`
}

type Balance struct {
	Amount    string `json:"amount"`
	Partition string `json:"partition"`
}

// Use a standard query to get the captable from the graph
func FindAllCaptableByOrgnrListFromTheGraph(orgnrList []string) ([]CapTable, error) {
	query := `
	query GetMultipleCapTables($orgnrList: [String!]) {
		capTables(where: { orgnr_in: $orgnrList }) {
			id
			name
			partitions
			status
			registry {
				id
			}
			tokenHolders {
				address
				balances {
					amount
					partition
				}
			}
			totalSupply
			owner
			minter
			controllers
			orgnr
		}
	}
	`

	var response GraphQLResponse

	err := utils.ExecuteGraphQLQuery(query, map[string]interface{}{"orgnrList": orgnrList}, &response)
	if err != nil {
		return []CapTable{}, err
	}

	if response.Data.CapTables == nil || len(response.Data.CapTables) == 0 {
		return []CapTable{}, nil
	}

	return response.Data.CapTables, nil
}

// Use a standard query to get the captable from the graph
func FindCaptableByOrgnrFromTheGraph(orgnr string) (CapTable, error) {
	query := `
	query getCapTableForOrgnr($orgnr: String!) {
		capTables(where: { orgnr: $orgnr }) {
			id
			name
			partitions
			status
			registry {
				id
			}
			tokenHolders {
				address
				balances {
					amount
					partition
				}
			}
			totalSupply
			owner
			minter
			controllers
			orgnr
		}
	}
	`

	var response GraphQLResponse

	err := utils.ExecuteGraphQLQuery(query, map[string]interface{}{"orgnr": orgnr}, &response)
	if err != nil {
		logrus.Error(err)
		return CapTable{}, err
	}

	if response.Data.CapTables == nil || len(response.Data.CapTables) == 0 {
		return CapTable{}, nil
	}

	return response.Data.CapTables[0], nil
}

// FindForetakFromTheGraph returns 25 foretak from TheGraph
// user can use the queryparameter "page" to paginate
func FindForetakFromTheGraph(page int) ([]CapTable, error) {
	query := `
	query getCapTables($skip: Int!) {
		capTables(first: 25, skip: $skip) {
			id
			name
			partitions
			status
			registry {
				id
			}
			tokenHolders {
				address
				balances {
					amount
					partition
				}
			}
			totalSupply
			owner
			minter
			controllers
			orgnr
		}
	}
	`

	var response GraphQLResponse

	err := utils.ExecuteGraphQLQuery(query, map[string]interface{}{"skip": page * 25}, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.CapTables, nil
}
