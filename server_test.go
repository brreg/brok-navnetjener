package main

import (
	"brok/navnetjener/model"
	"bytes"
	"crypto/rand"
	"encoding/hex"
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

func TestWalletCreation(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	testWallet := model.Wallet{
		FirstName:            "Kari",
		LastName:             "Norman",
		SosialSecurityNumber: "01019844422",
		WalletAddress:        randomWalletAddress(),
	}
	json, _ := json.Marshal(testWallet)
	req, _ := http.NewRequest("POST", "/wallet", bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	storedWallet, _ := model.FindWalletByAddress(testWallet.WalletAddress)

	assert.Equal(t, storedWallet.WalletAddress, testWallet.WalletAddress)

}

func TestWalletAddressDoesNotExist(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/wallet/wallet2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func randomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}
