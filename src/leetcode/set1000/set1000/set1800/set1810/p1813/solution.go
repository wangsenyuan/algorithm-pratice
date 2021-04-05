package p1813

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	var i int
	for i < len(sentence1) && i < len(sentence2) && sentence1[i] == sentence2[i] {
		i++
	}
	if i == len(sentence1) {
		return i == len(sentence2) || sentence2[i] == ' '
	}
	if i == len(sentence2) {
		return sentence1[i] == ' '
	}
	// s1[i] can't be blank space
	for i >= 0 && sentence1[i] != ' ' {
		i--
	}

	// i is the first position s1[i] != s2[i], either we continue i from s1 with some j in s2
	// or from s2 with some j in s1
	j := len(sentence1) - 1
	k := len(sentence2) - 1
	for j > i && k > i && sentence1[j] == sentence2[k] {
		j--
		k--
	}
	return j == i || k == i
}
