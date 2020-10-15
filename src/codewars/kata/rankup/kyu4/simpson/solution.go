package main

import (
	"math"
	"fmt"
)

func Simpson(n int) float64 {
	// your code
	// a = 0, b = pi

	f := func(x float64) float64 {
		y := math.Sin(x)
		return 1.5 * y * y * y
	}
	y0 := f(0.0)
	y1 := f(math.Pi)

	h := math.Pi / float64(n)

	y2 := 0.0

	for i := 1; i <= n/2; i++ {
		y2 += f(float64(2*i-1) * h)
	}

	y2 *= 4.0

	y3 := 0.0

	for i := 1; i < n/2; i++ {
		y3 += f(float64(2 * i) * h)
	}

	y3 *= 2.0

	return math.Pi * (y0 + y1 + y2 + y3) / (3.0 * float64(n))
}

func main() {
	fmt.Println(Simpson(290))
	fmt.Println(Simpson(72))
	fmt.Println(Simpson(252))
	fmt.Println(Simpson(40))
}
