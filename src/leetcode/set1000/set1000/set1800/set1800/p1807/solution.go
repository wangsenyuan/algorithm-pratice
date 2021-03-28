package p1807

import "bytes"

func evaluate(s string, knowledge [][]string) string {
	mem := make(map[string]string)
	for _, x := range knowledge {
		mem[x[0]] = x[1]
	}

	var buf bytes.Buffer

	for i := 0; i < len(s); {
		if s[i] == '(' {
			j := i + 1
			for s[j] != ')' {
				j++
			}
			k := s[i+1 : j]
			if v, found := mem[k]; found {
				buf.WriteString(v)
			} else {
				buf.WriteByte('?')
			}
			i = j + 1
		} else {
			buf.WriteByte(s[i])
			i++
		}
	}

	return buf.String()
}
