package main

import (
	"fmt"
)

func addBoldTag(s string, dict []string) string {
	bold := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		var l int
		for _, word := range dict {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				if len(word) > l {
					l = len(word)
				}
			}
		}

		for j := 0; j < l; j++ {
			bold[i+j] = true
		}
	}

	var res string
	for i := 0; i < len(bold); i++ {
		if !bold[i] {
			res += string(s[i])
			continue
		}
		j := i
		for ; j < len(bold) && bold[j]; j++ {

		}

		res += "<b>" + s[i:j] + "</b>"
		i = j - 1
	}

	return res
}

func main() {
	//s := "abcxyz123"
	//dict := []string{"abc", "123"}
	//s := "aaabbcc"
	//dict := []string{"aaa", "aab", "bc"}
	s := "aaabbcc"
	//dict := []string{"a", "b", "c"}
	dict := []string{"a", "c"}
	fmt.Println(addBoldTag(s, dict))
}
