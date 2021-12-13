package p2084

func countWords(words1 []string, words2 []string) int {
	cnt1 := count(words1)
	cnt2 := count(words2)
	var res int
	for k, v := range cnt1 {
		if v == 1 && cnt2[k] == 1 {
			res++
		}
	}
	return res
}

func count(ws []string) map[string]int {
	res := make(map[string]int)

	for _, w := range ws {
		res[w]++
	}

	return res
}
