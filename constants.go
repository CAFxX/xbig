package xbig

import "math/big"

// E is the constant e (Euler's number) with the given precision (in bits).
func E(prec uint) *big.Float {
	return ExpFloat(NewFloat(1).SetPrec(prec))
}

// func Pi(prec uint) *big.Float
// func Phi(prec uint) *big.Float
