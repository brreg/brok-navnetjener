package api

import (
	"brok/navnetjener/model"
	"brok/navnetjener/utils"
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
			wallet.OwnerPersonBirthYear = utils.FindBirthYear(wallet.OwnerPersonFnr)
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
