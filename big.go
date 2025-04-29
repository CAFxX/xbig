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
