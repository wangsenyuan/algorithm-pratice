package p1831

func checkIfPangram(sentence string) bool {
	cnt := make([]int, 26)

	for i := 0; i < len(sentence); i++ {
		cnt[int(sentence[i]-'a')]++
	}

	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			return false
		}
	}
	return true
}
