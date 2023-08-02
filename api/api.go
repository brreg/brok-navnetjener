package api

import (
	"brok/navnetjener/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWalletByID(context *gin.Context) {
	walletAddress := context.Param("walletAddress")

	wallet, err := model.FindWalletByAddress(walletAddress)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if wallet == (model.PublicWalletInfo{}) {
		context.JSON(http.StatusNotFound, gin.H{"error": "no wallet with address " + walletAddress})
		return
	}

	context.JSON(http.StatusOK, wallet)
}

func CreateWallet(context *gin.Context) {
	var input model.Wallet

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse body"})
		return
	}

	savedWallet, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not store in database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"wallet": savedWallet})
}

func GetAllWallets(context *gin.Context) {
	wallets, err := model.FindAllWallets()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, wallets)
}
