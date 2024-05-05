package p3136

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	cnt := []int{0, 0}
	for i := 0; i < len(word); i++ {
		if !isNumber(word[i]) && !isLetter(word[i]) {
			return false
		}
		if isLetter(word[i]) {
			if isVowel(word[i]) {
				cnt[0]++
			} else {
				cnt[1]++
			}
		}
	}

	return cnt[0] > 0 && cnt[1] > 0
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'A' || x == 'o' || x == 'O' || x == 'e' || x == 'E' || x == 'i' || x == 'I' || x == 'u' || x == 'U'
}

func isLetter(x byte) bool {
	return x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z'
}

func isNumber(x byte) bool {
	return x >= '0' && x <= '9'
}
