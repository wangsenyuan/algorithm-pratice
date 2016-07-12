package main

import "fmt"

func main() {
	s := "()"
	fmt.Printf("%d valid parentheses in %s\n", longestValidParentheses(s), s)
}

func longestValidParentheses(s string) int {
	stack := make([]int, len(s))
	p := 0
	r := 0
	for i, c := range s {
		if c == ')' && p > 0 && s[stack[p - 1]] == '(' {
			if p == 1 {
				r = max(r, i + 1)
			} else {
				r = max(r, i - stack[p - 2])
			}
			p--
		} else {
			stack[p] = i
			p++
		}

	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

