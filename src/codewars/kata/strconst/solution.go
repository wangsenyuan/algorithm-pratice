package main

import (
	"fmt"
)

func StringConstructing(a, s string) int {

	var construct func(t string, pos int, ans int) int

	construct = func(t string, pos int, ans int) int {
		//fmt.Println(t)
		i := pos
		for i < len(t) && i < len(s) && t[i] == s[i] {
			i++
		}

		if i == len(t) && i == len(s) {
			return ans
		}

		if i == len(t) {
			// all is same, just append a
			return construct(t+a, i, ans+1)
		}

		if i == len(s) {
			// remove all the extra characters
			return construct(t[:i], i, ans+len(t)-len(s))
		}

		return construct(t[:i]+t[i+1:], i, ans+1)
	}

	return construct("", 0, 0)
}

func main() {
	fmt.Println(StringConstructing("aba", "abbabba"))
	fmt.Println(StringConstructing("aba", "abaa"))
	fmt.Println(StringConstructing("aba", "a"))
	fmt.Println(StringConstructing("a", "a"))
	fmt.Println(StringConstructing("a", "aaa"))
	fmt.Println(StringConstructing("abcdefgh", "hgfedcba"))
	fmt.Println(StringConstructing("sxdfcgdgxdfgdxxf", "xgdfsxgdxgfsgdfxgfdgfgdfgdgsdfxgfdxgdfxgdgdfgdfxgsdxgdfxgfgsxfgdfgsdfxgdfxgsgsfgxsgsdxgsdfxgsgsdfxgsdfx"))
}
