package utils

import (
	"brok/navnetjener/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	r "math/rand"

	"github.com/bxcodec/faker/v3"
)

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
