package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct1("bacc"))
}

func lengthOfLongestSubstringTwoDistinct1(s string) int {
	res := 0
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			j++
		}

		k := j
		for k < len(s) && (s[k] == s[i] || s[k] == s[j]) {
			k++
		}

		if k-i > res {
			res = k - i
		}
		i = j
	}

	return res
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	cnts := make(map[byte]int)
	res := 0
	from := 0
	for i := 0; i < len(s); i++ {
		b := s[i]

		cnts[b]++

		for len(cnts) > 2 {
			d := s[from]
			c := cnts[d]
			if c == 1 {
				delete(cnts, d)
			} else {
				cnts[d]--
			}
			from++
		}

		if i-from+1 > res {
			res = i - from + 1
		}
	}

	return res
}
