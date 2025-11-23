package fft

import (
	"math"
	"math/cmplx"
)

// IFFT returns the inverse discrete Fourier transform of the input sequence X
func IFFT(X []complex128) []complex128 {
	return FFT(X, true)
}

// FFT returns the discrete Fourier transform of the input sequence x
// using the Cooley-Tukey FFT algorithm.
// len(x) must be a power of 2.
func FFT(x []complex128, inverse ...bool) []complex128 {
	N := len(x)
	if N == 0 {
		return nil
	}

	// out
	X := make([]complex128, N)
	copy(X, x)

	// bit-reverse
	for i := range N {
		j := reverse(i, N)
		if j > i {
			X[i], X[j] = X[j], X[i]
		}
	}

	sign := -1.0
	if len(inverse) > 0 && inverse[0] {
		sign = 1.0
	}

	for s := 1; s < log2(N)+1; s++ {
		m := 1 << s
		m2 := m >> 1

		a := complex(0, sign*2*math.Pi/float64(m))
		wm := cmplx.Exp(a)

		for k := 0; k < N; k += m {
			w := complex(1, 0)
			for j := range m2 {
				t, u := w*X[k+j+m2], X[k+j]
				X[k+j], X[k+j+m2] = u+t, u-t

				w *= wm
			}
		}
	}

	if len(inverse) > 0 && inverse[0] {
		for i := range N {
			X[i] /= complex(float64(N), 0)
		}
	}

	return X
}

func reverse(i, n int) int {
	bits := log2(n)

	var result int
	for b := range bits {
		if (i>>b)&1 != 1 {
			continue
		}

		result |= 1 << (bits - 1 - b)
	}

	return result
}

func log2(N int) int {
	var count int
	for N > 1 {
		N >>= 1
		count++
	}

	return count
}
