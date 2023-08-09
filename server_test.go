package main

import (
	"brok/navnetjener/model"
	"brok/navnetjener/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

	testWallet := utils.CreateTestWallet()

	testWallet.Save()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/wallet/"+testWallet.WalletAddress, nil)
	router.ServeHTTP(w, req)

	var receivedWallet model.PublicWalletInfo
	json.Unmarshal([]byte(w.Body.String()), &receivedWallet)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testWallet.WalletAddress, receivedWallet.WalletAddress)
}

func TestApiShouldCreateNewWalletEntryInDatabase(t *testing.T) {
	router := setup()

	testWallet := utils.CreateTestWallet()

	json, _ := json.Marshal(testWallet)
	req, _ := http.NewRequest("POST", "/wallet", bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	storedWallet, _ := model.FindWalletByWalletAddress(testWallet.WalletAddress)

	assert.Equal(t, storedWallet.WalletAddress, testWallet.WalletAddress)

}
