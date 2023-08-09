package utils

import (
	"brok/navnetjener/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	r "math/rand"

	"github.com/bxcodec/faker/v3"
)

func CreateFiveTestWalletsForOnePerson() []model.Wallet {
	firstName := faker.FirstNameFemale()
	lastName := faker.LastName()

	dateBorn := randomNumber(0, 30)
	mountBorn := randomNumber(1, 12)
	yearBorn := randomNumber(68, 99)

	pnr := fmt.Sprintf("%v%v%v00000", dateBorn, mountBorn, yearBorn)

	var wallets []model.Wallet

	for i := 0; i < 5; i++ {
		wallets = append(wallets, model.Wallet{
			FirstName:     firstName,
			LastName:      lastName,
			Orgnr:         randomNumber(11111111, 99999999), // Use 8 digits orgnr for testing
			Pnr:           pnr,
			YearBorn:      string(yearBorn),
			WalletAddress: randomWalletAddress(),
		})
	}

	return wallets
}

func CreateTestWallet() model.Wallet {
	dateBorn := randomNumber(0, 30)
	mountBorn := randomNumber(1, 12)
	yearBorn := randomNumber(68, 99)
	return model.Wallet{
		FirstName:     faker.FirstNameFemale(),
		LastName:      faker.LastName(),
		Orgnr:         randomNumber(11111111, 99999999), // Use 8 digits orgnr for testing
		Pnr:           string(dateBorn) + string(mountBorn) + string(yearBorn) + "00000",
		YearBorn:      string(yearBorn),
		WalletAddress: randomWalletAddress(),
	}
}

func randomNumber(min int, max int) string {
	return fmt.Sprintf("%02d", r.Intn(max-min)+min)
}

func randomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}
