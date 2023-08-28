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
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup Logrus
	logrus.SetLevel(logrus.DebugLevel)

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
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		logrus.Warn("SERVER_PORT environment variable not set, using default port 8080")
		port = "8080"
	}

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

	// Group API endpoints in /v1
	v1 := router.Group("/v1")

	v1.GET("/wallet", api.GetAllWallets)
	v1.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)
	v1.POST("/wallet", api.CreateWallet)

	v1.GET("/person/:pnr", api.GetWalletByPnr)

	v1.GET("/company/:orgnr", api.GetWalletByOrgnr)

	v1.GET("/foretak/:orgnr", api.GetForetakByOrgnr)
	v1.GET("/foretak/", api.GetForetak)
	v1.GET("/foretak", api.GetForetak)

	return router
}

func MaxBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}
