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
