package p843

import "sort"

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

func findSecretWord(wordlist []string, master Master) {

	var loop func(words []string)

	loop = func(words []string) {
		if len(words) == 0 {
			return
		}
		sort.Strings(words)
		first := words[0]
		ret := master.Guess(first)
		if ret == 6 {
			return
		}
		next := make([]string, 0, len(words))

		for i := 1; i < len(words); i++ {
			if similar(words[i], first) == ret {
				next = append(next, words[i])
			}
		}
		loop(next)
	}

	loop(wordlist)
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
