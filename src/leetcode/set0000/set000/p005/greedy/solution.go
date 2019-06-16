package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("forgeeksskeegfor"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abb"))

}

func longestPalindrome(s string) string {
	ls := ""

	for i := 0; i < len(s); i++ {
		if len(ls) == 0 {
			ls = s[i : i+1]
		}
		a := expand(s, i, i+1)
		if len(a) > len(ls) {
			ls = a
		}

		b := expand(s, i-1, i+1)
		if len(b) > len(ls) {
			ls = b
		}
	}

	return ls

}

func expand(s string, lo, hi int) string {
	ls := ""
	ln := 0
	for lo >= 0 && hi < len(s) {
		if s[lo] != s[hi] {
			break
		}

		if hi-lo+1 > ln {
			ln = hi - lo + 1
			ls = s[lo : hi+1]
		}
		lo--
		hi++
	}
	return ls
}
