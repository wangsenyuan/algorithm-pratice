package p2589

func vowelStrings(words []string, left int, right int) int {
	var res int

	for i := left; i <= right; i++ {
		if checkVowel(words[i]) {
			res++
		}
	}

	return res
}

func checkVowel(a string) bool {
	return isVowel(a[0]) && isVowel(a[len(a)-1])
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'e' || x == 'i' || x == 'u'
}
