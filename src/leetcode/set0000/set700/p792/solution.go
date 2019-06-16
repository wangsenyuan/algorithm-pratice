package p792

func numMatchingSubseq(S string, words []string) int {
	xs := make([][]int, 26)
	for i := 0; i < 26; i++ {
		xs[i] = make([]int, 0, 100)
	}
	for i := 0; i < len(S); i++ {
		x := int(S[i] - 'a')
		xs[x] = append(xs[x], i)
	}

	p := make([]int, 26)

	check := func(word string) bool {
		for i := 0; i < 26; i++ {
			p[i] = 0
		}
		at := -1
		for i := 0; i < len(word); i++ {
			x := int(word[i] - 'a')

			j := p[x]
			for j < len(xs[x]) && xs[x][j] <= at {
				j++
			}
			if j == len(xs[x]) {
				return false
			}
			at = xs[x][j]
		}

		return true
	}

	var ans int
	for _, word := range words {
		if check(word) {
			ans++
		}
	}

	return ans
}
