package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var res []string
	last := strings.LastIndexByte(s, '@')
	if last < 0 {
		return "No solution"
	}
	last = len(s) - last
	for len(s) > 0 {
		if s[0] == '@' {
			return "No solution"
		}
		i := strings.IndexByte(s, '@')
		if i < 0 || i == len(s)-1 {
			return "No solution"
		}
		if len(s)-i == last {
			res = append(res, s)
			s = ""
		} else {
			if s[i+1] == '@' {
				return "No solution"
			}
			res = append(res, s[:i+2])
			s = s[i+2:]
		}
	}
	return strings.Join(res, ",")
}
