package p1935

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make([]bool, 26)

	for i := 0; i < len(brokenLetters); i++ {
		broken[int(brokenLetters[i]-'a')] = true
	}

	check := func(word string) bool {
		for i := 0; i < len(word); i++ {
			if broken[int(word[i]-'a')] {
				return false
			}
		}
		return true
	}

	var cnt int
	words := strings.Split(text, " ")

	for _, word := range words {
		if check(word) {
			cnt++
		}
	}
	return cnt
}
