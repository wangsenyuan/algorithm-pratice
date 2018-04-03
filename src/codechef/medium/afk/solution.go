package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var A, B, C int64
		fmt.Scanf("%d %d %d", &A, &B, &C)
		fmt.Println(solve(A, B, C))
		t--
	}
}

func solve(A, B, C int64) int64 {
	d := abs(2*B - A - C)
	x := d / 2
	y := d - 2*x
	return x + y
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
