package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)
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

func solve(a, b string) int {
	var w int
	n := len(a)
	m := len(b)
	var diff int
	for i := 0; i < m; i++ {
		if a[i] != b[i] {
			w ^= 1
		}
		if i > 0 && a[i] != a[i-1] {
			diff ^= 1
		}
	}
	var res int
	if w == 0 {
		res++
	}

	for i := m; i < n; i++ {
		if a[i] != a[i-1] {
			diff ^= 1
		}

		w ^= diff
		if w == 0 {
			res++
		}
		if a[i-m] != a[i-m+1] {
			diff ^= 1
		}
	}

	return res
}
