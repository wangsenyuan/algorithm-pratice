package p2425

import "sort"

func equalFrequency(word string) bool {
	freq := make([]int, 26)

	for i := 0; i < len(word); i++ {
		freq[int(word[i]-'a')]++
	}

	sort.Ints(freq)

	var p int
	for p < 26 && freq[p] == 0 {
		p++
	}

	freq = freq[p:]

	n := len(freq)

	if n == 1 {
		return true
	}

	// n > 1
	if n == 2 {
		if freq[0] == 1 {
			return true
		}
		if freq[1] == freq[0]+1 {
			return true
		}
		return false
	}

	// n > 2
	if freq[n-1] == freq[1] {
		return freq[0] == 1
	}

	if freq[n-2] == freq[0] {
		return freq[n-1] == freq[0]+1
	}
	return false
}
