package p2434

import "bytes"

func robotWithString(s string) string {
	// 其实是一个stack
	n := len(s)

	min_pos := make([]byte, n+1)
	min_pos[n] = 'z'

	for i := n - 1; i >= 0; i-- {
		min_pos[i] = s[i]
		if i+1 < n && s[i] > min_pos[i+1] {
			min_pos[i] = min_pos[i+1]
		}
	}

	stack := make([]byte, n)
	var p int

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		stack[p] = s[i]
		p++

		for p > 0 && stack[p-1] <= min_pos[i+1] {
			buf.WriteByte(stack[p-1])
			p--
		}
	}

	for p > 0 {
		buf.WriteByte(stack[p-1])
		p--
	}

	return buf.String()
}
