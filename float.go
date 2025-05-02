package xbig

import (
	"math/big"
	"math/bits"

	"github.com/ALTree/bigfloat"
)

type floatNums interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~*big.Int | ~*big.Float | ~*big.Rat |
		~string
}

// NewFloat creates a new [big.Float] from the given float value.
// The input is not modified.
// If the input is a string, it is parsed using [*math/big.Float).SetString(x.]
// The function returns the new [big.Float] on success. On failure, it returns nil.
// The precision of the new [big.Float] depends on the input value precision.
func NewFloat[T floatNums](x T) *big.Float {
	return SetFloat(new(big.Float), x)
}

// SetFloat sets the value of the given [big.Float] to the given float value.
// The input float value is not modified.
// If the input is a string, it is parsed using [*math/big.Float).SetString(x.]
// If the [big.Int] is nil, a new one is created. Otherwise, the existing one is modified.
// The function returns the modified [big.Float] on success.
// On failure, it returns nil and the state of the big.Float is undefined.
// The precision of the modified [big.Float] depends on the input value precision.
func SetFloat[T floatNums](f *big.Float, x T) *big.Float {
	if f == nil {
		f = new(big.Float)
	}
	switch x := any(x).(type) {
	case int:
		return f.SetInt64(int64(x)).SetPrec(bits.UintSize)
	case int64:
		return f.SetInt64(x)
	case int32:
		return f.SetInt64(int64(x)).SetPrec(32)
	case int16:
		return f.SetInt64(int64(x)).SetPrec(16)
	case int8:
		return f.SetInt64(int64(x)).SetPrec(8)
	case uint:
		return f.SetUint64(uint64(x)).SetPrec(bits.UintSize)
	case uint64:
		return f.SetUint64(x)
	case uint32:
		return f.SetUint64(uint64(x)).SetPrec(32)
	case uint16:
		return f.SetUint64(uint64(x)).SetPrec(16)
	case uint8:
		return f.SetUint64(uint64(x)).SetPrec(8)
	case float32:
		return f.SetFloat64(float64(x)).SetPrec(32)
	case float64:
		return f.SetFloat64(x)
	case *big.Float:
		return f.Set(x)
	case *big.Rat:
		return f.SetRat(x)
	case *big.Int:
		return f.SetInt(x)
	case string:
		r, _ := f.SetString(x)
		return r
	}
	panic("unreachable")
}

func toFloat[T floatNums](x T) *big.Float {
	if f, ok := any(x).(*big.Float); ok {
		return f
	}
	return NewFloat(x)
}

// AddFloat adds two float values and returns the result as a new [big.Float].
// The input values are not modified.
//
// See [big.Float.Add] for more details.
func AddFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Add(toFloat(x), toFloat(y))
}

// SubFloat subtracts two float values and returns the result as a new [big.Float].
// The input values are not modified.
//
// See [big.Float.Sub] for more details.
func SubFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Sub(toFloat(x), toFloat(y))
}

// MulFloat multiplies two float values and returns the result as a new [big.Float].
// The input values are not modified.
//
// See [big.Float.Mul] for more details.
func MulFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Mul(toFloat(x), toFloat(y))
}

// QuoFloat divides two float values and returns the result as a new [big.Float].
// The input values are not modified.
//
// See [big.Float.Quo] for more details.
func QuoFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Quo(toFloat(x), toFloat(y))
}

// AbsFloat returns the absolute value of the given float value as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.Abs] for more details.
func AbsFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Abs(toFloat(x))
}

// NegFloat returns the negation of the given float value as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.Neg] for more details.
func NegFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Neg(toFloat(x))
}

// CmpFloat compares two float values and returns:
//
//	-1 if x < y
//	0 if x == y
//	1 if x > y
//
// The input values are not modified.
//
// See [big.Float.Cmp] for more details.
func CmpFloat[T, U floatNums](x T, y U) int {
	return toFloat(x).Cmp(toFloat(y))
}

// SqrtFloat returns the square root of the given float value as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.Sqrt] for more details.
func SqrtFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Sqrt(toFloat(x))
}

// SetModeFloat sets the rounding mode of the given float value and returns the result as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.SetMode] for more details.
func SetModeFloat[T floatNums](x T, mode big.RoundingMode) *big.Float {
	return NewFloat(x).SetMode(mode)
}

// SetPrecFloat sets the precision of the given float value and returns the result as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.SetPrec] for more details.
func SetPrecFloat[T floatNums](x T, prec uint) *big.Float {
	return NewFloat(x).SetPrec(prec)
}

// SetMantExpFloat sets the mantissa and exponent of the given float value and returns the result as a new [big.Float].
// The input value is not modified.
//
// See [big.Float.SetMantExp] for more details.
func SetMantExpFloat[T floatNums](x T, exp int) *big.Float {
	return new(big.Float).SetMantExp(toFloat(x), exp)
}

// remove/reimplement these when/if math/big adds them

// PowFloat computes x^y and returns the result as a new [big.Float].
// The input values are not modified.
//
// See [github.com/ALTree/bigfloat.Pow] for more details.
func PowFloat[T, U floatNums](x T, y U) *big.Float {
	return bigfloat.Pow(toFloat(x), toFloat(y))
}

// LogFloat computes the natural logarithm of x and returns the result as a new [big.Float].
// The input value is not modified.
//
// See [github.com/ALTree/bigfloat.Log] for more details.
func LogFloat[T floatNums](x T) *big.Float {
	return bigfloat.Log(toFloat(x))
}

// ExpFloat computes e^x and returns the result as a new [big.Float].
// The input value is not modified.
//
// See [github.com/ALTree/bigfloat.Exp] for more details.
func ExpFloat[T floatNums](x T) *big.Float {
	return bigfloat.Exp(toFloat(x))
}

// FMAFloat computes the fused multiply-add of x, y, and z and returns the result as a new [big.Float].
// The input values are not modified.
// The precision of the result is the maximum of the precision of x, y, and z.
func FMAFloat[T, U, V floatNums](x T, y U, z V) *big.Float {
	return NewFloat(FMARat(x, y, z)).SetPrec(max(toFloat(x).Prec(), toFloat(y).Prec(), toFloat(z).Prec()))
}

// LogBaseFloat computes the logarithm of a in base b and returns the result as a new [big.Float].
// The input values are not modified.
func LogBaseFloat[T, U floatNums](a T, b U) *big.Float {
	return QuoFloat(LogFloat(a), LogFloat(b))
}
