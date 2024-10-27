package p3330

func possibleStringCount(word string) int {
	res := 1
	for i := 0; i < len(word); {
		j := i
		for i < len(word) && word[i] == word[j] {
			i++
		}
		if i-j > 1 {
			res += (i - j - 1)
		}
	}
	return res
}
