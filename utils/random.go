package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	r "math/rand"
	"time"
)

func RandomNumber(min int, max int) string {
	random := r.New(r.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%02d", random.Intn(max-min)+min)
}

func RandomOrgnr() string {
	return RandomNumber(111111111, 999999999)
}

func RandomFnr() string {
	return RandomNumber(11111111111, 99999999999)
}

func RandomWalletAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}
