package main

import "fmt"

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"foo", "bar", "the"}

	rs := findSubstring(s, words)

	fmt.Printf("%v\n", rs)
}

func findSubstring(s string, words []string) []int {
	dict := make(map[string]int)

	for _, word := range words {
		dict[word]++
	}

	k := len(words[0])
	rs := make([]int, 0, 10)

	for i := 0; i < k; i++ {
		start := i
		tmp := make(map[string]int)

		for j := i; j+k <= len(s) && start+len(words)*k <= len(s); j += k {
			test := s[j : j+k]
			if dict[test] == 0 {
				start = j + k
				tmp = make(map[string]int)
				continue
			}

			for m := start; tmp[test] == dict[test]; m += k {
				start += k
				tmp[s[m:m+k]]--
			}

			if start+len(words)*k == j+k {
				rs = append(rs, start)
			}

			tmp[test]++
		}
	}

	return rs
}
