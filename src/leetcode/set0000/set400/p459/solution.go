package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("abb"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))

}

func repeatedSubstringPattern(str string) bool {
	lps := computeLPS(str)
	fmt.Println(lps)
	n := len(lps)
	last := lps[n-1]
	if last > 0 && n%(n-last) == 0 {
		return true
	}
	return false
}

func computeLPS(str string) []int {
	lps := make([]int, len(str))
	length := 0
	i := 1
	for i < len(str) {
		if str[i] == str[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length > 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func repeatedSubstringPattern1(str string) bool {
	for i := 1; i <= len(str)/2; i++ {
		sub := str[:i]
		j := i
		for ; j+i <= len(str); j += i {
			if str[j:j+i] != sub {
				break
			}
		}
		if j == len(str) {
			return true
		}
	}

	return false
}
