package main

/**

!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----!
|																																																|
| 	Dette settet med tester fungerer KUN når en lokal database med navnet navnetjener kjører,		|
| 	OG en lokal hardhat kjede med TheGraph som er startet fra https://github.com/brreg/brok!		|
|																																																|
|-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----!

*/

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

func TestApiShouldFindAllCapTablesBelongingToPerson(t *testing.T) {
	// Setup
	router := setup()

	// Test
	req, _ := http.NewRequest("GET", API_VERSION+"/person/21058000000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var receivedDataFromApi []model.CapTable
	if err := json.Unmarshal(w.Body.Bytes(), &receivedDataFromApi); err != nil {
		panic(err.Error())
	}
	assert.Equal(t, 1, len(receivedDataFromApi))

	assert.Equal(t, "Ryddig Bobil AS", receivedDataFromApi[0].Name)
}

func TestApiShouldGiveAllShareholdersForCompany(t *testing.T) {
	// Setup
	router := setup()

	// Test
	req, _ := http.NewRequest("GET", API_VERSION+"/foretak/815493000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var receivedDataFromApi model.CapTable
	if err := json.Unmarshal(w.Body.Bytes(), &receivedDataFromApi); err != nil {
		panic(err.Error())
	}

	assert.Equal(t, len(receivedDataFromApi.TokenHolders), 6)

	for _, shareholder := range receivedDataFromApi.TokenHolders {
		owner, _ := model.FindWalletByWalletAddress(shareholder.Address)
		assert.Equal(t, shareholder.Owner, owner.Owner)
	}

}

func TestCreateWalletWithToLargeRequestBody(t *testing.T) {
	router := setup()

	// Generate a large payload using fields in model.Wallet
	largeName := strings.Repeat("a", 1024*2) // 2KiB

	wallet := map[string]interface{}{
		"owner_person_first_name": largeName,
		"owner_person_last_name":  "Doe",
		"cap_table_orgnr":         12345678,
		"owner_person_fnr":        "200501021234",
		"wallet_address":          "0x1234567890abcdef",
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

func CreateTestWallet() model.Wallet {
	dateBorn := randomNumber(0, 30)
	mountBorn := randomNumber(1, 12)
	yearBorn := randomNumber(68, 99)

	birthDate := dateBorn + mountBorn + yearBorn

	return model.Wallet{
		OwnerPersonFirstName: faker.FirstNameFemale(),
		OwnerPersonLastName:  faker.LastName(),
		CapTableOrgnr:        randomOrgnr(),
		OwnerPersonFnr:       birthDate + "00000",
		OwnerPersonBirthDate: birthDate,
		WalletAddress:        randomWalletAddress(),
	}
}

func randomNumber(min int, max int) string {
	random := r.New(r.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%02d", random.Intn(max-min)+min)
}

func randomOrgnr() string {
	return randomNumber(111111111, 999999999)
}

func randomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}
