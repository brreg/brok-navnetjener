package api_test

import (
	"brok/navnetjener/api"
	"brok/navnetjener/model"
	"brok/navnetjener/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOneForetak(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/:orgnr", api.GetForetakByOrgnr)

	orgnr := "815493000"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/"+orgnr, nil)
	router.ServeHTTP(w, req)

	var receivedCapTable model.CapTable
	json.Unmarshal([]byte(w.Body.String()), &receivedCapTable)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, 6, len(receivedCapTable.TokenHolders))
}

func TestFindOneForetakOnFirstPage(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/", api.GetForetak)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/?page=0", nil)
	router.ServeHTTP(w, req)

	var receivedCapTable []model.CapTable
	json.Unmarshal([]byte(w.Body.String()), &receivedCapTable)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Greater(t, len(receivedCapTable), 0)
}

func TestFindZeroForetakOnSecondPage(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/", api.GetForetak)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/?page=1", nil)
	router.ServeHTTP(w, req)

	var receivedCapTable []model.CapTable
	json.Unmarshal([]byte(w.Body.String()), &receivedCapTable)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, 0, len(receivedCapTable))
}

func TestAmountOfSharesForOwnerOfCaptable(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/:orgnr/balanse/:id", api.GetNumberOfSharesForOwnerOfAForetak)

	orgnr := "815493000"
	ownerId := "310767859"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/"+orgnr+"/balanse/"+ownerId, nil)
	router.ServeHTTP(w, req)

	var receivedNumberOfShares string
	json.Unmarshal([]byte(w.Body.String()), &receivedNumberOfShares)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "1000", receivedNumberOfShares)
}

func TestAmountOfSharesForOwnerOfCaptable2(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/:orgnr/balanse/:id", api.GetNumberOfSharesForOwnerOfAForetak)

	orgnr := "815493000"
	ownerId := "15097600002"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/"+orgnr+"/balanse/"+ownerId, nil)
	router.ServeHTTP(w, req)

	var receivedNumberOfShares string
	json.Unmarshal([]byte(w.Body.String()), &receivedNumberOfShares)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "200", receivedNumberOfShares)
}

func TestShouldNotFindAnySharesForOwner(t *testing.T) {
	router := utils.Setup()
	router.GET("/aksjebok/:orgnr/balanse/:id", api.GetNumberOfSharesForOwnerOfAForetak)

	orgnr := "815493000"
	ownerId := "815493009"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/aksjebok/"+orgnr+"/balanse/"+ownerId, nil)
	router.ServeHTTP(w, req)

	var receivedNumberOfShares string
	json.Unmarshal([]byte(w.Body.String()), &receivedNumberOfShares)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetOwnersForForetak(t *testing.T) {
	// Initialize router and other setup
	router := utils.Setup()
	router.POST("/aksjebok/:orgnr/aksjeeier", api.GetOwnersForForetak)

	// Prepare the request body
	reqBody := api.BulkLookupRequest{
		Identifiers: []string{"12345678901", "123456789"}, // Mock fødselsnummer and orgnr
	}

	parentOrgnr := "123456789"

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/aksjebok/"+parentOrgnr+"/aksjeeier", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Initialize the recorder
	w := httptest.NewRecorder()

	// Perform the test
	router.ServeHTTP(w, req)

	// Check if the status code is what you expect
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response
	var response api.BulkLookupResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Perform your assertions here based on what you expect `response` to contain
	// For example:
	assert.NotNil(t, response.Wallets)
	// More detailed tests can be done based on your specific requirements
}
