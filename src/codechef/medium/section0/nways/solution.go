package main

import "fmt"

const MAX = 2000012

const MOD = 1000000007

var fact []int64
var invFact []int64

func init() {
	fact = make([]int64, MAX)
	invFact = make([]int64, MAX)
	fact[0] = 1
	invFact[0] = 1
	invFact[1] = 1
	for i := 1; i < MAX; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	for i := 2; i < MAX; i++ {
		invFact[i] = int64(MOD-MOD/i) * invFact[MOD%i] % MOD
	}
	for i := 1; i < MAX; i++ {
		invFact[i] = (invFact[i-1] * invFact[i]) % MOD
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)
		fmt.Println(solve(n, k))
		t--
	}
}

func solve(n, k int) int64 {
	x := n + k + 1
	y := k + 2
	ans := nCr(x, y)

	return ((2*ans)%MOD - int64(n) + MOD) % MOD
}

func nCr(n, r int) int64 {
	return (((fact[n] * invFact[r]) % MOD) * invFact[n-r]) % MOD
}
