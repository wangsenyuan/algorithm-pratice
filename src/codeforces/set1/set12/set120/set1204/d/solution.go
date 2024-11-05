package main

import (
	"bufio"
	"fmt"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)

	fmt.Println(res)
}

func solve(s string) string {
	n := len(s)

	stack := make([]int, n)
	var top int

	buf := []byte(s)

	for i, x := range []byte(s) {
		if x == '0' {
			if top > 0 && s[stack[top-1]] == '1' {
				top--
				continue
			}
		}
		stack[top] = i
		top++
	}

	for top > 0 && s[stack[top-1]] == '1' {
		buf[stack[top-1]] = '0'
		top--
	}

	return string(buf)
}
