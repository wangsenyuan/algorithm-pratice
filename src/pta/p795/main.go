package main

import "fmt"

func main() {
	var limit float64
	fmt.Scanf("%f", &limit)

	pi := calculate(limit)

	fmt.Printf("%.6f\n", pi)
}

func calculate(limit float64) float64 {

	var f func(prev float64, a, b, n int64) float64

	f = func(prev float64, a, b, n int64) float64 {
		diff := float64(a) * float64(n) / float64(b) / float64(2*n+1)
		//fmt.Printf("%f * %d / %d => %.6f\n", prev, n, 2*n+1, diff)
		if diff <= limit {
			return prev + diff
		}
		return f(prev+diff, a*n, b*(2*n+1), n+1)
	}

	return f(float64(1), 1, 1, 1) * float64(2)
}
