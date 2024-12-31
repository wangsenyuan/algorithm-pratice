package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)
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

func solve(s string, t string) string {

	pos := make([]int, 26)
	for i := len(s) - 1; i > 0; i-- {
		x := int(s[i] - 'a')
		pos[x] = i
	}
	at := -1
	var l int
	n := len(t)
	for i := n - 2; i >= 0; i-- {
		x := int(t[i] - 'a')
		if pos[x] == 0 {
			continue
		}
		if l == 0 || pos[x]+n-i < l {
			l = pos[x] + n - i
			at = i
		}
	}

	if at < 0 {
		return "-1"
	}
	x := int(t[at] - 'a')
	ans := s[:pos[x]] + t[at:]
	return ans
}
