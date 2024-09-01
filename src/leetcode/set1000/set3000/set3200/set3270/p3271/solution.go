package p3271

import (
	"bytes"
)

func stringHash(s string, k int) string {
	var res bytes.Buffer

	for i := 0; i < len(s); i += k {
		var code int
		for j := i; j < i+k; j++ {
			code += int(s[j] - 'a')
			if code >= 26 {
				code -= 26
			}
		}
		res.WriteByte(byte(code + 'a'))
	}

	return res.String()
}
