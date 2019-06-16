package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := []string{"hello", "world"}
	//sentence := []string{"a", "bcd", "e"}
	//sentence := []string{"I", "had", "apple", "pie"}
	//sentence := []string{"f", "p", "a"}
	rows := 2
	cols := 8
	fmt.Println(wordsTyping1(sentence, rows, cols))
}

func wordsTyping(sentence []string, rows int, cols int) int {
	n := len(sentence)
	num := make([]int, n)
	for i := range num {
		cnt := 0
		for j, l := i, 0; l+len(sentence[j]) <= cols; j = (j + 1) % n {
			cnt++
			l += len(sentence[j]) + 1
		}
		num[i] = cnt
	}
	start, total := 0, 0
	for i := 0; i < rows; i++ {
		total += num[start]
		start = total % n
	}
	return total / n
}

func wordsTyping1(sentence []string, rows int, cols int) int {
	start := 0
	str := strings.Join(sentence, " ") + " "
	n := len(str)

	for i := 0; i < rows; i++ {
		start += cols
		if str[start%n] == ' ' {
			start++
			continue
		}
		for start > 0 && str[(start-1)%n] != ' ' {
			start--
		}
	}

	return start / n
}
