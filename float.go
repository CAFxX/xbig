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

func NewFloat[T floatNums](x T) *big.Float {
	return SetFloat(new(big.Float), x)
}

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

func AddFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Add(toFloat(x), toFloat(y))
}
func SubFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Sub(toFloat(x), toFloat(y))
}
func MulFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Mul(toFloat(x), toFloat(y))
}
func DivFloat[T, U floatNums](x T, y U) *big.Float {
	return new(big.Float).Quo(toFloat(x), toFloat(y))
}
func AbsFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Abs(toFloat(x))
}
func NegFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Neg(toFloat(x))
}
func CmpFloat[T, U floatNums](x T, y U) int {
	return toFloat(x).Cmp(toFloat(y))
}
func SqrtFloat[T floatNums](x T) *big.Float {
	return new(big.Float).Sqrt(toFloat(x))
}
func SetModeFloat[T floatNums](x T, mode big.RoundingMode) *big.Float {
	return NewFloat(x).SetMode(mode)
}
func SetPrecFloat[T floatNums](x T, prec uint) *big.Float {
	return NewFloat(x).SetPrec(prec)
}
func SetMantExpFloat[T floatNums](x T, exp int) *big.Float {
	return new(big.Float).SetMantExp(toFloat(x), exp)
}

// remove/reimplement these when/if math/big adds them
func PowFloat[T, U floatNums](x T, y U) *big.Float {
	return bigfloat.Pow(toFloat(x), toFloat(y))
}
func LogFloat[T floatNums](x T) *big.Float {
	return bigfloat.Log(toFloat(x))
}
func ExpFloat[T floatNums](x T) *big.Float {
	return bigfloat.Exp(toFloat(x))
}

func FMAFloat[T, U, V floatNums](x T, y U, z V) *big.Float {
	return NewFloat(FMARat(x, y, z)).SetPrec(max(toFloat(x).Prec(), toFloat(y).Prec(), toFloat(z).Prec()))
}
