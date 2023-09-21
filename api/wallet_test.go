package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"brok/navnetjener/api"
	"brok/navnetjener/model"
	"brok/navnetjener/utils"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatingMultipleWallets(t *testing.T) {
	// Initialize router and other setup
	router := utils.Setup()
	router.POST("/v1/wallet/", api.CreateWallet)

	// Prepare the request body
	reqBody := []model.Wallet{}
	for i := 1; i < 10; i++ {
		wallet := model.Wallet{
			OwnerPersonFirstName: faker.FirstName(),
			OwnerPersonLastName:  faker.LastName(),
			OwnerPersonFnr:       utils.RandomFnr(),
			CapTableOrgnr:        utils.RandomOrgnr(),
			WalletAddress:        utils.RandomWalletAddress(),
		}
		reqBody = append(reqBody, wallet)
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/v1/wallet/", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Initialize the recorder
	w := httptest.NewRecorder()

	// Perform the test
	router.ServeHTTP(w, req)

	// Check if the status code is what you expect
	assert.Equal(t, http.StatusCreated, w.Code)
}
