package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d", &x)
	y := reverse(x)
	fmt.Printf("%d\n", y)
}

func reverse(x int) int {
	var y = 0
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}

	return y
}
