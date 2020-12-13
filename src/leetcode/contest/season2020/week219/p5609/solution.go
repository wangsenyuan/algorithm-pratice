package p5609

func countConsistentStrings(allowed string, words []string) int {
	cnt := count(allowed)

	var res int

	for _, word := range words {
		var ok = true
		for i := 0; i < len(word) && ok; i++ {
			ok = cnt[int(word[i]-'a')] > 0
		}
		if ok {
			res++
		}
	}

	return res
}

func count(str string) [26]int {
	var res [26]int
	for i := 0; i < len(str); i++ {
		res[int(str[i]-'a')]++
	}
	return res
}
