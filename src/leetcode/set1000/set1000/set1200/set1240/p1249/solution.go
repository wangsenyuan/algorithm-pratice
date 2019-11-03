package p1249

import "bytes"

func minRemoveToMakeValid(s string) string {

	n := len(s)
	stack := make([]int, n)
	var p int
	valid := make([]bool, n)

	for i := 0; i < n; i++ {
		if s[i] == ')' {
			if p == 0 {
				// invalid
			} else {
				valid[stack[p-1]] = true
				valid[i] = true
				p--
			}
		} else if s[i] == '(' {
			stack[p] = i
			p++
		}
	}

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == ')' {
			if valid[i] {
				buf.WriteByte(s[i])
			}
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}
