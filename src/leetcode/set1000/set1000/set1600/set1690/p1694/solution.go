package p1694

import (
	"bytes"
	"strings"
)

func reformatNumber(number string) string {
	str := number
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "-", "")
	var buf bytes.Buffer
	n := len(str)
	if n <= 3 {
		return str
	}
	for i := 0; i < n; {
		if i+4 == n {
			//2-2
			buf.WriteString(str[i : i+2])
			i += 2
			buf.WriteByte('-')
			buf.WriteString(str[i : i+2])
			i += 2
			continue
		}
		buf.WriteString(str[i:min(i+3, n)])
		i += 3
		if i < n {
			buf.WriteByte('-')
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
