package main

import (
	"brok/navnetjener/model"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	r "math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setup() *gin.Engine {
	loadEnv()
	loadDatabase()
	gin.SetMode(gin.TestMode)
	return routerConfig()
}

var API_VERSION string = "/v1"

func TestWalletRoute(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", API_VERSION+"/wallet", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiShouldReturnOneWalletWithCorrectWalletAddress(t *testing.T) {
	router := setup()

	testWallet := CreateTestWallet()

	testWallet.Save()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", API_VERSION+"/wallet/"+testWallet.WalletAddress, nil)
	router.ServeHTTP(w, req)

	var receivedWallet model.PublicWalletInfo
	json.Unmarshal([]byte(w.Body.String()), &receivedWallet)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testWallet.WalletAddress, receivedWallet.WalletAddress)
}

func TestApiShouldCreateNewWalletEntryInDatabase(t *testing.T) {
	router := setup()

	testWallet := CreateTestWallet()

	json, _ := json.Marshal(testWallet)
	req, _ := http.NewRequest("POST", API_VERSION+"/wallet", bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	storedWallet, _ := model.FindWalletByWalletAddress(testWallet.WalletAddress)

	assert.Equal(t, storedWallet.WalletAddress, testWallet.WalletAddress)

}

// func TestShouldFindAllWalletsBelongingToPerson(t *testing.T) {
// 	// Setup
// 	router := setup()

// 	// testWallets := CreateFiveTestWalletsForOnePerson()
// 	// for _, wallet := range testWallets {
// 	// 	wallet.Save()
// 	// }

// 	// Test
// 	req, _ := http.NewRequest("GET", API_VERSION+"/person/"+testWallets[0].Pnr, nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var receivedWallet []struct {
// 		WalletAddress string `json:"wallet_address"`
// 	}
// 	if err := json.Unmarshal(w.Body.Bytes(), &receivedWallet); err != nil {
// 		panic(err.Error())
// 	}
// 	assert.Equal(t, len(receivedWallet), 5)

// 	for i, wallet := range receivedWallet {
// 		assert.Equal(t, testWallets[i].WalletAddress, wallet.WalletAddress)
// 	}

// }

// func TestShouldFindAllShareholderForCompany(t *testing.T) {
// 	// Setup
// 	router := setup()

// 	testWallets := CreateSevenTestWalletsForOneCompany()
// 	for _, wallet := range testWallets {
// 		wallet.Save()
// 	}

// 	// Test
// 	req, _ := http.NewRequest("GET", API_VERSION+"/company/"+fmt.Sprint(testWallets[0].Orgnr), nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var receivedWallet []model.PublicWalletInfo
// 	if err := json.Unmarshal(w.Body.Bytes(), &receivedWallet); err != nil {
// 		panic(err.Error())
// 	}
// 	assert.Equal(t, len(receivedWallet), 7)

// 	for i, wallet := range receivedWallet {
// 		assert.Equal(t, testWallets[i].WalletAddress, wallet.WalletAddress)
// 	}

// }

func TestCreateWalletWithToLargeRequestBody(t *testing.T) {
	router := setup()

	// Generate a large payload using fields in model.Wallet
	largeName := strings.Repeat("a", 1024*2) // 2KiB

	wallet := map[string]interface{}{
		"first_name":     largeName,
		"last_name":      "Doe",
		"orgnr":          12345678,
		"pnr":            "200501021234",
		"wallet_address": "0x1234567890abcdef",
	}
	payloadBytes, err := json.Marshal(wallet)
	if err != nil {
		t.Fatalf("Failed to marshal wallet: %v", err)
	}

	req, _ := http.NewRequest(http.MethodPost, API_VERSION+"/wallet", bytes.NewReader(payloadBytes))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// Since Gin has a 1KiB request body limit, it should return a 413 Payload Too Large.
	assert.Equal(t, http.StatusRequestEntityTooLarge, resp.Code)
}

func CreateSevenTestWalletsForOneCompany() []model.Wallet {
	orgnr := randomNumberInt(11111111, 99999999) // Use 8 digits orgnr for testing

	var wallets []model.Wallet

	for i := 0; i < 7; i++ {
		dateBorn := randomNumber(1, 30)
		mountBorn := randomNumber(1, 12)
		yearBorn := randomNumber(68, 99)
		birthDate := dateBorn + mountBorn + yearBorn

		wallets = append(wallets, model.Wallet{
			FirstName:     faker.FirstNameFemale(),
			LastName:      faker.LastName(),
			Orgnr:         orgnr,
			Pnr:           birthDate + "00000",
			BirthDate:     birthDate,
			WalletAddress: randomWalletAddress(),
		})
	}

	return wallets
}

func CreateFiveTestWalletsForOnePerson() []model.Wallet {
	firstName := faker.FirstNameFemale()
	lastName := faker.LastName()

	dateBorn := randomNumber(0, 30)
	mountBorn := randomNumber(1, 12)
	yearBorn := randomNumber(68, 99)

	birthDate := dateBorn + mountBorn + yearBorn

	var wallets []model.Wallet

	for i := 0; i < 5; i++ {
		wallets = append(wallets, model.Wallet{
			FirstName:     firstName,
			LastName:      lastName,
			Orgnr:         randomNumberInt(11111111, 99999999), // Use 8 digits orgnr for testing
			Pnr:           birthDate + "00000",
			BirthDate:     birthDate,
			WalletAddress: randomWalletAddress(),
		})
	}

	return wallets
}

func CreateTestWallet() model.Wallet {
	dateBorn := randomNumber(0, 30)
	mountBorn := randomNumber(1, 12)
	yearBorn := randomNumber(68, 99)

	birthDate := dateBorn + mountBorn + yearBorn

	return model.Wallet{
		FirstName:     faker.FirstNameFemale(),
		LastName:      faker.LastName(),
		Orgnr:         randomNumberInt(11111111, 99999999), // Use 8 digits orgnr for testing
		Pnr:           birthDate + "00000",
		BirthDate:     birthDate,
		WalletAddress: randomWalletAddress(),
	}
}

func randomNumber(min int, max int) string {
	random := r.New(r.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%02d", random.Intn(max-min)+min)
}

func randomNumberInt(min int, max int) int {
	random := r.New(r.NewSource(time.Now().UnixNano()))
	return random.Intn(max-min) + min
}

func randomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}
