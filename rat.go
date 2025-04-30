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

// NewRat creates a new [big.Rat] from the given rational value.
// The input is not modified.
// If the input is a string, it is parsed using [math/big.(*Rat).SetString(x)].
// The function returns the new [big.Rat] on success. On failure, it returns nil.
func NewRat[T ratNums](x T) *big.Rat {
	return SetRat(new(big.Rat), x)
}

// SetRat sets the value of the given [big.Rat] to the given rational value.
// The input rational value is not modified.
// If the input is a string, it is parsed using [math/big.(*Rat).SetString(x)].
// If the input is a float, and it can not be converted to a rational [math/big.Rat], it returns nil.
// The function returns the modified [big.Rat] on success.
// On failure, it returns nil and the state of the [big.Rat] is undefined.
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

// NewRatFrac creates a new [big.Rat] from the given rational values.
// The inputs are not modified.
func NewRatFrac[T, U ratNums](x T, y U) *big.Rat {
	return SetRatFrac(new(big.Rat), x, y)
}

// SetRatFrac sets the value of the given [big.Rat] to the given rational values.
// The inputs are not modified.
func SetRatFrac[T, U ratNums](f *big.Rat, x T, y U) *big.Rat {
	if isIntLike(x) && isIntLike(y) {
		return f.SetFrac(isIntLikeToInt(x), isIntLikeToInt(y))
	}
	f = SetRat(f, x)
	return f.Quo(f, toRat(y))
}

func isIntLike[T ratNums](x T) bool {
	switch x := any(x).(type) {
	case *big.Int, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return true
	case *big.Rat:
		return x.IsInt()
	case *big.Float:
		return x.IsInt()
	case float64:
		return x == float64(int64(x))
	case float32:
		return x == float32(int32(x))
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
	case *big.Rat:
		return x.Num()
	case *big.Float:
		r, a := x.Int(nil)
		if a != big.Exact {
			panic("unreachable")
		}
		return r
	case float64:
		return toInt(int64(x))
	case float32:
		return toInt(int32(x))
	}
	return nil
}

// AddRat adds two rational values and returns the result as a new [big.Rat].
// The input values are not modified.
//
// See [math/big.(*Rat).Add] for more details.
func AddRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Add(toRat(x), toRat(y))
}

// SubRat subtracts two rational values and returns the result as a new [big.Rat].
// The input values are not modified.
//
// See [math/big.(*Rat).Sub] for more details.
func SubRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Sub(toRat(x), toRat(y))
}

// MulRat multiplies two rational values and returns the result as a new [big.Rat].
// The input values are not modified.
//
// See [math/big.(*Rat).Mul] for more details.
func MulRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Mul(toRat(x), toRat(y))
}

// QuoRat divides two rational values and returns the result as a new [big.Rat].
// The input values are not modified.
//
// See [math/big.(*Rat).Quo] for more details.
func QuoRat[T, U ratNums](x T, y U) *big.Rat {
	return new(big.Rat).Quo(toRat(x), toRat(y))
}

// AbsRat returns the absolute value of the given rational value as a new [big.Rat].
// The input value is not modified.
//
// See [math/big.(*Rat).Abs] for more details.
func AbsRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Abs(toRat(x))
}

// NegRat returns the negation of the given rational value as a new [big.Rat].
// The input value is not modified.
//
// See [math/big.(*Rat).Neg] for more details.
func NegRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Neg(toRat(x))
}

// InvRat returns the inverse of the given rational value as a new [big.Rat].
// The input value is not modified.
//
// See [math/big.(*Rat).Inv] for more details.
func InvRat[T ratNums](x T) *big.Rat {
	return new(big.Rat).Inv(toRat(x))
}

// CmpRat compares two rational values and returns:
//
//	-1 if x < y
//	0 if x == y
//	1 if x > y
//
// The input values are not modified.
//
// See [math/big.(*Rat).Cmp] for more details.
func CmpRat[T, U ratNums](x T, y U) int {
	return toRat(x).Cmp(toRat(y))
}

// FMARat computes the fused multiply-add of three rational values and returns the result as a new [big.Rat].
// The inputs are not modified.
func FMARat[T, U, V ratNums](x T, y U, z V) *big.Rat {
	rx := NewRat(x)
	rx.Mul(rx, toRat(y))
	return rx.Add(rx, toRat(z))
}
