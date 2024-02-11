package p3034

import "slices"

func maxPalindromesAfterOperations(words []string) int {
	cnt := make([]int, 26)
	var odd int

	for _, word := range words {
		odd += len(word) & 1
		for i := 0; i < len(word); i++ {
			cnt[int(word[i]-'a')]++
		}
	}
	var odd2 int

	for i := 0; i < 26; i++ {
		odd2 += cnt[i] & 1
	}

	left := odd2 - odd
	res := len(words)

	slices.SortFunc(words, func(a, b string) int {
		return len(b) - len(a)
	})

	for _, w := range words {
		if left <= 0 {
			break
		}
		left -= len(w) / 2 * 2
		res--
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
