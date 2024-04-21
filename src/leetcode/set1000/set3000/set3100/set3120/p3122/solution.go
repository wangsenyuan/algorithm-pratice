package p3122

func numberOfSpecialChars(word string) int {
	first := make([]int, 52)
	for i := 0; i < 52; i++ {
		first[i] = -1
	}
	last := make([]int, 52)
	for i := 0; i < len(word); i++ {
		var x int
		if word[i] >= 'a' && word[i] <= 'z' {
			x = int(word[i] - 'a')
		} else {
			x = 26 + int(word[i]-'A')
		}

		if first[x] < 0 {
			first[x] = i
		}
		last[x] = i
	}
	var ans int
	for i := 0; i < 26; i++ {
		if first[i] < 0 || first[i+26] < 0 {
			continue
		}
		if last[i] < first[i+26] {
			ans++
		}
	}

	return ans
}
