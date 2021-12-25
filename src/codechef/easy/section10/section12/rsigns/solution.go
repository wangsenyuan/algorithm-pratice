package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		t--
		var k int
		fmt.Scanf("%d", &k)
		fmt.Println(solve(k))
	}
}

const MOD = 1e9 + 7

func solve(k int) int {
	r := pow(2, k)
	return int((5 * int64(r)) % MOD)
}

func pow(a int, b int) int {
	res := int64(1)
	x := int64(a)
	for b > 0 {
		if b&1 == 1 {
			res *= x
			res %= MOD
		}
		x *= x
		x %= MOD
		b >>= 1
	}

	return int(res)
}
