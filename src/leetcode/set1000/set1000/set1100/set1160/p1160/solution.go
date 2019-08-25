package p1160

func countCharacters(words []string, chars string) int {
	cnt := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		cnt[chars[i]-'a']++
	}
	cnt2 := make([]int, 26)
	check := func(word string) bool {
		for i := 0; i < len(cnt2); i++ {
			cnt2[i] = 0
		}
		for i := 0; i < len(word); i++ {
			cnt2[word[i]-'a']++
		}
		for i := 0; i < len(cnt); i++ {
			if cnt2[i] > cnt[i] {
				return false
			}
		}
		return true
	}

	var res int
	for _, word := range words {
		if check(word) {
			res += len(word)
		}
	}

	return res
}
