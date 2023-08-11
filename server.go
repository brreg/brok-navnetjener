package main

import (
	"brok/navnetjener/api"
	"brok/navnetjener/database"
	"brok/navnetjener/model"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Wallet{})
}

func serveApplication() {
	router := routerConfig()
	port := os.Getenv("SERVER_PORT")
	router.Run(":" + port)
	fmt.Printf("Server running at port %s", port)
}

func routerConfig() *gin.Engine {
	router := gin.Default()
	router.GET("/wallet", api.GetAllWallets)
	router.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)
	router.POST("/wallet", api.CreateWallet)

	router.GET("/person/:pnr", api.GetWalletByPnr)

	router.GET("/company/:orgnr", api.GetWalletByOrgnr)

	return router
}
