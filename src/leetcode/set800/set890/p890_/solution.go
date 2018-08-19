package p890_

func findAndReplacePattern(words []string, pattern string) []string {

	check := func(word string) bool {
		if word == pattern {
			return true
		}

		ii := make([]int, 26)
		jj := make([]int, 26)
		for i := 0; i < 26; i++ {
			ii[i] = -1
			jj[i] = -1
		}
		n := len(word)
		for i := 0; i < n; i++ {
			a := int(word[i] - 'a')
			b := int(pattern[i] - 'a')

			if ii[a] < 0 && jj[b] < 0 {
				ii[a] = b
				jj[b] = a
				continue
			}
			if ii[a] < 0 || jj[b] < 0 || ii[a] != b || jj[b] != a {
				return false
			}
		}
		return true
	}

	res := make([]string, 0, len(words))
	for _, word := range words {
		if check(word) {
			res = append(res, word)
		}
	}
	return res
}
