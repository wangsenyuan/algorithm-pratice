package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
}

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}

	cnt := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			cnt++
		}
	}

	if cnt == 0 || cnt == len(word) {
		return true
	} else if cnt > 1 {
		return false
	} else {
		return word[0] >= 'A' && word[0] <= 'Z'
	}

}
