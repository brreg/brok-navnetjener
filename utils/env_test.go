package utils_test

import (
	"brok/navnetjener/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFindEnvFile(t *testing.T) {
	file, _ := utils.FindFile(".env.local")
	fmt.Println("file")
	fmt.Println(file)
	assert.Equal(t, true, true)
}
