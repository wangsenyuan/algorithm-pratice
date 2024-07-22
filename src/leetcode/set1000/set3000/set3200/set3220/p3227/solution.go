package p3227

func doesAliceWin(s string) bool {
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			return true
		}
	}

	return false
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'u' || x == 'i' || x == 'e'
}
