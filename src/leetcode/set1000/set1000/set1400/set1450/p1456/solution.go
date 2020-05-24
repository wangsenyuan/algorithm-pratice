package p1456

func maxVowels(s string, k int) int {
	var best int

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			best++
		}
	}

	var sum = best

	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			sum++
		}
		if isVowel(s[i-k]) {
			sum--
		}
		if sum > best {
			best = sum
		}
	}
	return best
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'o' || c == 'e' || c == 'i' || c == 'u'
}
