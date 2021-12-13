package p2063

func countVowels(word string) int64 {
	// any vowels
	var res int64
	n := len(word)

	vowels := map[byte]bool{'a': true, 'o': true, 'e': true, 'i': true, 'u': true}

	for i := 0; i < n; i++ {
		if vowels[word[i]] {
			res += int64(i+1) * int64(n-i)
		}
	}

	return res
}
