package main

import "fmt"

func validPalindrome(s string) bool {
	n := len(s)

	i, j := 0, n-1

	for i < j && s[i] == s[j] {
		i++
		j--
	}
	if i >= j {
		return true
	}

	a, b := i+1, j
	for a < b && s[a] == s[b] {
		a++
		b--
	}
	if a >= b {
		return true
	}

	a, b = i, j-1
	for a < b && s[a] == s[b] {
		a++
		b--
	}

	return a >= b
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
