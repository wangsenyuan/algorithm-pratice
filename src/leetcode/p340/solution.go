package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := make(map[byte]int)
	ans := 0
	j := 0
	for i := range s {
		m[s[i]]++
		for len(m) > k {
			m[s[j]]--
			if m[s[j]] == 0 {
				delete(m, s[j])
			}
			j++
		}

		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	return ans
}
