package main

import "fmt"

func main() {
	res := generatePalindromes("aaaaaa")
	for _, x := range res {
		fmt.Println(x)
	}
}

func generatePalindromes(s string) []string {
	cnts := make(map[rune]int)
	for _, b := range s {
		cnts[b]++
	}

	if !canPermutePalindrome(cnts, len(s)) {
		return nil
	}

	center := ""
	half := ""

	for r, v := range cnts {
		if v%2 == 1 {
			center = string(r)
		}
		half += repeat(r, v/2)
	}

	rs := []rune(half)

	nn := factorial(len(rs))

	res := make([]string, 0, nn)
	used := make(map[string]bool)
	for i := 0; i < nn; i++ {
		nextPermutation(rs)
		x := string(rs) + center + string(reverse(rs))
		if used[x] {
			continue
		}
		used[x] = true
		res = append(res, x)
	}
	return res
}

func reverse(rs []rune) []rune {
	sr := make([]rune, len(rs))
	for i := 0; i < len(rs); i++ {
		sr[i] = rs[len(rs)-1-i]
	}
	return sr
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
func nextPermutation(s []rune) {
	n := len(s)
	i := n - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}

	if i >= 0 {
		j := n
		for j > i+1 && s[j-1] <= s[i] {
			j--
		}
		s[i], s[j-1] = s[j-1], s[i]
	}

	for k := 1; i+k < n-k; k++ {
		s[i+k], s[n-k] = s[n-k], s[i+k]
	}
}

func repeat(r rune, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += string(r)
	}
	return s
}

func canPermutePalindrome(cnts map[rune]int, n int) bool {

	odd := false
	for _, v := range cnts {
		if v%2 == 1 {
			if !odd {
				odd = true
			} else {
				return false
			}
		}
	}

	if n%2 == 1 {
		return odd == true
	}
	return true
}
