package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	for i, j := 0, 0; i < len(s); i, j = i+1, j+1 {
		for j < len(t) {
			if s[i] == t[j] {
				break
			}
			j++
		}
		if j == len(t) {
			return false
		}
	}

	return true
}
