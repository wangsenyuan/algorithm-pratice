package main

import "fmt"

func main() {
	fmt.Printf("%s is valid? %v\n", "()[]{}", isValid("()[]{}"))
}

func isValid(s string) bool {
	for i := 0; i < len(s); {
		j := pairedIndex(s, i+1, s[i], 1)
		if j == len(s) {
			return false
		}
		if !isValid(s[i+1 : j]) {
			return false
		}
		i = j + 1
	}

	return true
}

func pairedIndex(s string, i int, x byte, level int) int {
	if i >= len(s) {
		return i
	}
	if s[i] == x {
		return pairedIndex(s, i+1, x, level+1)
	}

	if pair(s[i], x) && level == 1 {
		return i
	}

	if pair(s[i], x) {
		return pairedIndex(s, i+1, x, level-1)
	}

	return pairedIndex(s, i+1, x, level)
}

func pair(y, x byte) bool {
	if x == '(' {
		return y == ')'
	}

	if x == '[' {
		return y == ']'
	}

	if x == '{' {
		return y == '}'
	}

	return false
}
