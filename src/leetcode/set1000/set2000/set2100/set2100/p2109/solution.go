package p2109

import "bytes"

func addSpaces(s string, spaces []int) string {
	var buf bytes.Buffer

	var p int
	for i := 0; i < len(s); i++ {
		if p < len(spaces) && spaces[p] == i {
			buf.WriteByte(' ')
			p++
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
