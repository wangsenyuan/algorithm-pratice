package p1108

import (
	"bytes"
)

func defangIPaddr(address string) string {
	var buf bytes.Buffer

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			buf.WriteString("[.]")
		} else {
			buf.WriteByte(address[i])
		}
	}

	return buf.String()
}
