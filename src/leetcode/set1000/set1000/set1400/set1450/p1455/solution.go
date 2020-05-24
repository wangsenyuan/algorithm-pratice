package p1455

func isPrefixOfWord(sentence string, searchWord string) int {
	startWith := func(i int, n int) bool {
		var j int
		for i+j < n && j < len(searchWord) && sentence[i+j] == searchWord[j] {
			j++
		}
		return j == len(searchWord)
	}

	var j int
	var k int = 1
	for i := 1; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			if startWith(j, i) {
				return k
			}
			j = i + 1
			k++
		}
	}

	return -1
}
