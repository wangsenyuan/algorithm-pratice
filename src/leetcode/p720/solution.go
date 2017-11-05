package main

import (
	"sort"
	"fmt"
)

func longestWord(words []string) string {
	wordMap := make(map[string]bool)

	for _, word := range words {
		wordMap[word] = true
	}

	sort.Sort(SS(words))

	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		can := true
		for j := len(word) - 1; j > 0; j-- {
			if found, _ := wordMap[word[:j]]; !found {
				can = false
				break
			}
		}

		if can {
			return word
		}
	}
	return ""
}

type SS []string

func (this SS) Len() int {
	return len(this)
}

func (this SS) Less(i, j int) bool {
	return len(this[i]) < len(this[j]) || (len(this[i]) == len(this[j]) && this[i] > this[j])
}

func (this SS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(longestWord([]string{"m", "mo", "moc", "moch", "mocha", "l", "la", "lat", "latt", "latte", "c", "ca", "cat"}))
}
