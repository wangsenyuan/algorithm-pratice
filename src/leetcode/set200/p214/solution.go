package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}

func shortestPalindrome(s string) string {

	expend := func(a, b int) (int, int) {
		for a >= 0 && b < len(s) && s[a] == s[b] {
			a, b = a-1, b+1
		}
		return a + 1, b - 1
	}
	l := 0
	for i := 0; i < len(s); i++ {
		a, b := expend(i, i+1)
		if a == 0 && b - a + 1 > l {
			l = b - a + 1
		}
		a, b = expend(i-1, i+1)
		if a == 0 && b - a + 1 >l {
			l = b - a + 1
		}
	}

	var buf bytes.Buffer

	for i := len(s) - 1; i >= l; i-- {
		buf.WriteByte(s[i])
	}

	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
	}
	return buf.String()
}
