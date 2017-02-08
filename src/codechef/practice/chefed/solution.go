package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var res = 0
	for i := 0; i <= 97; i++ {
		if f(n-i) == n {
			res++
		}
	}

	fmt.Println(res)
}

func f(x int) int {
	y := s(x)
	return x + y + s(y)
}

func s(x int) int {
	var y = 0
	for x > 0 {
		y += x % 10
		x /= 10
	}
	return y
}
