package main

import "fmt"

func main() {
	var x int
	var y float64
	fmt.Scanf("%d %f", &x, &y)
	res := solve(x, y)
	fmt.Printf("%.2f\n", res)
}

func solve(x int, y float64) float64 {
	if x%5 != 0 {
		return y
	}

	if float64(x)+0.5 > y {
		return y
	}

	return y - 0.5 - float64(x)
}
