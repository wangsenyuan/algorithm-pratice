package p5617

import "bytes"

func interpret(command string) string {
	var buf bytes.Buffer

	for i := 0; i < len(command); {
		if command[i] == 'G' {
			buf.WriteByte(command[i])
			i++
			continue
		}
		// comm[i] == '('
		if command[i+1] == ')' {
			buf.WriteByte('o')
			i += 2
			continue
		}
		buf.WriteString("al")
		i += 4
	}

	return buf.String()
}
