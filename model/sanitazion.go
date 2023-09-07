package model

import (
	"github.com/mrz1836/go-sanitize"
)

// Only allow AlphaNumeric character's to return
func SanitizeString(input string) string {
	return sanitize.AlphaNumeric(input, true)
}

func SanitizeWallet(w *Wallet) {

	// Sanitize string fields
	w.OwnerPersonFirstName = SanitizeString(w.OwnerPersonFirstName)
	w.OwnerPersonLastName = SanitizeString(w.OwnerPersonLastName)
	w.OwnerPersonPnr = SanitizeString(w.OwnerPersonPnr)
	w.OwnerPersonBirthDate = SanitizeString(w.OwnerPersonBirthDate)
	w.OwnerCompanyName = SanitizeString(w.OwnerCompanyName)
	w.OwnerCompanyOrgnr = SanitizeString(w.OwnerCompanyOrgnr)
	w.WalletAddress = SanitizeString(w.WalletAddress)
}
