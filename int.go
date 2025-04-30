package xbig

import (
	"math"
	"math/big"
	"math/rand"
)

type intNums interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~*big.Int |
		~string |
		~[]big.Word | ~[]byte
}

func NewInt[T intNums](x T) *big.Int {
	return SetInt(new(big.Int), x)
}

func SetInt[T intNums](f *big.Int, x T) *big.Int {
	if f == nil {
		f = new(big.Int)
	}
	switch x := any(x).(type) {
	case int:
		return f.SetInt64(int64(x))
	case int64:
		return f.SetInt64(x)
	case int32:
		return f.SetInt64(int64(x))
	case int16:
		return f.SetInt64(int64(x))
	case int8:
		return f.SetInt64(int64(x))
	case uint:
		return f.SetUint64(uint64(x))
	case uint64:
		return f.SetUint64(x)
	case uint32:
		return f.SetUint64(uint64(x))
	case uint16:
		return f.SetUint64(uint64(x))
	case uint8:
		return f.SetUint64(uint64(x))
	case *big.Int:
		return f.Set(x)
	case string:
		r, _ := f.SetString(x, 0)
		return r
	case []big.Word:
		return f.SetBits(x)
	case []byte:
		return f.SetBytes(x)
	}
	panic("unreachable")
}

func toInt[T intNums](x T) *big.Int {
	if x, ok := any(x).(*big.Int); ok {
		return x
	}
	return NewInt(x)
}

func AddInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Add(toInt(x), toInt(y))
}
func SubInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Sub(toInt(x), toInt(y))
}
func MulInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Mul(toInt(x), toInt(y))
}
func DivInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Div(toInt(x), toInt(y))
}
func ModInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Mod(toInt(x), toInt(y))
}
func DivModInt[T, U intNums](x T, y U) (*big.Int, *big.Int) {
	return new(big.Int).DivMod(toInt(x), toInt(y), nil)
}
func ModInverseInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).ModInverse(toInt(x), toInt(y))
}
func ModInverseSqrtInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).ModSqrt(toInt(x), toInt(y))
}
func QuoInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Quo(toInt(x), toInt(y))
}
func RemInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Rem(toInt(x), toInt(y))
}
func QuoRemInt[T, U intNums](x T, y U) (*big.Int, *big.Int) {
	return new(big.Int).QuoRem(toInt(x), toInt(y), new(big.Int))
}
func ExpInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Exp(toInt(x), toInt(y), nil)
}
func ExpModInt[T, U, V intNums](x T, y U, z V) *big.Int {
	return new(big.Int).Exp(toInt(x), toInt(y), toInt(z))
}
func GCDInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).GCD(nil, nil, toInt(x), toInt(y))
}
func GCPPolyInt[T, U intNums](a T, b U) (*big.Int, *big.Int, *big.Int) {
	x, y := new(big.Int), new(big.Int)
	return x, y, new(big.Int).GCD(x, y, toInt(a), toInt(b))
}
func AbsInt[T intNums](x T) *big.Int {
	return new(big.Int).Abs(toInt(x))
}
func NegInt[T intNums](x T) *big.Int {
	return new(big.Int).Neg(toInt(x))
}
func CmpInt[T, U intNums](x T, y U) int {
	return toInt(x).Cmp(toInt(y))
}
func CmpAbsInt[T, U intNums](x T, y U) int {
	return toInt(x).CmpAbs(toInt(y))
}
func AndInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).And(toInt(x), toInt(y))
}
func AndNotInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).AndNot(toInt(x), toInt(y))
}
func OrInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Or(toInt(x), toInt(y))
}
func XorInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Xor(toInt(x), toInt(y))
}
func NotInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Not(toInt(x))
}
func LshInt[T, U intNums](x T, y U) *big.Int {
	yi := toInt(y)
	if !yi.IsUint64() || yi.Uint64() > math.MaxUint {
		panic("shift too large")
	}
	return new(big.Int).Lsh(toInt(x), uint(yi.Uint64()))
}
func RshInt[T, U intNums](x T, y U) *big.Int {
	yi := toInt(y)
	if !yi.IsUint64() || yi.Uint64() > math.MaxUint {
		panic("shift too large")
	}
	return new(big.Int).Rsh(toInt(x), uint(yi.Uint64()))
}
func RandInt[T intNums](r *rand.Rand, x T) *big.Int {
	return new(big.Int).Rand(r, toInt(x))
}

func FMAInt[T, U, V intNums](x T, y U, z V) *big.Int {
	rx := NewInt(x)
	rx.Mul(rx, toInt(y))
	return rx.Add(rx, toInt(z))
}
