package main

import (
	"brok/navnetjener/api"
	"brok/navnetjener/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()

	loggerConfig()

	loadDatabase()

	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Print("Error loading .env.local file")
	}
}

func loggerConfig() {
	logLevel := os.Getenv("LOG_LEVEL")

	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func loadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&model.Wallet{})
}

func serveApplication() {
	router := routerConfig()
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		logrus.Warn("SERVER_PORT environment variable not set, using default port 8080")
		port = "8080"
	}

	router.Run(":" + port)
}

func routerConfig() *gin.Engine {

	env, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		logrus.Warn("ENVIRONMENT environment variable not set, using default value: dev")
		env = "development"
	}

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if env == "development" {
		gin.SetMode(gin.DebugMode)
	}

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
	router.Use(maxBodySize(1024)) // 1 KiB limit

	// Group API endpoints in /v1
	v1 := router.Group("/v1")

	v1.GET("/wallet", api.GetAllWallets)
	v1.GET("/wallet/:walletAddress", api.GetWalletByWalletAddress)
	v1.POST("/wallet", api.CreateWallet)

	v1.GET("/person/:pnr", api.GetAllForetakForPerson)

	v1.GET("/foretak/:orgnr", api.GetForetakByOrgnr)
	v1.GET("/foretak/", api.GetForetak)
	v1.GET("/foretak", api.GetForetak)

	return router
}

func maxBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}
