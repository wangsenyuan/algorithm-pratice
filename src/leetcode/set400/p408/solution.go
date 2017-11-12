package main

import "fmt"

func main() {
	fmt.Println(validWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(validWordAbbreviation("apple", "a2e"))
	fmt.Println(validWordAbbreviation("hi", "2i"))
	fmt.Println(validWordAbbreviation("a", "01"))
}

func validWordAbbreviation(word string, abbr string) bool {
	x := 0
	for i := 0; i < len(abbr); i++ {
		if abbr[i] >= '0' && abbr[i] <= '9' {
			if x == 0 && abbr[i] == '0' {
				return false
			}
			x = x*10 + int(abbr[i]-'0')
			continue
		}
		if x >= len(word) {
			return false
		}
		word = word[x:]
		x = 0
		if abbr[i] != word[0] {
			return false
		}
		word = word[1:]
	}

	return x == len(word)
}
