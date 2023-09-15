package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func loadEnv() {
	file, err := FindFile(".env")

	if err != nil {
		logrus.Error(err)
	}

	err = godotenv.Load(file)
	if err != nil {
		logrus.Error("Error loading .env.local file")
	}
}

// Looks for the file in current and parent directory.
//
// Returns the relative filepath to the file
func FindFile(filename string) (string, error) {
	currentDir := "."
	// Check if the file exists in the current directory
	filePath := filepath.Join(currentDir, filename)
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	} else if !os.IsNotExist(err) {
		return "", err
	}

	filePath = filepath.Join("./../", filePath)
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	} else if !os.IsNotExist(err) {
		return "", err
	}

	return "", errors.New("Could not find " + filename)
}
