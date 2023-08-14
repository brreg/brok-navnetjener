package main

import (
	"brok/navnetjener/api"
	"brok/navnetjener/database"
	"brok/navnetjener/model"
	"fmt"
	"log"
	"net/http"
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
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Print("Error loading .env.local file")
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

	/*
		Set a limit on the entire request body to 1 KiB
		Limit calculation (in bytes):
			257 (FirstName)
			257 (LastName)
			20 (Orgnr)
			13 (Pnr)
			6 (BirthDate)
			44 (WalletAddress)
			150 (JSON overhead)

		Total: 747, rounding up to 1024 to have some wiggle room
	*/
	router.Use(MaxBodySize(1024)) // 1 KiB limit

	router.GET("/wallet", api.GetAllWallets)
	router.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)
	router.POST("/wallet", api.CreateWallet)

	router.GET("/person/:pnr", api.GetWalletByPnr)

	router.GET("/company/:orgnr", api.GetWalletByOrgnr)

	return router
}

func MaxBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}
