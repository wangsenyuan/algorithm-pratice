package p318

func maxProduct(words []string) int {
	ws := make([]Word, len(words))
	for i, w := range words {
		ws[i] = NewWord(w)
	}
	var best int
	for i := 0; i < len(ws); i++ {
		for j := i + 1; j < len(ws); j++ {
			if noSameWord(ws[i], ws[j]) {
				best = max(best, ws[i].ll*ws[j].ll)
			}
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

type Word struct {
	cnt [26]int
	ll  int
}

func NewWord(s string) Word {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}
	return Word{cnt, len(s)}
}

func noSameWord(a, b Word) bool {
	for i := 0; i < 26; i++ {
		if a.cnt[i] > 0 && b.cnt[i] > 0 {
			return false
		}
	}
	return true
}
