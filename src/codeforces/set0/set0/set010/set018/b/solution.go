package main

import "fmt"

func main() {
	var n, d, m, l int
	fmt.Scan(&n, &d, &m, &l)
	res := solve(n, d, m, l)
	fmt.Println(res)
}

func solve(n int, d int, m int, l int) int {

	best := ((n-1)*m + l + 1 + d - 1) / d * d
	for i := l + 1; i < m; i++ {
		// (k * m + i) % d = 0
		x := i % d
		if x != 0 {
			x = d - x
		}
		// k * m % d = x
		g := gcd(m, d)
		if x%g != 0 {
			continue
		}
		// a * b % d = i
		u := m / g
		v := d / g
		x /= g
		// k * u % v = x
		k := x * modularInverse(u, v) % v
		best = min(k*m+i, best)
	}

	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// 扩展欧几里得算法，用于求解方程 ax + by = gcd(a, b)
func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x, y := extendedGCD(b, a%b)
	return gcd, y, x - (a/b)*y
}

// 求a模m的乘法逆元
func modularInverse(a, m int) int {
	g, x, _ := extendedGCD(a, m)
	if g != 1 {
		return -1 // a和m不互质，不存在模逆元
	}
	return (x%m + m) % m // 确保返回的是最小正整数
}
