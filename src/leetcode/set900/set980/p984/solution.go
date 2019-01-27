package p984

import "bytes"

func strWithout3a3b(A int, B int) string {
	var buf bytes.Buffer
	n := A + B
	if A == B {
		for i := 0; i < n; i += 2 {
			buf.WriteByte('a')
			buf.WriteByte('b')
		}
		return buf.String()
	}

	var first, second byte

	if A > B {
		first, second = 'a', 'b'
	} else {
		first, second = 'b', 'a'
		A, B = B, A
	}
	//A > B
	//aab
	n = A + B

	for i := 0; i < n && A > B; {
		if A > 0 {
			buf.WriteByte(first)
			A--
			i++
		}
		if A > 0 {
			buf.WriteByte(first)
			A--
			i++
		}
		if B > 0 {
			buf.WriteByte(second)
			B--
			i++
		}
	}
	if A > 0 {
		for i := 0; i < A; i++ {
			buf.WriteByte(first)
			buf.WriteByte(second)
		}
	}

	return buf.String()
}
