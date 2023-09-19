package api

import (
	"brok/navnetjener/database"
	"brok/navnetjener/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

// List out all wallets belonging to a person
type PersonResponse struct {
	WalletAddress string `json:"wallet_address"`
}

func GetWalletByWalletAddress(context *gin.Context) {
	walletAddress := context.Param("walletAddress")

	wallet, err := model.FindWalletByWalletAddress(walletAddress)
	if wallet == (model.PublicWalletInfo{}) || err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "no wallet with address " + walletAddress})
		return
	}

	context.JSON(http.StatusOK, wallet)
}

func CreateWallet(context *gin.Context) {
	var newWallet []model.Wallet

	if err := context.ShouldBindJSON(&newWallet); err != nil {
		// Check if the error is because of large request body
		if err.Error() == "http: request body too large" {
			context.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Request body exceeds limit"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse body"})
		return
	}

	var savedWalletList []*model.Wallet

	for _, wallet := range newWallet {
		if wallet.OwnerPersonFnr != "" {
			wallet.OwnerPersonBirthDate = wallet.OwnerPersonFnr[:6]
		}
		savedWallet, err := wallet.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Could not store in database"})
			return
		}
		savedWalletList = append(savedWalletList, savedWallet)
	}

	context.JSON(http.StatusCreated, gin.H{"wallet": savedWalletList})
}

// Used for bulk lookup
type WalletInfo struct {
	Identifier    string  `json:"identifier"`
	WalletAddress *string `json:"walletAddress"`
}

type BulkLookupRequest struct {
	Identifiers []string `json:"identifiers"`
	ParentOrgnr string   `json:"parentOrgnr"`
}

type BulkLookupResponse struct {
	Wallets []WalletInfo `json:"wallets"`
}

func FindWalletByPersonFnrAndParentOrg(fnr string, parentOrgnr string) (string, error) {
	var wallet model.Wallet
	safeFnr := model.SanitizeString(fnr)
	safeParentOrgnr := model.SanitizeString(parentOrgnr)

	logrus.Info("Sanitized fnr: ", safeFnr, ", Sanitized parentOrgnr: ", safeParentOrgnr) // Log sanitized inputs

	err := database.Database.Where("owner_person_fnr=? AND cap_table_orgnr=?", safeFnr, safeParentOrgnr).First(&wallet).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.Info("Record not found") // Log record not found
			return "", nil                  // return nil error when record is not found
		}
		logrus.Error("Could not find wallet in db with fnr: ", fnr, " and parentOrgnr: ", parentOrgnr)
		return "", err
	}

	logrus.Info("Found wallet: ", wallet.WalletAddress) // Log the found wallet

	return wallet.WalletAddress, nil
}

func GetWalletsForIdentifiers(context *gin.Context) {
	var req BulkLookupRequest // Antatt å være definert et annet sted

	// Parse the JSON body
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	logrus.Info("Parsed ParentOrgnr: ", req.ParentOrgnr)

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
			walletAddress, err = model.FindWalletByPersonFnrAndParentOrg(id, req.ParentOrgnr)
		} else if len(id) == 9 { // It's an orgnr
			walletAddress, err = model.FindWalletByOrgnrAndParentOrg(id, req.ParentOrgnr)
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
