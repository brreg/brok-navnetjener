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
		err := godotenv.Load(".env")
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
	router.Run(":8080")
	fmt.Println("Server running at port 8080")
}

func routerConfig() *gin.Engine {
	router := gin.Default()
	router.GET("/wallet", api.GetAllWallets)
	router.GET("/wallet/:walletAddress", api.GetWalletByID)
	router.POST("/wallet", api.CreateWallet)

	return router
}
