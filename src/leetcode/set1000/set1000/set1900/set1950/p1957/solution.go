package p1957

import "bytes"

func makeFancyString(s string) string {
	var buf bytes.Buffer
	for i, j := 1, 0; i <= len(s); i++ {
		if i == len(s) || s[i] != s[i-1] {
			cnt := min(i-j, 2)
			for k := 0; k < cnt; k++ {
				buf.WriteByte(s[i-1])
			}
			j = i
		}
	}

	return buf.String()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
