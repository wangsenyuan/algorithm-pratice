package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(solve(n, m))
}

const mod = 998244353

func add(a, b int) int {
	a += b

	for a >= mod {
		a -= mod
	}

	for a < 0 {
		a += mod
	}

	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, m int) int {
	var ans int
	for i := 1; i <= n; i++ {
		ans = add(ans, pow(m%mod, i))
	}

	cnt := 1
	cur := 1

	for i := 1; i <= n; i++ {
		if cur > m {
			break
		}
		if prime(i) {
			cur *= i
		}
		cnt = mul(cnt, m/cur%mod)
		ans = sub(ans, cnt)
	}
	return ans
}

func prime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
