package p748

func shortestCompletingWord(licensePlate string, words []string) string {
	cnt := make(map[byte]int)
	for i := 0; i < len(licensePlate); i++ {
		if isLetter(licensePlate[i]) {
			cnt[lower(licensePlate[i])]++
		}
	}
	var best string
	for _, word := range words {
		if len(best) > 0 && len(word) > len(best) {
			continue
		}
		tmp := make(map[byte]int)
		for i := 0; i < len(word); i++ {
			if _, found := cnt[word[i]]; found {
				tmp[word[i]]++
			}
		}
		if len(tmp) < len(cnt) {
			continue
		}
		complete := true
		for b, c := range cnt {
			if tmp[b] < c {
				complete = false
				break
			}
		}
		if complete {
			if len(best) == 0 || len(best) > len(word) {
				best = word
			}
		}
	}
	return best
}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return byte(b - 'A' + 'a')
	}
	return b
}

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z'
}
