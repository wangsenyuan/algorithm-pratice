package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int64
		var w int
		fmt.Scanf("%d %d", &n, &w)
		fmt.Println(solve(n, w))
		t--
	}
}

const MOD = 1000000007

func solve(n int64, w int) int64 {
	if n == 1 {
		return 0
	}
	x := pow(10, n-2)

	var y int64

	for d1 := 1; d1 < 10; d1++ {
		dn := d1 + w
		if dn >= 0 && dn < 10 {
			y++
		}
	}

	if y == 0 {
		return 0
	}

	return (x * y) % MOD
}

func pow(a, b int64) int64 {
	if b == 0 {
		return 1
	}
	c := pow(a, b/2)
	d := (c * c) % MOD
	if b%2 == 1 {
		return (d * a) % MOD
	}
	return d
}
