package model

import (
	"errors"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeString(input string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(input)
}

func SanitizeWallet(w *Wallet) error {
	p := bluemonday.StrictPolicy()

	// Sanitize string fields
	w.FirstName = p.Sanitize(w.FirstName)
	w.LastName = p.Sanitize(w.LastName)
	w.Pnr = p.Sanitize(w.Pnr)
	w.BirthDate = p.Sanitize(w.BirthDate)
	w.WalletAddress = p.Sanitize(w.WalletAddress)

	// Validate Ethereum address
	if err := validateEthereumAddress(w.WalletAddress); err != nil {
		return err
	}

	return nil
}

func validateEthereumAddress(address string) error {
	// Check for length and starting characters
	if len(address) != 42 || address[:2] != "0x" {
		return errors.New("invalid wallet address format")
	}

	// Ensure the remaining characters are hexadecimal
	matched, err := regexp.MatchString(`^0x[a-fA-F0-9]{40}$`, address)
	if err != nil || !matched {
		return errors.New("invalid wallet address format")
	}

	return nil
}
