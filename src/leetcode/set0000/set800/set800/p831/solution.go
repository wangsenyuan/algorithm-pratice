package p831

import (
	"strings"
	"bytes"
)

func maskPII(S string) string {
	if strings.Contains(S, "@") {
		return maskEmail(strings.ToLower(S))
	}
	return maskPhone(S)
}
func maskPhone(S string) string {
	var x int
	nums := make([]byte, 13)
	for i := 0; i < len(S); i++ {
		if isDigit(S[i]) {
			nums[x] = S[i]
			x++
		}
	}
	nums = nums[:x]
	var buf bytes.Buffer
	buf.WriteByte(nums[x-1])
	x--
	buf.WriteByte(nums[x-1])
	x--
	buf.WriteByte(nums[x-1])
	x--
	buf.WriteByte(nums[x-1])
	x--

	for i := 0; i < 2; i++ {
		buf.WriteByte('-')
		for j := 0; j < 3; j++ {
			buf.WriteByte('*')
			x--
		}
	}

	if x > 0 {
		buf.WriteByte('-')
		for x > 0 {
			buf.WriteByte('*')
			x--
		}
		buf.WriteByte('+')
	}

	ss := buf.Bytes()
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	return string(ss)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func maskEmail(S string) string {
	x := strings.Index(S, "@")
	first := S[:x]
	a := first[0]
	b := first[len(first)-1]
	var buf bytes.Buffer
	buf.WriteByte(a)
	buf.WriteString("*****")
	buf.WriteByte(b)
	buf.WriteString(S[x:])
	return buf.String()
}
