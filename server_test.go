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
	"testing"

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

func TestWalletRoute(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/wallet", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiShouldReturnOneWalletWithCorrectWalletAddress(t *testing.T) {
	router := setup()

	testWallet := createTestWallet()

	testWallet.Save()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/wallet/"+testWallet.WalletAddress, nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Body)
	var receivedWallet model.PublicWalletInfo
	json.Unmarshal([]byte(w.Body.String()), &receivedWallet)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testWallet.WalletAddress, receivedWallet.WalletAddress)
}

func TestApiShouldCreateNewWalletEntryInDatabase(t *testing.T) {
	router := setup()

	testWallet := createTestWallet()

	json, _ := json.Marshal(testWallet)
	req, _ := http.NewRequest("POST", "/wallet", bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	storedWallet, _ := model.FindWalletByAddress(testWallet.WalletAddress)

	assert.Equal(t, storedWallet.WalletAddress, testWallet.WalletAddress)

}

func randomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}

func createTestWallet() model.Wallet {
	return model.Wallet{
		FirstName:            faker.FirstNameFemale(),
		LastName:             faker.LastName(),
		CompanyOrgnr:         randomNumber(11111111, 99999999), // Use 8 digits orgnr for testing
		SosialSecurityNumber: randomNumber(0, 30) + randomNumber(1, 12) + randomNumber(78, 99) + "00000",
		WalletAddress:        randomWalletAddress(),
	}
}

func randomNumber(min int, max int) string {
	return fmt.Sprintf("%02d", r.Intn(max-min)+min)
}
