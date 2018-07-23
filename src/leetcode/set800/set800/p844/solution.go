package p844

import "bytes"

func backspaceCompare(S string, T string) bool {
	s := process(S)
	t := process(T)

	return s == t
}

func process(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if buf.Len() > 0 {
				buf.Truncate(buf.Len() - 1)
			}
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
