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

## Notes

This package is a WIP: [APIs](https://pkg.go.dev/github.com/CAFxX/xbig) are not stable yet.

Ergonomy comes at a slight increase in memory allocations and copies, as well as dispatch logic. Improvements in the go compiler/runtime should eventually be able to remove most of these overheads: 
[golang/go#59591](https://github.com/golang/go/issues/59591), 
[golang/go#64824](https://github.com/golang/go/issues/64824),
[golang/go#48849](https://github.com/golang/go/issues/48849)
and so on.
