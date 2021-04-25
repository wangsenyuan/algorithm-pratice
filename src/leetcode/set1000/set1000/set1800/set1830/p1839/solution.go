package p1839

func longestBeautifulSubstring(word string) int {
	vowls := "aeiou"
	pos := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}
	var start int
	prev := -1
	var best int
	for i := 0; i < len(word); i++ {
		if word[i] == 'a' {
			if prev < 0 || word[prev] != 'a' {
				start = i
			}
			prev = i
			continue
		}
		if p, found := pos[word[i]]; found {
			if prev < 0 || word[prev] != word[i] && word[prev] != vowls[p-1] {
				prev = -1
				continue
			}
			prev = i
			if p == 4 {
				best = max(best, i-start+1)
			}
		} else {
			prev = -1
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
