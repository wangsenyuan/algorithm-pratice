package p1604

import "bytes"

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)

	x := manachers(a)
	if x[0] == n-1 {
		return true
	}
	y := manachers(b)

	if y[0] == n-1 {
		return true
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if a[i] != b[j] {
			break
		}
		if i+1 > j-1 {
			return true
		}
		if x[i+1] == j-1 {
			return true
		}
		if y[i+1] == j-1 {
			return true
		}
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if b[i] != a[j] {
			break
		}
		if i+1 > j-1 {
			return true
		}
		if x[i+1] == j-1 {
			return true
		}
		if y[i+1] == j-1 {
			return true
		}
	}

	return false
}

func manachers(s string) []int {
	ss := modify(s)
	n := len(ss)
	P := make([]int, n)
	var c, r int

	for i := 0; i < n; i++ {
		mirror := 2*c - r

		if i < r {
			P[i] = min(r-i, P[mirror])
		}
		a := i + 1 + P[i]
		b := i - (1 + P[i])
		for a < n && b >= 0 && ss[a] == ss[b] {
			P[i]++
			a++
			b--
		}

		if i+P[i] > r {
			r = i + P[i]
			c = i
		}
	}
	f := make([]int, len(s))

	for i := 0; i < n; i++ {
		//l := P[i]
		// #a#a#b#
		a := i + P[i]
		b := i - P[i]
		a--
		b++
		if a > b || i%2 != 0 {
			f[b/2] = a / 2
		}
	}
	return f
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func modify(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte('#')
		buf.WriteByte(s[i])
	}
	buf.WriteByte('#')
	return buf.String()
}
