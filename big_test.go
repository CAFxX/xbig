package xbig

import (
	"math"
	"math/big"
	"strconv"
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

func TestFMAInt(t *testing.T) {
	if CmpInt(FMAInt(1, 2, 3), 5) != 0 {
		t.Error("FMAInt(1, 2, 3) != big.NewInt(5)")
	}
	if CmpInt(FMAInt(7, "3", -1), 20) != 0 {
		t.Error("FMAInt(7, \"3\", -1) != big.NewInt(20)")
	}
}

func TestConstants(t *testing.T) {
	if e, _ := E(64).Float64(); e != math.E {
		t.Errorf("E: got %g, want %g", e, math.E)
	}
	if pi, _ := Pi(64).Float64(); pi != math.Pi {
		t.Errorf("Pi: got %g, want %g", pi, math.Pi)
	}
	if phi, _ := Phi(64).Float64(); phi != math.Phi {
		t.Errorf("Phi: got %g, want %g", phi, math.Phi)
	}
}

func BenchmarkPi(b *testing.B) {
	for i := 6; i <= 20; i += 2 {
		b.Run(strconv.Itoa(1<<i), func(b *testing.B) {
			for b.Loop() {
				Pi(1 << i)
			}
		})
	}
}

func BenchmarkE(b *testing.B) {
	for i := 6; i <= 20; i += 2 {
		b.Run(strconv.Itoa(1<<i), func(b *testing.B) {
			for b.Loop() {
				E(1 << i)
			}
		})
	}
}

func BenchmarkPhi(b *testing.B) {
	for i := 6; i <= 20; i += 2 {
		b.Run(strconv.Itoa(1<<i), func(b *testing.B) {
			for b.Loop() {
				Phi(1 << i)
			}
		})
	}
}
