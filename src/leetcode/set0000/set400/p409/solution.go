package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	cnts := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		cnts[s[i]]++
	}

	sum := 0
	mid := 0
	for _, v := range cnts {
		if v%2 == 0 {
			sum += v
			continue
		}
		mid = 1
		sum += v - 1
	}

	return sum + mid
}
