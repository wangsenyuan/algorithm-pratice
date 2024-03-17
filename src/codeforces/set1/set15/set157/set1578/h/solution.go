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
	fmt.Println(solve(s))
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) int {
	// split it
	s = strings.ReplaceAll(s, ">", "")
	n := len(s)
	stack := make([]int, n)
	var top int
	pair := make([]int, n)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack[top] = i
			top++
		} else if s[i] == ')' {
			j := stack[top-1]
			pair[j] = i
			pair[i] = j
			top--
		}
	}

	var dfs func(l int, r int) int
	dfs = func(l int, r int) int {
		if l+1 == r {
			return 0
		}

		if pair[l] == r {
			return dfs(l+1, r-1)
		}
		// else a function type
		a := dfs(l, pair[l])
		// pair[l] + 1 = '->'
		b := dfs(pair[l]+2, r)
		return max(a+1, b)
	}
	return dfs(0, n-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
