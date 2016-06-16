package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abb"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	t := preProcess(s)
	n := len(t)

	P := make([]int, n, n)
	C := 0
	R := 0
	for i := 1; i < n-1; i++ {
		j := 2*C - i // equals to i' = C - (i-C)

		if R > i {
			P[i] = min(R-i, P[j])
		}
		// Attempt to expand palindrome centered at i
		for t[i+1+P[i]] == t[i-1-P[i]] {
			P[i]++
		}

		// If palindrome centered at i expand past R,
		// adjust center based on expanded palindrome.
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	// Find the maximum element in P.
	var maxLen = 0
	var centerIndex = 0
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	return s[(centerIndex-1-maxLen)/2 : (centerIndex-1-maxLen)/2+maxLen]
}

func preProcess(s string) string {
	rs := make([]rune, len(s)*2+3, len(s)*2+3)
	i := 0
	rs[i] = '^'

	for _, r := range s {
		i++
		rs[i] = '#'
		i++
		rs[i] = r
	}
	i++
	rs[i] = '#'
	i++
	rs[i] = '$'

	return string(rs)
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
