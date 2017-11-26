package p734

func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

	dict := make(map[string]map[string]bool)

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if dict[a] == nil {
			dict[a] = make(map[string]bool)
		}

		if dict[b] == nil {
			dict[b] = make(map[string]bool)
		}
		dict[a][b] = true
		dict[b][a] = true
	}

	for i := 0; i < len(words1); i++ {
		a, b := words1[i], words2[i]
		if a != b && !dict[a][b] {
			return false
		}
	}

	return true
}
