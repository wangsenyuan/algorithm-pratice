package p804

func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	res := make(map[string]bool)

	for _, word := range words {
		var tmp string
		for i := 0; i < len(word); i++ {
			x := int(word[i] - 'a')
			tmp += codes[x]
		}
		res[tmp] = true
	}

	return len(res)
}
