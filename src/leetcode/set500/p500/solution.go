package main

import "fmt"

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

func findWords(words []string) []string {
	keyboard := make(map[byte]int)
	keyboard['q'] = 1
	keyboard['w'] = 1
	keyboard['e'] = 1
	keyboard['r'] = 1
	keyboard['t'] = 1
	keyboard['y'] = 1
	keyboard['u'] = 1
	keyboard['i'] = 1
	keyboard['o'] = 1
	keyboard['p'] = 1
	keyboard['a'] = 2
	keyboard['s'] = 2
	keyboard['d'] = 2
	keyboard['f'] = 2
	keyboard['g'] = 2
	keyboard['h'] = 2
	keyboard['j'] = 2
	keyboard['k'] = 2
	keyboard['l'] = 2
	keyboard['z'] = 3
	keyboard['x'] = 3
	keyboard['c'] = 3
	keyboard['v'] = 3
	keyboard['b'] = 3
	keyboard['n'] = 3
	keyboard['m'] = 3

	for k, r := range keyboard {
		keyboard[byte(k - 'a' + 'A')] = r
	}

	sameRow := func(word string) bool {
		row := keyboard[word[0]]
		for i := 1; i < len(word); i++ {
			if keyboard[word[i]] != row {
				return false
			}
		}
		return true
	}

	var res []string
	for _, word := range words {
		if sameRow(word) {
			res = append(res, word)
		}
	}
	return res
}
