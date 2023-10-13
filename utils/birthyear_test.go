package utils_test

import (
	"brok/navnetjener/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test50YearOldPerson(t *testing.T) {
	birthYear := utils.FindBirthYear("11117300000")

	assert.Equal(t, "1973", birthYear)
}

func Test10YearOldPerson(t *testing.T) {
	birthYear := utils.FindBirthYear("11111300000")

	assert.Equal(t, "2013", birthYear)
}

func Test120YearOldPerson(t *testing.T) {
	birthYear := utils.FindBirthYear("11110300000")

	assert.Equal(t, "2003", birthYear)
}
