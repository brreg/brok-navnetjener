package database

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Database *gorm.DB

func Connect() {
	var err error

	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		logrus.Error("DB_HOST environment variable not set")
	}

	username, exists := os.LookupEnv("DB_USER")
	if !exists {
		logrus.Error("DB_USER environment variable not set")
	}

	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		logrus.Error("DB_PASSWORD environment variable not set")
	}

	databaseName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		logrus.Error("DB_NAME environment variable not set")
	}

	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		logrus.Error("DB_PORT environment variable not set")
	}

	sslmode := "disable"
	if os.Getenv("DB_SSL_ENCRYPTION") == "true" {
		sslmode = "require"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Oslo", host, username, password, databaseName, port, sslmode)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "navnetjener.",
			SingularTable: false,
		}}

	timeToWaitBetweenRetries := 10 * time.Second
	for {
		Database, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err == nil {
			break
		}
		logrus.Errorf("Attempt to connect to database failed. Retrying in %s...\n", timeToWaitBetweenRetries)
		time.Sleep(timeToWaitBetweenRetries)
	}

	logrus.Info("Successfully connected to the database")
}
