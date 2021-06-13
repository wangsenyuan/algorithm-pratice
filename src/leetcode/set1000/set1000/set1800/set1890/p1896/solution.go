package p1896

func makeEqual(words []string) bool {
	cnt := make([]int, 26)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			cnt[int(word[i]-'a')]++
		}
	}
	n := len(words)

	for i := 0; i < 26; i++ {
		if cnt[i]%n != 0 {
			return false
		}
	}
	return true
}
