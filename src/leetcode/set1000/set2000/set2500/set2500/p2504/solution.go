package p2504

func similarPairs(words []string) int {
	var res int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if flag(words[i]) == flag(words[j]) {
				res++
			}
		}
	}
	return res
}

func flag(word string) int {
	var res int
	for i := 0; i < len(word); i++ {
		x := int(word[i] - 'a')
		res |= 1 << x
	}
	return res
}
