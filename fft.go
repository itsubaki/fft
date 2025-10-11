package fft

import (
	"math"
	"math/cmplx"
)

type FFT func(x []complex128) []complex128

// CooleyTukey returns the discrete Fourier transform of the input sequence x
func CooleyTukey(x []complex128) []complex128 {
	N := len(x)
	if N < 2 {
		return x
	}

	even, odd := make([]complex128, N/2), make([]complex128, N/2)
	for i := range N / 2 {
		even[i], odd[i] = x[2*i], x[2*i+1]
	}
	Feven, Fodd := CooleyTukey(even), CooleyTukey(odd)

	X := make([]complex128, N)
	for k := range N / 2 {
		a := -2i * math.Pi * complex(float64(k), 0) / complex(float64(N), 0)
		t := cmplx.Exp(a) * Fodd[k]

		X[k], X[k+N/2] = Feven[k]+t, Feven[k]-t
	}

	return X
}
