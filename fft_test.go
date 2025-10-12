package fft_test

import (
	"fmt"
	"math"

	"github.com/itsubaki/fft"
)

func ExampleCooleyTukey() {
	x := []complex128{1, 2, 3, 4, 5, 6, 7, 8}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (36.0000+0.0000i)
	// (-4.0000+9.6569i)
	// (-4.0000+4.0000i)
	// (-4.0000+1.6569i)
	// (-4.0000+0.0000i)
	// (-4.0000-1.6569i)
	// (-4.0000-4.0000i)
	// (-4.0000-9.6569i)
}

func ExampleCooleyTukey_step() {
	x := []complex128{1, 1, 1, 1, 0, 0, 0, 0}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (4.0000+0.0000i)
	// (1.0000-2.4142i)
	// (0.0000+0.0000i)
	// (1.0000-0.4142i)
	// (0.0000+0.0000i)
	// (1.0000+0.4142i)
	// (0.0000+0.0000i)
	// (1.0000+2.4142i)
}

func ExampleCooleyTukey_sin() {
	N := 8
	x := make([]complex128, N)
	for i := range N {
		x[i] = complex(math.Sin(2*math.Pi*float64(i)/float64(N)), 0)
	}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (0.0000+0.0000i)
	// (-0.0000-4.0000i)
	// (0.0000-0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (-0.0000+4.0000i)
}

func ExampleCooleyTukey_sincos() {
	N := 8
	x := make([]complex128, N)
	for i := range N {
		c := math.Cos(2 * math.Pi * float64(i) / float64(N))
		s := 0.5 * math.Sin(4*math.Pi*float64(i)/float64(N))
		x[i] = complex(c+s, 0)
	}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (-0.0000+0.0000i)
	// (4.0000-0.0000i)
	// (0.0000-2.0000i)
	// (0.0000-0.0000i)
	// (0.0000+0.0000i)
	// (0.0000-0.0000i)
	// (-0.0000+2.0000i)
	// (4.0000+0.0000i)
}

func ExampleCooleyTukey_impluse() {
	x := []complex128{1, 0, 0, 0, 0, 0, 0, 0}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
	// (1.0000+0.0000i)
}

func ExampleCooleyTukey_const() {
	x := []complex128{5, 5, 5, 5, 5, 5, 5, 5}

	X := fft.CooleyTukey(x)
	for _, v := range X {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (40.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
	// (0.0000+0.0000i)
}

func ExampleCooleyTukey_inverse() {
	x := []complex128{1, 2, 3, 4, 5, 6, 7, 8}

	X := fft.CooleyTukey(x)
	xi := fft.CooleyTukey(X, true)
	for _, v := range xi {
		fmt.Printf("%.4f\n", v)
	}

	// Output:
	// (1.0000+0.0000i)
	// (2.0000-0.0000i)
	// (3.0000+0.0000i)
	// (4.0000+0.0000i)
	// (5.0000+0.0000i)
	// (6.0000-0.0000i)
	// (7.0000-0.0000i)
	// (8.0000-0.0000i)
}
