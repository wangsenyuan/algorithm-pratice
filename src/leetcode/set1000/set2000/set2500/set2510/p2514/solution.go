package p2514

import "math"

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	g := gcd(divisor1, divisor2)

	// 假设最终的结果是x
	// 那么 <= x, 整除 div1 的， 算入 集合2， <=x, 整除 div2的部分， 算入集合1，
	lcm := int64(divisor1) / int64(g) * int64(divisor2)

	check := func(num int) bool {
		x := num / divisor1
		y := num / divisor2
		z := int(int64(num) / lcm)
		// 既不能被div1整除，又不能div2整除的部分
		a := num - x - y + z
		var cnt1, cnt2 int

		if divisor2%divisor1 != 0 {
			cnt1 = y - z
		}
		if divisor1%divisor2 != 0 {
			cnt2 = x - z
		}

		diff1 := max(0, uniqueCnt1-cnt1)
		diff2 := max(0, uniqueCnt2-cnt2)

		return diff1+diff2 <= a
	}

	var l, r int = 0, math.MaxInt64

	for l < r {
		mid := (l + r) / 2

		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
