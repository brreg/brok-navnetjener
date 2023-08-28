package database

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	logrus.Debug(dsn)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
