package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := solve(n)

	fmt.Println(res)
}

const N = 1_000_010

const mod = 1e9 + 7

var P [N]int
var pow [N]int

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func init() {
	P[0] = 1
	pow[0] = 1

	for i := 1; i < N; i++ {
		P[i] = mul(i, P[i-1])
		pow[i] = mul(2, pow[i-1])
	}
}

func solve(n int) int {
	return sub(P[n], pow[n-1])
}
