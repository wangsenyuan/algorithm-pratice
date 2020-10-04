package p1556

import (
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	var buf strings.Builder

	var l int

	for n > 0 {
		x := n % 10
		buf.WriteString(strconv.Itoa(x))
		l++
		if l == 3 {
			buf.WriteByte('.')
			l = 0
		}
		n /= 10
	}
	res := make([]byte, 0, buf.Len())
	m := buf.Len()
	s := buf.String()
	if s[m-1] == '.' {
		m--
	}
	for i := m - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	return string(res)
}
