package main

import "fmt"

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	bs := []byte(s)
	return play(bs, k)
}

func play(s []byte, k int) int {
	snake := make(map[byte]int)
	j := 0
	res := 0

	meet := func(snake map[byte]int) bool {
		max := 0
		rep := 0
		for _, v := range snake {
			if v > max {
				max = v
			}
			rep += v
		}
		rep -= max
		return rep <= k
	}

	for i := 0; i <= len(s); i++ {
		if i < len(s) {
			snake[s[i]]++
		}
		if i-j > res {
			res = i - j
		}
		for !meet(snake) {
			snake[s[j]]--
			if snake[s[j]] == 0 {
				delete(snake, s[j])
			}
			j++
		}
	}

	return res
}
