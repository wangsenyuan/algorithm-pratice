package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("forgeeksskeegfor"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abb"))

}

func longestPalindrome(s string) string {
	n := len(s)

	fx := make([][]int, n, n)

	for i := 0; i < n; i++ {
		fx[i] = make([]int, n, n)
		fx[i][i] = 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp(fx, s, i, j)
		}
	}

	start := 0
	len := 1
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if fx[i][j] == 1 {
				if j-i+1 > len {
					start = i
					len = j - i + 1
				}
			}
		}
	}

	return s[start : start+len]
}

func dp(fx [][]int, s string, i, j int) int {
	if i >= j {
		return 1
	}

	if fx[i][j] != 0 {
		return fx[i][j]
	}

	if s[i] != s[j] {
		fx[i][j] = -1
		return fx[i][j]
	}

	fx[i][j] = dp(fx, s, i+1, j-1)
	return fx[i][j]
}
