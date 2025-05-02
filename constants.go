package xbig

import (
	"math"
	"math/big"
)

// E is the constant e (Euler's number) with the given precision (in bits).
func E(prec uint) *big.Float {
	return ExpFloat(NewFloat(1).SetPrec(prec))
}

// Pi is the constant π with the given precision (in bits).
func Pi(prec uint) *big.Float {
	// Pi calculates Pi to the specified precision in bits using the Gauss-Legendre algorithm.
	// prec must be > 0. If prec is >= big.MaxPrec, the function panics.
	// The algorithm has quadratic convergence, roughly doubling the correct bits each iteration.
	if prec == 0 {
		panic("picalc.Pi: precision must be > 0")
	}
	// Ensure precision is within the limits supported by big.Float calculations.
	// We need internalPrec > prec, so check prec against MaxPrec.
	if prec >= big.MaxPrec {
		panic("picalc.Pi: precision is too large (>= big.MaxPrec)")
	}

	// Use higher internal precision for calculations to mitigate rounding errors
	// during intermediate steps. The number of guard bits needed grows roughly
	// logarithmically with the target precision. Add a safety buffer.
	internalPrec := prec + uint(math.Log2(float64(prec))) + 32
	// Clamp internalPrec to MaxPrec if the calculation overflows or exceeds the limit.
	if internalPrec < prec || internalPrec > big.MaxPrec { // Check for overflow and limit
		internalPrec = big.MaxPrec
	}

	// --- Constants ---
	// It's crucial these constants use internalPrec.
	one := big.NewFloat(1).SetPrec(internalPrec)
	two := big.NewFloat(2).SetPrec(internalPrec)
	four := big.NewFloat(4).SetPrec(internalPrec)
	half := big.NewFloat(0.5).SetPrec(internalPrec)

	// --- Initial values (Gauss-Legendre a_0, b_0, t_0, p_0) ---
	a := new(big.Float).SetPrec(internalPrec).Set(one) // a_0 = 1
	// b_0 = 1 / sqrt(2) which is equivalent to sqrt(1/2)
	b := new(big.Float).SetPrec(internalPrec).Sqrt(half)
	t := new(big.Float).SetPrec(internalPrec).Quo(one, four) // t_0 = 1/4
	p := new(big.Float).SetPrec(internalPrec).Set(one)       // p_0 = 1

	// --- Workspace variables (pre-allocate to reuse memory in the loop) ---
	// Variables for the next iteration's state
	aNext := new(big.Float).SetPrec(internalPrec)
	bNext := new(big.Float).SetPrec(internalPrec)
	tNext := new(big.Float).SetPrec(internalPrec)
	pNext := new(big.Float).SetPrec(internalPrec)
	// Other temporaries needed for calculations
	aDiff := new(big.Float).SetPrec(internalPrec) // Stores a_n - a_{n+1}
	tmp := new(big.Float).SetPrec(internalPrec)   // General temporary storage

	// --- Iteration loop ---
	// The number of iterations required grows logarithmically with precision.
	// log2(prec) gives a baseline; add a margin for safety and convergence.
	numIterations := int(math.Log2(float64(prec))) + 5

	for i := 0; i < numIterations; i++ {
		// a_{n+1} = (a_n + b_n) / 2
		aNext.Add(a, b)
		aNext.Mul(aNext, half) // Multiply by 0.5 is equivalent to dividing by 2

		// b_{n+1} = sqrt(a_n * b_n)
		tmp.Mul(a, b)   // tmp = a_n * b_n
		bNext.Sqrt(tmp) // b_{n+1} = sqrt(tmp)

		// Calculate the difference needed for t_{n+1} *before* updating a
		aDiff.Sub(a, aNext) // aDiff = a_n - a_{n+1}

		// p_{n+1} = 2 * p_n
		pNext.Mul(two, p)

		// t_{n+1} = t_n - p_n * (a_n - a_{n+1})^2
		aDiff.Mul(aDiff, aDiff) // aDiff = (a_n - a_{n+1})^2
		tmp.Mul(p, aDiff)       // tmp = p_n * (a_n - a_{n+1})^2
		tNext.Sub(t, tmp)       // tNext = t_n - tmp

		// Update state variables for the next iteration by reusing allocated memory.
		// a, b, t, p = a_{n+1}, b_{n+1}, t_{n+1}, p_{n+1}
		a.Set(aNext)
		b.Set(bNext)
		t.Set(tNext)
		p.Set(pNext)
	}

	// --- Final calculation ---
	// pi ≈ (a_n + b_n)^2 / (4 * t_n)
	// Note: After the loop, a, b, t hold the state corresponding to n+1
	pi := new(big.Float).SetPrec(internalPrec).Add(a, b) // pi = a_{n+1} + b_{n+1}
	pi.Mul(pi, pi)                                       // pi = (a_{n+1} + b_{n+1})^2
	tmp.Mul(four, t)                                     // tmp = 4 * t_{n+1}
	pi.Quo(pi, tmp)                                      // pi = pi / tmp

	// Set the final desired precision on the result and return.
	return pi.SetPrec(prec)
}

// Phi is the golden ratio φ constant with the given precision (in bits).
func Phi(prec uint) *big.Float {
	r := big.NewFloat(5).SetPrec(prec + 4)
	r.Sqrt(r)
	r.Add(r, big.NewFloat(1))
	r.Quo(r, big.NewFloat(2))
	return r.SetPrec(prec)
}
