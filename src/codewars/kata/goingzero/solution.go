package main

import "fmt"

func Going(n int) float64 {

	res := 1.0

	for i := 2; i <= n; i++ {
		res = 1.0 + res/float64(i)
	}

	return res
}

func main() {
	for i := 1; i <= 8; i++ {
		fmt.Printf("%d => %.10f\n", i, Going(i))
	}
}
