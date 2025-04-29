package xbig

import (
	"math"
	"math/big"
	"testing"
)

func TestInt(t *testing.T) {
	if NewInt(1).Cmp(big.NewInt(1)) != 0 {
		t.Error("NewInt(1) != big.NewInt(1)")
	}
	if NewInt(int8(1)).Cmp(big.NewInt(1)) != 0 {
		t.Error("NewInt(int8(1)) != big.NewInt(1)")
	}
	if NewInt(uint8(1)).Cmp(big.NewInt(1)) != 0 {
		t.Error("NewInt(uint8(1)) != big.NewInt(1)")
	}
	if NewInt("1").Cmp(big.NewInt(1)) != 0 {
		t.Error("NewInt(\"1\") != big.NewInt(1)")
	}
}

func TestMulInt(t *testing.T) {
	if MulInt(-5, uint64(math.MaxUint64)).Cmp(new(big.Int).Mul(big.NewInt(-5), new(big.Int).SetUint64(math.MaxUint64))) != 0 {
		t.Error("MulInt(-5, math.MaxUint64) != new(big.Int).Mul(big.NewInt(-5), new(big.Int).SetUint64(math.MaxUint64))")
	}
}

func TestComplex(t *testing.T) {
	SubFloat(1, QuoRat(MulRat(-3, math.Pi), ExpInt(2, 128)))
}

func TestComplexStrings(t *testing.T) {
	MulRat("5.7", ExpInt("2", "123"))
}
