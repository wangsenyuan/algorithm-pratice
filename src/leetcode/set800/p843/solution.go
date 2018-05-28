package p843

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 */
type Master struct {
	cnt    int
	secret string
	found  bool
}

func (this *Master) Guess(word string) int {
	this.cnt++
	if word == this.secret {
		this.found = true
	}
	return similar(word, this.secret)
}

func findSecretWord(wordlist []string, master *Master) {
	prop := make([][]int, 6)
	for i := 0; i < 6; i++ {
		prop[i] = make([]int, 26)
	}

	for _, word := range wordlist {
		for i := 0; i < 6; i++ {
			prop[i][int(word[i]-'a')]++
		}
	}

	best := func(words []string) string {
		var k int
		var max int
		for i, word := range words {
			var score int
			for j := 0; j < 6; j++ {
				score += prop[j][int(word[j]-'a')]
			}
			if score > max {
				max, k = score, i
			}
		}
		return words[k]
	}

	words := make([]string, len(wordlist))
	copy(words, wordlist)
	next := make([]string, len(wordlist))
	for len(words) > 0 {
		candidate := best(words)
		ret := master.Guess(candidate)
		if ret == 6 {
			break
		}
		var j int
		for _, word := range words {
			if similar(word, candidate) == ret {
				next[j] = word
				j++
				continue
			}
			for i := 0; i < 6; i++ {
				prop[i][int(word[i]-'a')]--
			}
		}
		words, next = next[:j], words
	}
}

func similar(a, b string) int {
	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			cnt++
		}
	}
	return cnt
}
