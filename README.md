# xbig

[`github.com/CAFxX/xbig`](https://github.com/CAFxX/xbig) is a generics-based convenience wrapper package that aims at making the [`math/big`](https://pkg.go.dev/math/big) API more ergonomic to use.

The following are all valid ways to use `xbig`:

```go
QuoFloat(MulFloat(math.PI, 2), ExpInt(2, 128)) // 2π/2¹²⁸
MulInt("4793825749327547329903472", "9843728974589574362758946543567483265783432")
```

This package is a WIP: APIs are not stable yet. [![Go Reference](https://pkg.go.dev/badge/github.com/CAFxX/xbig.svg)](https://pkg.go.dev/github.com/CAFxX/xbig)
