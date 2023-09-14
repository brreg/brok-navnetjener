package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"brok/navnetjener/api"
	"brok/navnetjener/model"

	"brok/navnetjener/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

// START OF COPYPASTE FROM PACKAGE MAIN
// ----------------- NOTE TO LEAD DEVELOPER -----------------
//
// The following block of code is copied over from the main package to enable testing.
// This is not an ideal practice and poses several limitations:
//
//  1. Code Duplication: The same piece of logic resides in two places, which violates DRY (Don't Repeat Yourself) principles.
//     This leads to increased maintenance overhead.
//
//  2. Test-Production Parity: Since the logic is duplicated, there's a risk that changes in the main code may not get reflected in tests immediately,
//     leading to false positives/negatives.
//
// 3. Coupling: Having most of the logic in the `main` package increases coupling, making it hard to write isolated unit tests for individual components.
//
// TODO: Recommended Action:
//
// Refactor the logic in server.go or other parts of the `main` package into separate, well-defined packages. This would allow:
//
// - Easier testing by enabling these packages to be imported in test files.
// - Better separation of concerns, which in turn makes the codebase easier to understand and maintain.
// - Greater reusability of code components in future projects.
//
// By doing so, we'll adhere to best practices in Go development and ensure that our code is both clean and maintainable.
//
// ----------------------------------------------------------------
func setup() *gin.Engine {
	loadEnv()
	loadDatabase()
	gin.SetMode(gin.TestMode)
	return routerConfig()
}

func loadEnv() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Print("Error loading .env.local file")
		panic(err)
	}
}

func loadDatabase() {
	database.Connect()
	autoMigrate := os.Getenv("DB_AUTO_MIGRATE")
	if autoMigrate == "true" {
		logrus.Info("Auto migrating database")
		database.Database.AutoMigrate(&model.Wallet{})
	}
}

func routerConfig() *gin.Engine {

	env, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		logrus.Warn("ENVIRONMENT environment variable not set, using default value: dev")
		env = "development"
	}

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if env == "development" {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	/*
		Set a limit on the entire request body to 1 KiB
		Limit calculation (in bytes):
			257 (FirstName)
			257 (LastName)
			20 (Orgnr)
			13 (Fnr)
			6 (BirthDate)
			44 (WalletAddress)
			150 (JSON overhead)

		Total: 747, rounding up to 1024 to have some wiggle room
	*/
	router.Use(maxBodySize(1024)) // 1 KiB limit

	// Group API endpoints in /v1
	v1 := router.Group("/v1")

	v1.POST("/wallets/bulk", api.GetWalletsForIdentifiers)

	return router
}

func maxBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}

/*
* END COPYPASTE
 */

func TestGetWalletsForIdentifiers(t *testing.T) {
	// Initialize router and other setup
	router := setup()

	// Prepare the request body
	reqBody := api.BulkLookupRequest{
		Identifiers: []string{"12345678901", "123456789"}, // Mock fødselsnummer and orgnr
		ParentOrgnr: "123456789",                          // Mock parent orgnr
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/v1/wallets/bulk", bytes.NewBuffer(reqBodyBytes))
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
