package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	ps := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		ps[p[i]]++
	}
	ss := make(map[byte]int)
	var ans []int
	j := 0
	for i := 0; i <= len(s); i++ {
		if i < len(s) {
			if _, has := ps[s[i]]; !has {
				j = i + 1
				ss = make(map[byte]int)
				continue
			}
			ss[s[i]]++
			if ss[s[i]] > ps[s[i]] {
				for j < i && ss[s[i]] > ps[s[i]] {
					ss[s[j]]--
					j++
				}
			}
		}

		if reflect.DeepEqual(ss, ps) {
			ans = append(ans, j)
			ss[s[j]]--
			j++
		}
	}

	return ans
}
