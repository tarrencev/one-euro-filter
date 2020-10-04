package oneeurofilter

import "math"

type OneEuroFilter struct {
	minCutoff float64
	beta      float64
	dCutoff   float64
	xPrev     float64
	dxPrev    float64
	tPrev     float64
}

func smoothingFactor(tE, cutoff float64) float64 {
	r := 2 * math.Pi * cutoff * tE
	return r / (r + 1)
}

func exponentialSmoothing(a, x, xPrev float64) float64 {
	return a*x + (1-a)*xPrev
}

// NewOneEuroFilter initializes a new One Euro Filter
func NewOneEuroFilter(t0, x0, dx0, minCutoff, beta, dCutoff float64) *OneEuroFilter {
	return &OneEuroFilter{
		minCutoff: minCutoff,
		beta:      beta,
		dCutoff:   dCutoff,
		xPrev:     x0,
		dxPrev:    dx0,
		tPrev:     t0,
	}
}

func (f *OneEuroFilter) Update(t, x float64) float64 {
	tE := t - f.tPrev

	// The filtered derivative of the signal.
	aD := smoothingFactor(tE, f.dCutoff)
	dx := (x - f.xPrev) / tE
	dxHat := exponentialSmoothing(aD, dx, f.dxPrev)

	// The filtered signal.
	cutoff := f.minCutoff + f.beta*math.Abs(dxHat)
	a := smoothingFactor(tE, cutoff)
	xHat := exponentialSmoothing(a, x, f.xPrev)

	// Memorize the previous values.
	f.xPrev = xHat
	f.dxPrev = dxHat
	f.tPrev = t

	return xHat
}
