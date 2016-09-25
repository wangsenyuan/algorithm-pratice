package main

import "fmt"

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	fmt.Println(shortestDistance(words, "coding", "practice"))
	fmt.Println(shortestDistance(words, "makes", "coding"))

}

func shortestDistance(words []string, word1 string, word2 string) int {
	a, b := -1, -1
	c := len(words)
	for i, word := range words {
		if word == word1 {
			a = i
		} else if word == word2 {
			b = i
		}

		if a >= 0 && b >= 0 && abs(a-b) < c {
			c = abs(a - b)
		}
	}

	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
