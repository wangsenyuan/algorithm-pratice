package main

import "fmt"

func main() {
	fmt.Printf("%s is valid? %v\n", "()[]{}", isValid("()[]{}"))
}

func isValid(s string) bool {
	n := len(s)
	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack[p] = i
			p++
			continue
		}
		if p == 0 {
			return false
		}

		if s[i] == '}' && s[stack[p-1]] != '{' {
			return false
		}

		if s[i] == ')' && s[stack[p-1]] != '(' {
			return false
		}

		if s[i] == ']' && s[stack[p-1]] != '[' {
			return false
		}
		p--
	}

	return p == 0
}
