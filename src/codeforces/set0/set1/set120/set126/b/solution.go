package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) string {
	n := len(s)
	p := make([]int, n)
	// ok[i] 表示 长度为i的前缀，在中间也出现了
	ok := make([]bool, n+1)

	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[j] != s[i] {
			j = p[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		p[i] = j
		if i < n-1 && j != i+1 {
			ok[j] = true
		}
	}

	j := p[n-1]
	for j > 0 && !ok[j] {
		j = p[j-1]
	}

	if j == 0 {
		return "Just a legend"
	}
	return s[:j]
}
