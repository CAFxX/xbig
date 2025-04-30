package xbig

import "math/big"

func E(prec uint) *big.Float {
	return ExpFloat(NewFloat(1).SetPrec(prec))
}

// func Pi(prec uint) *big.Float
// func Phi(prec uint) *big.Float
