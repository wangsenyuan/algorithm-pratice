package main

import "fmt"

func countSubstrings(s string) int {

	palindromic := func(x, y int) int {
		for x > 0 && y < len(s)-1 && s[x-1] == s[y+1] {
			x--
			y++
		}
		return y - x + 1
	}

	var res int
	for i := 0; i < len(s); i++ {
		a := palindromic(i, i)
		res += a/2 + 1
		if i < len(s)-1 && s[i] == s[i+1] {
			b := palindromic(i, i+1)
			res += b / 2
		}
	}

	return res
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
