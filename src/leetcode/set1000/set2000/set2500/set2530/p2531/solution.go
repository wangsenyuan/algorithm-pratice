package p2531

func isItPossible(word1 string, word2 string) bool {
	cnt1 := count(word1)
	cnt2 := count(word2)

	check := func(i int, j int) bool {
		if i == j {
			return len(cnt1) == len(cnt2)
		}
		a := len(cnt1)
		if cnt1[i] == 1 {
			a--
		}
		b := len(cnt2)
		if cnt2[j] == 1 {
			b--
		}
		if cnt1[j] == 0 {
			a++
		}
		if cnt2[i] == 0 {
			b++
		}
		return a == b
	}

	for i := 0; i < 26; i++ {
		if cnt1[i] > 0 {
			for j := 0; j < 26; j++ {
				if cnt2[j] > 0 && check(i, j) {
					return true
				}
			}
		}

	}
	return false
}

func count(s string) map[int]int {
	res := make(map[int]int)
	for i := 0; i < len(s); i++ {
		res[int(s[i]-'a')]++
	}

	return res
}
