package p1961

func isPrefixString(s string, words []string) bool {
	var i int
	for _, word := range words {
		for j := 0; j < len(word); j++ {
			if i == len(s) || word[j] != s[i] {
				return false
			}
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false
}
