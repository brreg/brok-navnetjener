package main

import (
	"brok/navnetjener/api"
	"brok/navnetjener/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := utils.Setup()
	apiRoutes(router)
	serveApplication(router)
}

func apiRoutes(router *gin.Engine) {

	// Group API endpoints in /v1
	v1 := router.Group("/v1")

	v1.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)
	v1.POST("/wallet", api.CreateWallet)
	v1.POST("/wallets/bulk", api.GetWalletsForIdentifiers)

	v1.GET("/person/:fnr", api.GetAllForetakForPerson)

	v1.GET("/foretak/:orgnr", api.GetForetakByOrgnr)
	v1.GET("/foretak", api.GetForetak)
}

func serveApplication(router *gin.Engine) {
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		logrus.Warn("SERVER_PORT environment variable not set, using default port 8080")
		port = "8080"
	}

	router.Run(":" + port)
}
