package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	x := readString(reader)
	res := solve(s, x)

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

func solve(s string, x string) int {
	n := len(s)
	m := len(x)
	// n <= m
	pref := make([]int, m+1)
	for i := 0; i < m; i++ {
		pref[i+1] = pref[i]
		if x[i] == '1' {
			pref[i+1]++
		}
	}
	// 01   111
	var res int
	for i := 0; i < n; i++ {
		tmp := pref[m-(n-i)+1] - pref[i]
		if s[i] == '1' {
			res += (m - (n - i) + 1 - i) - tmp
		} else {
			res += tmp
		}
	}

	return res
}
