package main

import "fmt"

func main() {
	fmt.Println(numDecodings("123"))
}

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	a, b := 1, 1

	for i := 0; i < n; i++ {
		c := a
		if s[i] == '0' {
			a = 0
		}

		if i >= 1 && canDecode(s[i-1], s[i]) {
			a += b
		}
		b = c
	}

	return a
}

func canDecode(a, b byte) bool {
	if a == '1' {
		return true
	}

	if a == '2' && b <= '6' {
		return true
	}
	return false
}
