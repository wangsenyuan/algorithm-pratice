package p5541

func countSubstrings(s string, t string) int {
	m := len(s)
	n := len(t)

	count := func(a, b int) int {
		var diff int
		var res int
		for a < m && b < n {
			if s[a] != t[b] {
				diff++
			}
			if diff == 1 {
				res++
			}
			if diff > 1 {
				break
			}
			a++
			b++
		}
		return res
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += count(i, j)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countDigits(num int) int {
	var cnt int
	for num > 0 {
		cnt++
		num -= num & -num
	}
	return cnt
}
