package model_test

import (
	"brok/navnetjener/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeString(t *testing.T) {
	word1 := model.SanitizeString("123 asdf")
	assert.Equal(t, word1, "123 asdf")

	word2 := model.SanitizeString("hello !#¤%&/()=?`'*¨^><|,.-;:_ world")
	assert.Equal(t, word2, "hello  world")

	word3 := model.SanitizeString("0x8be848ce9ebba1e304e6daa1d6b1b40f17e478fd")
	assert.Equal(t, word3, "0x8be848ce9ebba1e304e6daa1d6b1b40f17e478fd")

	word4 := strings.ToLower("0x8be848ce9ebba1e304e6daa1d6b1b40F17E478FD")
	assert.Equal(t, word4, "0x8be848ce9ebba1e304e6daa1d6b1b40f17e478fd")
}
