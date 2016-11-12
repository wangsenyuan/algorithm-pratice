package main

import "fmt"

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))
}

func palindromePairs(words []string) [][]int {
	wordMap := make(map[string]int)
	for i, word := range words {
		wordMap[word] = i
	}

	var ans [][]int
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[:j]) {
				if k, ok := wordMap[reverse(word[j:])]; ok && k != i {
					ans = append(ans, []int{k, i})
				}
			}

			if j < len(word) && isPalindrome(word[j:]) {
				if k, ok := wordMap[reverse(word[:j])]; ok && k != i {
					ans = append(ans, []int{i, k})
				}
			}
		}
	}

	return ans
}

func reverse(word string) string {
	rev := make([]byte, len(word))
	n := len(word)
	for i := range word {
		rev[n-i-1] = word[i]
	}
	return string(rev)
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
