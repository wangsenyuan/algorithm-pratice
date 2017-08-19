package main

import (
	"math"
	"fmt"
)

func Doubles(maxk, maxn int) float64 {
	// your code

	v := func(k, n int) float64 {
		return 1.0 / (float64(k) * math.Pow(float64(n+1), float64(2*k)))
	}

	u := func(k, n int) float64 {
		sum := 0.0

		for i := 1; i <= n; i++ {
			sum += v(k, i)
		}
		return sum
	}

	s := func(k, n int) float64 {
		sum := 0.0

		for i := 1; i <= k; i++ {
			sum += u(i, n)
		}

		return sum
	}

	return s(maxk, maxn)
}

func main() {
	fmt.Println(Doubles(1, 10))
	fmt.Println(Doubles(10, 1000))
	fmt.Println(Doubles(10, 10000))
}
