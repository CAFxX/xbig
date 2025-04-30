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

// NewInt creates a new [big.Int] from the given integer value.
// The input is not modified.
// The function the new [big.Int] on success. On failure, it returns nil.
func NewInt[T intNums](x T) *big.Int {
	return SetInt(new(big.Int), x)
}

// SetInt sets the value of the given [big.Int] to the given integer value.
// The input integer value is not modified.
// If the input is a string, it is parsed using math/big.(*Int).SetString(x, 0).
// If the input is a []byte, it is parsed using math/big.(*Int).SetBytes(x).
//
// If the [big.Int] is nil, a new one is created. Otherwise, the existing one is modified.
// The function returns the modified [big.Int] on success. On failure, it returns nil.
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

// AddInt adds two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Add for more details.
func AddInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Add(toInt(x), toInt(y))
}

// SubInt subtracts two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Sub for more details.
func SubInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Sub(toInt(x), toInt(y))
}

// MulInt multiplies two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See [math/big.(*Int).Mul] for more details.
func MulInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Mul(toInt(x), toInt(y))
}

// DivInt divides two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Div for more details.
func DivInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Div(toInt(x), toInt(y))
}

// ModInt computes the modulus of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Mod for more details.
func ModInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Mod(toInt(x), toInt(y))
}

// DivModInt divides two integers and returns the quotient and remainder as two new [big.Int]s.
// The inputs are unmodified.
//
// See math/big.(*Int).DivMod for more details.
func DivModInt[T, U intNums](x T, y U) (*big.Int, *big.Int) {
	return new(big.Int).DivMod(toInt(x), toInt(y), nil)
}

// ModInverseInt computes the modular inverse of x modulo y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).ModInverse for more details.
func ModInverseInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).ModInverse(toInt(x), toInt(y))
}

// ModSqrtInt computes the modular square root of x modulo y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).ModSqrt for more details.
func ModSqrtInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).ModSqrt(toInt(x), toInt(y))
}

// QuoInt divides two integers and returns the quotient as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Quo for more details.
func QuoInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Quo(toInt(x), toInt(y))
}

// RemInt computes the modulus of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Rem for more details.
func RemInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Rem(toInt(x), toInt(y))
}

// QuoRemInt divides two integers and returns the quotient and remainder as two new [big.Int]s.
// The inputs are unmodified.
//
// See math/big.(*Int).QuoRem for more details.
func QuoRemInt[T, U intNums](x T, y U) (*big.Int, *big.Int) {
	return new(big.Int).QuoRem(toInt(x), toInt(y), new(big.Int))
}

// ExpInt computes x to the power of y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Exp for more details.
func ExpInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Exp(toInt(x), toInt(y), nil)
}

// ExpModInt computes x to the power of y modulo z and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Exp for more details.
func ExpModInt[T, U, V intNums](x T, y U, z V) *big.Int {
	return new(big.Int).Exp(toInt(x), toInt(y), toInt(z))
}

// GCDInt computes the greatest common divisor of x and y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).GCD for more details.
func GCDInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).GCD(nil, nil, toInt(x), toInt(y))
}

// GCDPolyInt computes the greatest common divisor of x and y and returns the result as a new [big.Int],
// and two new [big.Int]s that are the coefficients of the GCD polynomial.
// The inputs are unmodified.
//
// See math/big.(*Int).GCD for more details.
func GCDPolyInt[T, U intNums](a T, b U) (*big.Int, *big.Int, *big.Int) {
	x, y := new(big.Int), new(big.Int)
	return x, y, new(big.Int).GCD(x, y, toInt(a), toInt(b))
}

// AbsInt computes the absolute value of x and returns the result as a new [big.Int].
// The input is unmodified.
//
// See math/big.(*Int).Abs for more details.
func AbsInt[T intNums](x T) *big.Int {
	return new(big.Int).Abs(toInt(x))
}

// NegInt computes the negation of x and returns the result as a new [big.Int].
// The input is unmodified.
//
// See math/big.(*Int).Neg for more details.
func NegInt[T intNums](x T) *big.Int {
	return new(big.Int).Neg(toInt(x))
}

// CmpInt compares two integers and returns:
//
//	-1 if x < y
//	0 if x == y
//	1 if x > y
//
// The inputs are unmodified.
//
// See math/big.(*Int).Cmp for more details.
func CmpInt[T, U intNums](x T, y U) int {
	return toInt(x).Cmp(toInt(y))
}

// CmpAbsInt compares the absolute values of two integers and returns:
//
//	-1 if |x| < |y|
//	0 if |x| == |y|
//	1 if |x| > |y|
//
// The inputs are unmodified.
//
// See math/big.(*Int).CmpAbs for more details.
func CmpAbsInt[T, U intNums](x T, y U) int {
	return toInt(x).CmpAbs(toInt(y))
}

// AndInt computes the bitwise AND of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).And for more details.
func AndInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).And(toInt(x), toInt(y))
}

// AndNotInt computes the bitwise AND NOT of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).AndNot for more details.
func AndNotInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).AndNot(toInt(x), toInt(y))
}

// OrInt computes the bitwise OR of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Or for more details.
func OrInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Or(toInt(x), toInt(y))
}

// XorInt computes the bitwise XOR of two integers and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Xor for more details.
func XorInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Xor(toInt(x), toInt(y))
}

// NotInt computes the bitwise NOT of x and returns the result as a new [big.Int].
// The input is unmodified.
//
// See math/big.(*Int).Not for more details.
func NotInt[T, U intNums](x T, y U) *big.Int {
	return new(big.Int).Not(toInt(x))
}

// LshInt computes the left shift of x by y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Lsh for more details.
func LshInt[T, U intNums](x T, y U) *big.Int {
	yi := toInt(y)
	if !yi.IsUint64() || yi.Uint64() > math.MaxUint {
		panic("shift too large")
	}
	return new(big.Int).Lsh(toInt(x), uint(yi.Uint64()))
}

// RshInt computes the right shift of x by y and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).Rsh for more details.
func RshInt[T, U intNums](x T, y U) *big.Int {
	yi := toInt(y)
	if !yi.IsUint64() || yi.Uint64() > math.MaxUint {
		panic("shift too large")
	}
	return new(big.Int).Rsh(toInt(x), uint(yi.Uint64()))
}

// RandInt generates a random integer in the range [0, x) and returns it as a new [big.Int].
// The input x is unmodified.
//
// See math/big.(*Int).Rand for more details.
func RandInt[T intNums](r *rand.Rand, x T) *big.Int {
	return new(big.Int).Rand(r, toInt(x))
}

// FMAInt computes the fused multiply-add of x, y, and z and returns the result as a new [big.Int].
// The inputs are unmodified.
//
// See math/big.(*Int).FMA for more details.
func FMAInt[T, U, V intNums](x T, y U, z V) *big.Int {
	rx := NewInt(x)
	rx.Mul(rx, toInt(y))
	return rx.Add(rx, toInt(z))
}
