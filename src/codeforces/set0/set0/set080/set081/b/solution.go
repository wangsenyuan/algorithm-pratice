package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	res := solve(s)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) string {
	var buf bytes.Buffer

	var last byte = ','

	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}
		// number
		j := i
		for i < len(s) && isDigit(s[i]) {
			i++
		}
		if j < i {
			if isDigit(last) {
				buf.WriteByte(' ')
			}

			buf.WriteString(s[j:i])
			last = s[i-1]
			continue
		}
		if s[i] == '.' {
			if last != ',' {
				buf.WriteByte(' ')
			}
			buf.WriteString("...")
			i += 3
			last = '.'
			continue
		}

		last = ','
		buf.WriteString(", ")
		i++
	}

	res := buf.String()

	if res[len(res)-1] == ' ' {
		return res[:len(res)-1]
	}
	return res
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
