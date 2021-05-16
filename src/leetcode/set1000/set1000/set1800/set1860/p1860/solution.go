package p1860

import (
	"bytes"
	"strings"
)

func sortSentence(s string) string {
	ss := strings.Split(s, " ")
	mem := make(map[int]string)
	for _, x := range ss {
		l := len(x)
		y := int(x[l-1] - '1')
		mem[y] = x[:l-1]
	}

	var buf bytes.Buffer

	for k := 0; k < len(mem); k++ {
		buf.WriteString(mem[k])
		if k+1 < len(mem) {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}
