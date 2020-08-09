package p1543

import "bytes"

func makeGood(s string) string {
	n := len(s)

	safe := func(i, j int) bool {
		x := ind(s[i])
		y := ind(s[j])
		return x != y || s[i] == s[j]
	}

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		if p == 0 || safe(stack[p-1], i) {
			stack[p] = i
			p++
		} else {
			p--
		}
	}
	var buf bytes.Buffer

	for i := 0; i < p; i++ {
		buf.WriteByte(s[stack[i]])
	}

	return buf.String()
}

func ind(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x - 'a')
	}
	return int(x - 'A')
}
