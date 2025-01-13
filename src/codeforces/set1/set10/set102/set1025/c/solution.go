package main

import (
	"bufio"
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

func solve(s string) int {
	var res int
	n := len(s)

	check := func(i int, j int) bool {
		a := s[i%n]
		b := s[j%n]
		if (j-i)%2 == 0 {
			return a == b
		}
		return a != b
	}

	for i := 0; i < 2*n; {
		j := i
		for i < 2*n && check(j, i) {
			i++
		}
		res = max(res, i-j)
	}

	return min(res, n)
}
