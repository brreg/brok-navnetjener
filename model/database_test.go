package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWalletToPublicInfo(t *testing.T) {
	personWallet := Wallet{
		OwnerPersonFirstName: "firstName",
		OwnerPersonLastName:  "lastName",
		OwnerPersonPnr:       "12345678901",
		OwnerPersonBirthDate: "123456",
		OwnerCompanyName:     "",
		OwnerCompanyOrgnr:    "",
	}

	publicPersonWallet := parseWalletToPublicInfo(personWallet)

	assert.Equal(t, publicPersonWallet.Owner.Person.FirstName, personWallet.OwnerPersonFirstName)
	assert.Equal(t, publicPersonWallet.Owner.Person.LastName, personWallet.OwnerPersonLastName)
	assert.Equal(t, publicPersonWallet.Owner.Person.BirthDate, personWallet.OwnerPersonBirthDate)
	assert.Equal(t, publicPersonWallet.Owner.Company.Name, personWallet.OwnerCompanyName)
	assert.Equal(t, publicPersonWallet.Owner.Company.Orgnr, personWallet.CapTableOrgnr)
}
