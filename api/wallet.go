package api

import (
	"brok/navnetjener/model"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var newWallet model.Wallet

	if err := context.ShouldBindJSON(&newWallet); err != nil {
		// Check if the error is because of large request body
		if err.Error() == "http: request body too large" {
			context.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Request body exceeds limit"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse body"})
		return
	}

	if newWallet.OwnerPersonPnr != "" {
		newWallet.OwnerPersonBirthDate = newWallet.OwnerPersonPnr[:6]
	}

	savedWallet, err := newWallet.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not store in database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"wallet": savedWallet})
}
