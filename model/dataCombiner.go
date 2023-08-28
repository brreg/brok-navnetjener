package model

// findCaptableByOrgnr combines data from TheGraph and the database
// to return a CapTable struct with person data
func FindCaptableByOrgnr(orgnr string) (CapTable, error) {
	captable, err := FindCaptableByOrgnrFromTheGraph(orgnr)
	if err != nil {
		return CapTable{}, err
	}

	for i, tokenHolder := range captable.TokenHolders {
		person, err := findPersonByWalletAddress(tokenHolder.Address)
		if err != nil {
			return CapTable{}, err
		}

		captable.TokenHolders[i].Person = person
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
		for j, tokenHolder := range captable.TokenHolders {
			person, err := findPersonByWalletAddress(tokenHolder.Address)
			if err != nil {
				return nil, err
			}

			captables[i].TokenHolders[j].Person = person
		}
	}

	return captables, nil
}
