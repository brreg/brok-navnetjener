package utils

import (
	"math/big"
)

// ToDecimal converts wei to decimal
// returns value as a string
// example: 1000000000000000000 wei = 1 ETH
func ToDecimal(wei string) string {
	decimals := big.NewFloat(1000000000000000000)
	weiFloat := new(big.Float)
	weiFloat.SetString(wei)
	weiFloat.Quo(weiFloat, decimals)
	weiFloatString := weiFloat.Text('f', 0)
	return weiFloatString
}
