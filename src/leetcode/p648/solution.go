package main

import (
	"strings"
	"fmt"
)

func replaceWords(dict []string, sentence string) string {
	dictMap := make(map[string]bool)

	for _, word := range dict {
		dictMap[word] = true
	}

	words := strings.Split(sentence, " ")

	var res string

	for _, word := range words {

		i := 1
		for i < len(word) {
			if dictMap[word[:i]] {
				break
			}
			i++
		}
		res += word[:i]
		res += " "
	}

	return res[:len(res)-1]
}

func main() {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"

	fmt.Println(replaceWords(dict, sentence))
}
