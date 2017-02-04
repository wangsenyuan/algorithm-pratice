package main

import "fmt"

var M int64 = 1000000007

func main() {
	var n, x int
	fmt.Scanf("%d", &n)

	for n > 0 {
		fmt.Scanf("%d", &x)

		y := count(int64(x))

		fmt.Println(y)

		n--
	}
}

func count(n int64) int64 {
	// returns the answer for the problem
	// it is 2*(26+...+26^k), k = (n+1)/2
	// minus 26^k for odd n
	// 2*(26+...+26^k) = 52*(26^k-1)/25
	var pow func(x, t int64) int64

	pow = func(x, t int64) int64 {
		if t == 0 {
			return 1
		}

		y := pow(x, t/2) % M
		y = (y * y) % M
		if t%2 == 1 {
			y = x * y % M
		}
		return y
	}
	// inv25 is the inverse residue of 25 modulo MOD:
	// 25 * inv25 % MOD = 1
	inv25 := pow(25, M-2);
	k := (n + 1) / 2
	p26 := pow(26, k);                    // 26^k % MOD
	ans := 52 * (p26 - 1) % M * inv25 % M // 52 * (26^k-1) / 25 % MOD
	if n%2 == 1 {
		ans = (ans + M - p26) % M // subtract 26^k for odd n
	}
	return ans;
}
