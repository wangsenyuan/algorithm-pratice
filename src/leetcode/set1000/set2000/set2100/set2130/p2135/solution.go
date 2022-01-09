package p2135

func wordCount(startWords []string, targetWords []string) int {
	mem := make(map[[26]int]bool)

	for _, word := range startWords {
		cnt := freq(word)
		mem[cnt] = true
	}
	var res int
	for _, word := range targetWords {
		cnt := freq(word)
		var ok bool
		for i := 0; i < 26 && !ok; i++ {
			if cnt[i] > 0 {
				cnt[i]--
				if mem[cnt] {
					ok = true
				}
				cnt[i]++
			}
		}
		if ok {
			res++
		}
	}
	return res
}

func freq(s string) [26]int {
	var res [26]int
	for i := 0; i < len(s); i++ {
		res[int(s[i]-'a')]++
	}
	return res
}
