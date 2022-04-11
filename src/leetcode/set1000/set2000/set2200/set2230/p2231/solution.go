package p2231

import "strings"

func minimizeResult(expression string) string {
	it := strings.Index(expression, "+")

	calc := func(l, r int) int64 {
		a := parse(expression[:l])

		b := parse(expression[l:it]) + parse(expression[it+1:r])

		c := parse(expression[r:])

		if a == 0 {
			a = 1
		}
		if b == 0 {
			b = 1
		}
		if c == 0 {
			c = 1
		}

		return a * b * c
	}

	var best int64 = 1 << 30
	l, r := 0, 0

	n := len(expression)
	for i := 0; i < it; i++ {
		// add ( before i
		for j := it + 2; j <= n; j++ {
			// add ) after j
			tmp := calc(i, j)
			if tmp < best {
				best = tmp
				l = i
				r = j
			}
		}
	}

	res := expression[:l] + "(" + expression[l:r] + ")" + expression[r:]

	return res
}

func parse(s string) int64 {
	var a int64
	for i := 0; i < len(s); i++ {
		a = a*10 + int64(s[i]-'0')
	}
	return a
}
