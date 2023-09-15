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
	"brok/navnetjener/api"
	"brok/navnetjener/model"
	"brok/navnetjener/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var API_VERSION string = "/v1"

func TestApiShouldReturnOneWalletWithCorrectWalletAddress(t *testing.T) {

	router := utils.Setup()
	router.GET("/v1/wallet/:walletAddress", api.GetWalletByWalletAddress)

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
	router := utils.Setup()

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
	router := utils.Setup()

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
	router := utils.Setup()

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
	router := utils.Setup()

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
	dateBorn := utils.RandomNumber(0, 30)
	mountBorn := utils.RandomNumber(1, 12)
	yearBorn := utils.RandomNumber(68, 99)

	birthDate := dateBorn + mountBorn + yearBorn

	return model.Wallet{
		OwnerPersonFirstName: faker.FirstNameFemale(),
		OwnerPersonLastName:  faker.LastName(),
		CapTableOrgnr:        utils.RandomOrgnr(),
		OwnerPersonFnr:       birthDate + "00000",
		OwnerPersonBirthDate: birthDate,
		WalletAddress:        utils.RandomWalletAddress(),
	}
}
