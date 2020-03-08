package p1374

import (
	"bytes"
)

func generateTheString(n int) string {
	var buf bytes.Buffer
	if n&1 == 1 {
		for i := 0; i < n; i++ {
			buf.WriteByte('a')
		}
		return buf.String()
	}

	buf.WriteByte('a')
	for i := 1; i < n; i++ {
		buf.WriteByte('b')
	}
	return buf.String()
}
