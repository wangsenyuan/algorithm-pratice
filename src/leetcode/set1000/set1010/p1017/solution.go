package p1017

import (
	"bytes"
)

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	var buf bytes.Buffer

	for N != 0 {
		// first num that should be set
		if N&1 == 1 {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}

		N = -(N >> 1)
	}

	bs := buf.Bytes()

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}
