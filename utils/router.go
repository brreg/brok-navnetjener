package utils

import (
	"brok/navnetjener/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Loads Env variables, configures logging and database connection, and creates a gin router.
//
// After running this function, API endpoints needs to be defined before usage of the router.
func Setup() *gin.Engine {
	loadEnv()
	loggerConfig()
	database.Connect()
	return routerConfig()
}

func routerConfig() *gin.Engine {

	router := gin.New()

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

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/v1/health/"),
		gin.Recovery(),
	)

	return router
}

func maxBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}
