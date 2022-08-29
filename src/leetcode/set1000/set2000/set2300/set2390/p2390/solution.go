package p2390

import "bytes"

func removeStars(s string) string {
	n := len(s)
	stack := make([]int, n)
	rem := make([]bool, n)
	var p int
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			rem[stack[p-1]] = true
			p--
			rem[i] = true
		} else {
			stack[p] = i
			p++
		}
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if !rem[i] {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
