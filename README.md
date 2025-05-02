# xbig

[`github.com/CAFxX/xbig`](https://github.com/CAFxX/xbig) is a generics-based convenience wrapper package that aims at making the [`math/big`](https://pkg.go.dev/math/big) API more ergonomic to use.

[![Go Reference](https://pkg.go.dev/badge/github.com/CAFxX/xbig.svg)](https://pkg.go.dev/github.com/CAFxX/xbig)

## Examples

The following are all valid ways to use `xbig`:

```go
// 2π/2¹²⁸
QuoFloat(MulFloat(math.PI, 2), ExpInt(2, 128))

// 4793825749327547329903472 * 9843728974589574362758946543567483265783432
MulInt("4793825749327547329903472", "9843728974589574362758946543567483265783432")
```

You can test `xbig` on the [Go playground](https://go.dev/play/p/OFIjjTg47yt).

## Functionalities

### Simplified instance construction

`NewInt`, `NewRat`, `NewFloat` accept the following source and create a new instance of the corresponding `big.Int`, `big.Rat`, or `big.Float` type.
`SetInt`, `SetRat`, `SetFloat` do the same, but set the provided `big.Int`, `big.Rat`, or `big.Float` instance.

The following table summarizes which conversions are supported (✅ means supported, and the conversion can not fail; ⚠️ means supported, but the conversion can fail depending on the source value; ❌ means not supported)

| Source type  | `*big.Int` | `*big.Rat`   | `*big.Float` |
| ------------ | ---------- | ------------ | ------------ |
| `int`        | ✅         | ✅           | ✅           |
| `int8`       | ✅         | ✅           | ✅           |
| `int16`      | ✅         | ✅           | ✅           |
| `int32`      | ✅         | ✅           | ✅           |
| `int64`      | ✅         | ✅           | ✅           |
| `uint`       | ✅         | ✅           | ✅           |
| `uint8`      | ✅         | ✅           | ✅           |
| `uint16`     | ✅         | ✅           | ✅           |
| `uint32`     | ✅         | ✅           | ✅           |
| `uint64`     | ✅         | ✅           | ✅           |
| `float32`    | ❌         | ⚠️ (±∞, NaN) | ⚠️ (NaN)     |
| `float64`    | ❌         | ⚠️ (±∞, NaN) | ⚠️ (NaN)     |
| `*big.Int`   | ✅         | ✅           | ✅           |
| `*big.Rat`   | ❌         | ✅           | ✅           |
| `*big.Float` | ❌         | ⚠️ (±∞)      | ✅           |
| `string`     | ⚠️         | ⚠️           | ⚠️           |
| `[]byte`     | ✅         | ❌           | ❌           |
| `[]big.Word` | ✅         | ❌           | ❌           |

### Simplified function composition

Most aritchmetic functions in the package automatically perform type conversions (see the previous section), so it is e.g. possible to pass a `big.Rat` or `int` to functions that in `math/big` only accept a `big.Float` as input.

Furthermore, all functions (apart from the `Set*` functions) allocate a new result and never modify their inputs, so accidental shadowing of variables is not possible.

### Arbitrary precision constants

The functions `Pi`, `Phi`, and `E` return the mathematical constants $\pi$, $\varphi$, and $e$ with arbitrary precision.

## Notes

This package is a WIP: [APIs](https://pkg.go.dev/github.com/CAFxX/xbig) are not stable yet.

Ergonomy comes at a slight increase in memory allocations and copies, as well as dispatch logic. Improvements in the go compiler/runtime should eventually be able to remove most of these overheads:
[golang/go#59591](https://github.com/golang/go/issues/59591),
[golang/go#64824](https://github.com/golang/go/issues/64824),
[golang/go#48849](https://github.com/golang/go/issues/48849)
and so on.
