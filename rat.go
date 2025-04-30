package xbig

import (
	"math/big"
)

type ratNums interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~*big.Int | ~*big.Float | ~*big.Rat |
		~string
}

func NewRat[T ratNums](x T) *big.Rat {
	return SetRat(new(big.Rat), x)
}

func SetRat[T ratNums](f *big.Rat, x T) *big.Rat {
	if f == nil {
		f = new(big.Rat)
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
	case float32:
		return f.SetFloat64(float64(x))
	case float64:
		return f.SetFloat64(x)
	case *big.Float:
		r, _ := x.Rat(f)
		return r
	case *big.Rat:
		return f.Set(x)
	case *big.Int:
		return f.SetInt(x)
	case string:
		r, _ := f.SetString(x)
		return r
	}
	panic("unreachable")
}

func toRat[T ratNums](x T) *big.Rat {
	if x, ok := any(x).(*big.Rat); ok {
		return x
	}
	return NewRat(x)
}

func NewRatFrac[T, U ratNums](x T, y U) *big.Rat {
	return SetRatFrac(new(big.Rat), x, y)
}

func SetRatFrac[T, U ratNums](f *big.Rat, x T, y U) *big.Rat {
	if isIntLike(x) && isIntLike(y) {
		return f.SetFrac(isIntLikeToInt(x), isIntLikeToInt(y))
	}
	f = SetRat(f, x)
	return f.Quo(f, NewRat(y))
}

func isIntLike[T ratNums](x T) bool {
	switch any(x).(type) {
	case *big.Int, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return true
	}
	return false
}

func isIntLikeToInt[T ratNums](x T) *big.Int {
	switch x := any(x).(type) {
	case *big.Int:
		return x
	case int:
		return toInt(x)
	case int64:
		return toInt(x)
	case int32:
		return toInt(x)
	case int16:
		return toInt(x)
	case int8:
		return toInt(x)
	case uint:
		return toInt(x)
	case uint64:
		return toInt(x)
	case uint32:
		return toInt(x)
	case uint16:
		return toInt(x)
	case uint8:
		return toInt(x)
	}
	return nil
}

func AddRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Add(toRat(x), toRat(y))
}
func SubRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Sub(toRat(x), toRat(y))
}
func MulRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Mul(toRat(x), toRat(y))
}
func QuoRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Quo(toRat(x), toRat(y))
}
func AbsRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Abs(toRat(x))
}
func NegRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Neg(toRat(x))
}
func InvRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Inv(toRat(x))
}
func CmpRat[T, U ratNums](x T, y U) int {
	return toRat(x).Cmp(toRat(y))
}
