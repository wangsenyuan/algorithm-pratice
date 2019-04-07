package p1024

import "bytes"

func removeOuterParentheses(S string) string {
	n := len(S)
	if n == 0 {
		return S
	}
	var left int
	var p int

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if S[i] == '(' {
			p++
		} else {
			p--
			if p == 0 {
				buf.WriteString(S[left+1 : i])
				left = i + 1
			}
		}
	}
	return buf.String()
}
