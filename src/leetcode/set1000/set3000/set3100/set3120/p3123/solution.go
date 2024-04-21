package p3123

func numberOfSpecialChars(word string) int {
	last := make([]int, 52)
	for i := 0; i < 52; i++ {
		last[i] = -1
	}
	for i := 0; i < len(word); i++ {
		var x int
		if word[i] >= 'a' && word[i] <= 'z' {
			x = int(word[i] - 'a')
		} else {
			x = 26 + int(word[i]-'A')
		}

		last[x] = i
	}
	var ans int
	for i := 0; i < 26; i++ {
		if last[i] >= 0 && last[i+26] >= 0 {
			ans++
		}
	}

	return ans
}
