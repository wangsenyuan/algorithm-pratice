package main

import "fmt"

var MOD int64 = 1000000007

func main() {
	var t int
	var m int
	fmt.Scanf("%d", &t)
	for t > 0 {
		fmt.Scanf("%d", &m)

		n := int64(m + 1)

		n1 := n / 2
		n2 := n - n1

		res := fastExp(2, n1) + fastExp(2, n2) - 2

		res = res % MOD

		println(res)

		t--
	}
}

func fastExp(a, b int64) int64 {
	if b == 0 {
		return 1
	}

	c := fastExp(a, b/2)

	c = (c * c) % MOD

	if b&1 == 1 {
		c = (c * a) % MOD
	}

	return c
}
