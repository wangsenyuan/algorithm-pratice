package p2564

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	sum := make([]int, n+1)

	for i, word := range words {
		sum[i+1] = sum[i] + checkVowel(word)
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		res[i] = sum[cur[1]+1] - sum[cur[0]]
	}
	return res
}

func checkVowel(s string) int {
	if isVowel(s[0]) && isVowel(s[len(s)-1]) {
		return 1
	}
	return 0
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'e' || x == 'i' || x == 'u'
}
