package main

import (
	"sort"
	"fmt"
)

func topKFrequent(words []string, k int) []string {
	cnts := make(map[string]int)

	mx := 0
	for _, word := range words {
		cnts[word]++
		if cnts[word] > mx {
			mx = cnts[word]
		}
	}

	rev := make([][]string, mx+1)

	for word, cnt := range cnts {
		if rev[cnt] == nil {
			rev[cnt] = make([]string, 0, k)
		}
		rev[cnt] = append(rev[cnt], word)
	}

	res := make([]string, k)

	for i, j := mx, 0; i > 0 && j < k; i-- {
		ss := rev[i]
		if ss == nil || len(ss) == 0 {
			continue
		}

		sort.Strings(ss)

		for a := 0; a < len(ss) && j < k; a++ {
			res[j] = ss[a]
			j++
		}
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
