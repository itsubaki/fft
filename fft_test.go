package fft_test

import (
	"fmt"
	"math"

	"github.com/itsubaki/fft"
)

func ExampleForward() {
	x := []complex128{1, 2, 3, 4, 5, 6, 7, 8}

	X := fft.Forward(x)
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

func ExampleForward_step() {
	x := []complex128{1, 1, 1, 1, 0, 0, 0, 0}

	X := fft.Forward(x)
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

func ExampleForward_sin() {
	N := 8
	x := make([]complex128, N)
	for i := range N {
		a := 2 * math.Pi * float64(i) / float64(N)
		x[i] = complex(math.Sin(a), 0)
	}

	X := fft.Forward(x)
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

func ExampleForward_sincos() {
	N := 8
	x := make([]complex128, N)
	for i := range N {
		a := math.Pi * float64(i) / float64(N)
		cos := math.Cos(2 * a)
		sin := math.Sin(4 * a)

		x[i] = complex(cos+0.5*sin, 0)
	}

	X := fft.Forward(x)
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

func ExampleForward_impluse() {
	x := []complex128{1, 0, 0, 0, 0, 0, 0, 0}

	X := fft.Forward(x)
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

func ExampleForward_const() {
	x := []complex128{5, 5, 5, 5, 5, 5, 5, 5}

	X := fft.Forward(x)
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

func ExampleForward_inverse() {
	x := []complex128{1, 2, 3, 4, 5, 6, 7, 8}

	xi := fft.Inverse(fft.Forward(x))
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

func ExampleForward_zero() {
	X := fft.Forward([]complex128{})
	fmt.Println(X)

	// Output:
	// []
}
