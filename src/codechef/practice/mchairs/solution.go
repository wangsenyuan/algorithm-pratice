package main

import "fmt"

var MOD int64 = 1000000007

func main() {
	var t, n int

	fmt.Scanf("%d", &t)
	
	for t > 0 {
		fmt.Scanf("%d", &n)

		m := fastExp(2, int64(n))

		m = (m - 1 + MOD) % MOD

		fmt.Println(m)

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
