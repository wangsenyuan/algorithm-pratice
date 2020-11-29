package p5557

func maxRepeating(sequence string, word string) int {
	m := len(sequence)
	n := len(word)
	if m < n {
		return 0
	}
	check := func(cnt int) bool {
		for i := 0; i+cnt*n <= m; i++ {
			var j int
			for j < n {
				ok := true
				for k := 0; k < cnt && ok; k++ {
					if word[j] != sequence[i+j+k*n] {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
				j++
			}
			if j == n {
				return true
			}
		}
		return false
	}

	// m >= n
	k := m / n

	for i := k; i > 0; i-- {
		if check(i) {
			return i
		}
	}

	return 0
}
