package main

import "fmt"

const N = 3000

var exp []float64

func init() {
	exp = make([]float64, N+1)
	exp[1] = 1
	for i := 2; i <= N; i++ {
		exp[i] = exp[i-1]/float64(i-1)*float64(i) + 1.0
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Printf("%.5f\n", solve(n))
	}
}

func solve(n int) float64 {
	if n > N {
		panic("wrong input")
	}
	return exp[n]
}
