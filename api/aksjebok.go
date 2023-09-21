package api

import (
	"brok/navnetjener/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetForetakByOrgnr(context *gin.Context) {
	orgnr := context.Param("orgnr")
	safeOrgnr := model.SanitizeString(orgnr)

	if len(safeOrgnr) != 9 {
		context.JSON(http.StatusBadRequest, gin.H{"error": safeOrgnr + " must be 9 valid digits"})
		return
	}

	if _, err := strconv.Atoi(safeOrgnr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": safeOrgnr + " must be 9 valid digits"})
		return
	}

	capTable, err := model.FindCaptableByOrgnr(safeOrgnr)
	if capTable.Name == "" || err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke aksjebok for orgnr " + safeOrgnr})
		return
	}

	context.JSON(http.StatusOK, capTable)
}

func GetNumberOfSharesForOwnerOfAForetak(context *gin.Context) {
	capTableOrgnr := context.Param("orgnr")
	safeCapTableOrgnr := model.SanitizeString(capTableOrgnr)

	if len(safeCapTableOrgnr) != 9 {
		context.JSON(http.StatusBadRequest, gin.H{"error": safeCapTableOrgnr + " must be 9 valid digits"})
		return
	}

	if _, err := strconv.Atoi(safeCapTableOrgnr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": safeCapTableOrgnr + " must be 9 valid digits"})
		return
	}

	capTable, err := model.FindCaptableByOrgnr(safeCapTableOrgnr)
	if capTable.Name == "" || err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke aksjebok for orgnr " + safeCapTableOrgnr})
		return
	}

	id := context.Param("id")
	safeId := model.SanitizeString(id)

	if len(safeId) != 11 && len(safeId) != 9 {
		logrus.Warn("error: id må være et gyldig orgnr eller fnr")
		context.JSON(http.StatusBadRequest, gin.H{"error": safeId + " må være et gyldig orgnr eller fnr"})
		return
	}

	if _, err := strconv.Atoi(safeId); err != nil {
		logrus.Warn("error: id må være et tall")
		context.JSON(http.StatusBadRequest, gin.H{"error": safeId + " må være et tall"})
		return
	}

	numberOfShares, err := model.FindNumberOfSharesForOwnerOfCaptable(capTable, id)
	if err != nil {
		logrus.Warn(err)
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke noen aksjer for " + safeId + " i foretak " + safeCapTableOrgnr})
		return
	}

	context.JSON(http.StatusOK, numberOfShares)
}

// GetForetak returns 25 foretak from TheGraph and the database
// user can use the queryparameter "page" to paginate
func GetForetak(context *gin.Context) {
	page := context.Query("page")
	safePage := model.SanitizeString(page)

	// convert page to int
	// if page is empty, set it to 0
	// if page is not a number, set it to 0
	// if page is less than 0, set it to 0
	if safePage == "" {
		safePage = "0"
	}

	safePageInt, err := strconv.Atoi(safePage)

	if err != nil {
		safePageInt = 0
	}

	capTables, err := model.FindForetak(safePageInt)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke noen aksjebøker"})
		return
	}

	context.JSON(http.StatusOK, capTables)
}

// Used for bulk lookup
type WalletInfo struct {
	Identifier    string  `json:"identifier"`
	WalletAddress *string `json:"walletAddress"`
}

type BulkLookupRequest struct {
	Identifiers []string `json:"identifiers"`
}

type BulkLookupResponse struct {
	Wallets []WalletInfo `json:"wallets"`
}

func GetOwnersForForetak(context *gin.Context) {
	parentOrgnr := context.Param("orgnr")
	safeParentOrgnr := model.SanitizeString(parentOrgnr)

	var req BulkLookupRequest

	// Parse the JSON body
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Initialize the response object
	response := make([]WalletInfo, len(req.Identifiers))

	logrus.Info("Request: ", req)

	// Loop through identifiers
	for i, id := range req.Identifiers {
		var walletInfo WalletInfo
		walletInfo.Identifier = id // set identifier

		var err error
		var walletAddress string // initialize walletAddress as an empty string

		// Check if identifier is for person or foretak based on its length
		if len(id) == 11 { // It's a fødselsnummer
			walletAddress, err = model.FindWalletByPersonFnrAndParentOrg(id, safeParentOrgnr)
		} else if len(id) == 9 { // It's an orgnr
			walletAddress, err = model.FindWalletByOrgnrAndParentOrg(id, safeParentOrgnr)
		} else {
			// Invalid identifier, skip to the next
			continue
		}

		if err != nil {
			logrus.Error("Error finding wallet: ", err) // Log the error
			continue
		}

		if walletAddress != "" {
			walletInfo.WalletAddress = &walletAddress // set wallet address
		}

		response[i] = walletInfo // populate the response array at index i
	}

	// Return the final response
	context.JSON(http.StatusOK, gin.H{"wallets": response})
}
