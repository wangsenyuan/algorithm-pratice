package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var p float64

	fmt.Scan(&n, &p)

	res := solve(n, p)

	fmt.Println(res)

}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}

	res := 1
	for i := 0; i < r; i++ {
		res *= (n - i)
		res /= (i + 1)
	}

	return res
}

const eps = 1e-8

func solve(n int, p float64) int {

	// 这个tot 不对
	tot := nCr(n, 3) * 2

	check := func(m int) float64 {
		// 选中m个
		res := nCr(n-m, 2) * nCr(m, 1)
		res += nCr(n-m, 1) * nCr(m, 2) * 2
		res += nCr(n-m, 0) * nCr(m, 3) * 2

		ans := float64(res) / float64(tot)

		return ans
	}

	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		tmp := check(mid)
		if tmp >= p || math.Abs(tmp-p) < eps {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
