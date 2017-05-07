package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	res := []byte(s)

	var j = 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverseWord(res, j, i)
			j = i + 1
		}
	}

	return string(res)
}

func reverseWord(str []byte, from, end int) {
	for i, j := from, end-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
