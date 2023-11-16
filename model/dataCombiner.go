package model

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// FindAllCapTables returns all captables for a person or organization
// first it looks in the database for a owner matching the id
// then it uses orgnr from the database to find captables from TheGraph
func FindAllCapTables(id string) ([]CapTable, error) {
	wallets, err := FindWallet(id)
	if err != nil {
		return []CapTable{}, err
	}

	var orgnrList []string
	for _, wallet := range wallets {
		orgnrList = append(orgnrList, fmt.Sprint(wallet.CapTableOrgnr))
	}

	logrus.Debug("found orgnrList in db with id: ", id[0:6], "*****", " orgnrList: ", orgnrList)

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

func filterOutOwnerWithNoShares(captable CapTable) CapTable {
	tokenHoldersWithShares := []TokenHolder{}
	for _, tokenHolder := range captable.TokenHolders {
		if tokenHolder.Balances[0].Amount != "0" {
			tokenHoldersWithShares = append(tokenHoldersWithShares, tokenHolder)
		}
	}
	captable.TokenHolders = tokenHoldersWithShares
	return captable
}

func FindNumberOfSharesForOwnerOfCaptable(capTable CapTable, ownerId string) (string, error) {
	ownerWallets, err := FindWallet(ownerId)
	if err != nil {
		return "", err
	}

	for _, ownerWallet := range ownerWallets {
		for _, tokenHolder := range capTable.TokenHolders {
			if ownerWallet.Owner == tokenHolder.Owner {
				return tokenHolder.Balances[0].Amount, nil
			}
		}
	}

	return "", errors.New("owner does not have any shares in this company")
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
	captable = filterOutOwnerWithNoShares(captable)
	for i, tokenHolder := range captable.TokenHolders {
		wallet, err := FindWalletByWalletAddress(tokenHolder.Address)
		if err != nil {
			logrus.Error(err)
			return CapTable{}, err
		}

		captable.TokenHolders[i].Owner = wallet.Owner
	}

	return captable, nil
}
