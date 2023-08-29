package model

import (
	"brok/navnetjener/utils"
	"fmt"
)

// FindAllCapTablesForPerson returns all captables for a person
// first it looks in the database for a person matching the pnr
// then it uses orgnr from the database to find captables from TheGraph
func FindAllCapTablesForPerson(pnr string) ([]CapTable, error) {
	wallets, err := FindWalletByPnr(pnr)
	if err != nil {
		return []CapTable{}, err
	}

	var orgnrList []string
	for _, wallet := range wallets {
		orgnrList = append(orgnrList, fmt.Sprint(wallet.Orgnr))
	}

	captables, err := FindAllCaptableByOrgnrListFromTheGraph(orgnrList)
	if err != nil {
		return []CapTable{}, err
	}

	for i, captable := range captables {
		captables[i], err = mergeDataFromTheGraphAndDatabase(captable)
		if err != nil {
			return []CapTable{}, err
		}
	}

	return captables, nil

	// var captables []CapTable
	// for _, wallet := range wallets {
	// 	// TODO optimize this
	// 	captable, err := FindAllCaptableByOrgnrListFromTheGraph(fmt.Sprint(wallet.Orgnr))
	// 	if err != nil {
	// 		return []CapTable{}, err
	// 	}
	// 	captable, err = mergeDataFromTheGraphAndDatabase(captable)
	// 	if err != nil {
	// 		return []CapTable{}, err
	// 	}
	// 	captables = append(captables, captable)
	// }

	// return captables, nil
}

// findCaptableByOrgnr combines data from TheGraph and the database
// to return a CapTable struct with person data
func FindCaptableByOrgnr(orgnr string) (CapTable, error) {
	captable, err := FindCaptableByOrgnrFromTheGraph(orgnr)
	if err != nil {
		return CapTable{}, err
	}

	captable, err = mergeDataFromTheGraphAndDatabase(captable)
	if err != nil {
		return CapTable{}, err
	}

	return captable, nil
}

// FindForetak returns 25 foretak from TheGraph and the database
// user can use the queryparameter "page" to paginate
func FindForetak(page int) ([]CapTable, error) {
	captables, err := FindForetakFromTheGraph(page)
	if err != nil {
		return nil, err
	}

	for i, captable := range captables {
		captable, err := mergeDataFromTheGraphAndDatabase(captable)
		if err != nil {
			return nil, err
		}
		captables[i] = captable
	}

	return captables, nil
}

// mergeDataFromTheGraphAndDatabase combines data from TheGraph and the database
// return a CapTable struct with person data
func mergeDataFromTheGraphAndDatabase(captable CapTable) (CapTable, error) {
	captable.TotalSupply = utils.ToDecimal(captable.TotalSupply)

	for i, tokenHolder := range captable.TokenHolders {
		person, err := findPersonByWalletAddress(tokenHolder.Address)
		tokenHolder = convertTokenHolderWeiToDecimals(tokenHolder)
		if err != nil {
			return CapTable{}, err
		}

		captable.TokenHolders[i].Person = person
	}

	return captable, nil
}

// ConvertTokenHolderWeiToDecimals converts the tokenHolder balance from wei to decimals
func convertTokenHolderWeiToDecimals(tokenHolder TokenHolder) TokenHolder {
	for i, balance := range tokenHolder.Balances {
		tokenHolder.Balances[i].Amount = utils.ToDecimal(balance.Amount)
	}
	return tokenHolder
}
