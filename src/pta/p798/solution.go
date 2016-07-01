package main

import "fmt"

func main() {
	var a3, a2, a1, a0 float64

	var a, b float64

	fmt.Scanf("%f %f %f %f", &a3, &a2, &a1, &a0)
	fmt.Scanf("%f %f", &a, &b)

	var f = fx(a3, a2, a1, a0)

	for b-a > 0.0001 {
		//fmt.Printf("check %.2f, %.2f\n", a, b)
		c := a + (b-a)/2
		z := f(c)
		//fmt.Printf("f(%.2f) == %.6f\n", c, z)
		if z == 0.0 {
			fmt.Printf("%.2f\n", c)
			return
		}
		if sign(z) != sign(f(a)) {
			b = c
		} else {
			a = c
		}
	}

	fmt.Printf("%.2f\n", a)
}

func fx(a3, a2, a1, a0 float64) func(float64) float64 {
	return func(x float64) float64 {
		return ((a3*x+a2)*x+a1)*x + a0
	}
}

func sign(x float64) int {
	if x < 0 {
		return -1
	}
	return 1
}
