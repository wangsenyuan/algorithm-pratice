package main

import "fmt"

func main() {
	fmt.Println(isOneEditDistance("tache", "teacher"))
}

func isOneEditDistance(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) && s[i] == t[j] {
		i++
		j++
	}

	check := func(a, b int) bool {
		if a > len(s) || b > len(t) {
			return false
		}
		c, d := a, b
		for c < len(s) && d < len(t) && s[c] == t[d] {
			c++
			d++
		}

		if c < len(s) || d < len(t) {
			return false
		}
		if (c-d == 1) || (d-c == 1) {
			return true
		}
		if a == b && c == d {
			return true
		}

		return false
	}

	if check(i+1, j) || check(i, j+1) || check(i+1, j+1) {
		return true
	}
	return false
}
