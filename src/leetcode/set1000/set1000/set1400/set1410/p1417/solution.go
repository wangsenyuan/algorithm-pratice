package p1417

import (
	"bytes"
)

func reformat(s string) string {
	n := len(s)
	letters := make([]byte, 0, n)
	digits := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		} else {
			letters = append(letters, s[i])
		}
	}

	var buf bytes.Buffer
	var a, b int

	if len(digits) <= len(letters) {
		buf.WriteByte(letters[b])
		b++
	}

	for a < len(digits) && b < len(letters) {
		buf.WriteByte(digits[a])
		a++
		buf.WriteByte(letters[b])
		b++
	}

	if b < len(letters) {
		// previous one is b
		return ""
	}

	if a < len(digits) {
		buf.WriteByte(digits[a])
		a++
		if a < len(digits) {
			return ""
		}
	}

	return buf.String()
}
