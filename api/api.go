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

func GetWalletByOrgnr(context *gin.Context) {
	orgnr := context.Param("orgnr")

	wallets, err := model.FindWalletByOrgnr(orgnr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(wallets) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "no wallet found"})
		return
	}

	context.JSON(http.StatusOK, wallets)
}

func GetWalletByPnr(context *gin.Context) {
	pnr := context.Param("pnr")
	var response []PersonResponse

	wallets, err := model.FindWalletByPnr(pnr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(wallets) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "no wallet found"})
		return
	}

	for _, wallet := range wallets {
		response = append(response, PersonResponse{
			WalletAddress: wallet.WalletAddress,
		})
	}

	context.JSON(http.StatusOK, response)
}

func GetWalletByWalletAddress(context *gin.Context) {
	walletAddress := context.Param("walletAddress")

	wallet, err := model.FindWalletByWalletAddress(walletAddress)
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
	var newWallet model.Wallet

	if err := context.ShouldBindJSON(&newWallet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse body"})
		return
	}

	newWallet.YearBorn = newWallet.Pnr[4:6]

	savedWallet, err := newWallet.Save()
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
