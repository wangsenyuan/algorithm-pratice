package p3348

import (
	"strings"
)

func smallestNumber(num string, t int64) string {
	target := int(t)

	divs := []int{2, 3, 5, 7}

	tmp := target
	for _, d := range divs {
		for tmp%d == 0 {
			tmp /= d
		}
	}
	if tmp > 1 {
		return "-1"
	}

	const add = 50

	num = strings.Repeat("0", add) + num
	n := len(num)

	vis := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make(map[int]bool)
	}

	var dfs func(i int, x int, flag bool) bool

	buf := make([]byte, n)
	var top int

	dfs = func(i int, x int, flag bool) bool {
		if i == n {
			return x == 1
		}
		if flag {
			if vis[i][x] {
				return false
			}
			vis[i][x] = true
		}

		v := num[i]

		var start int

		if !flag {
			if v == '0' {
				if i < add {
					start = 0
				} else {
					start = 1
				}
			} else {
				start = int(v - '0')
			}
		} else {
			start = 1
		}

		for nv := start; nv < 10; nv++ {
			buf[top] = byte(nv + '0')
			top++
			nx := x
			if nv > 0 {
				nx = x / gcd(x, nv)
			}
			nf := flag || nv > int(v-'0')
			if dfs(i+1, nx, nf) {
				return true
			}
			top--
		}

		return false
	}

	dfs(0, target, false)

	for i := 0; i < n; i++ {
		if buf[i] != '0' {
			return string(buf[i:])
		}
	}

	return "-1"
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
