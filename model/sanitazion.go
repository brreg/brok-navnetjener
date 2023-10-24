package model

import (
	"regexp"
)

var (
	alphaNumericWithSpacesRegExp = regexp.MustCompile(`[^a-zA-Z0-9æøåÆØÅ\s]`)
)

// emptySpace is an empty space for replacing
var emptySpace = []byte("")

// Only allow AlphaNumeric character's to return
func SanitizeString(input string) string {
	return string(alphaNumericWithSpacesRegExp.ReplaceAll([]byte(input), emptySpace))
}

func SanitizeWallet(w *Wallet) {

	// Sanitize string fields
	w.OwnerPersonFirstName = SanitizeString(w.OwnerPersonFirstName)
	w.OwnerPersonLastName = SanitizeString(w.OwnerPersonLastName)
	w.OwnerPersonFnr = SanitizeString(w.OwnerPersonFnr)
	w.OwnerPersonBirthYear = SanitizeString(w.OwnerPersonBirthYear)
	w.OwnerCompanyName = SanitizeString(w.OwnerCompanyName)
	w.OwnerCompanyOrgnr = SanitizeString(w.OwnerCompanyOrgnr)
	w.WalletAddress = SanitizeString(w.WalletAddress)
}
