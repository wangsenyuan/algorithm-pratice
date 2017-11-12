package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Sort(ByLen(words))

	dict := make(map[string]bool)
	var res []string
	for _, word := range words {
		if canConcatenate(word, dict) {
			res = append(res, word)
		}
		dict[word] = true
	}

	return res
}

func canConcatenate(word string, dict map[string]bool) bool {
	if len(dict) == 0 {
		return false
	}
	n := len(word)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			if dict[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
