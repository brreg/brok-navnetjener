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

	limit, exists := os.LookupEnv("LIMIT_ENDPOINTS")
	if !exists {
		logrus.Warn("LIMIT_ENDPOINTS environment variable not set, using default false")
		limit = "false"
	}

	// Don't use privileged endpoints if feature toggle LIMIT_ENDPOINTS is set
	if limit == "false" {
		v1.POST("/wallet", api.CreateWallet)
		v1.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)

		v1.GET("/aksjeeier/:id", api.GetAllForetakForAksjeeier)
		v1.GET("/aksjebok/:orgnr/balanse/:id", api.GetNumberOfSharesForOwnerOfAForetak)
		v1.POST("/aksjebok/:orgnr/aksjeeier", api.GetOwnersForForetak)
	}

	v1.GET("/aksjebok/", api.GetForetak)
	v1.GET("/aksjebok/:orgnr", api.GetForetakByOrgnr)

	v1.GET("/health/", api.Health)
}

func serveApplication(router *gin.Engine) {
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		logrus.Warn("SERVER_PORT environment variable not set, using default port 8080 ")
		port = "8080"
	}

	router.Run(":" + port)
}
