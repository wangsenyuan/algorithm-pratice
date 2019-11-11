package p1255

func maxScoreWords(words []string, letters []byte, score []int) int {
	cnt := count(letters)
	n := len(words)
	wc := make([][]int, n)

	for i := 0; i < n; i++ {
		wc[i] = count([]byte(words[i]))
	}

	N := 1 << uint(n)

	cc := make([]int, 26)

	var best int

	for state := 1; state < N; state++ {
		for i := 0; i < 26; i++ {
			cc[i] = 0
		}
		for i := 0; i < n; i++ {
			if (state & (1 << uint(i))) > 0 {
				for j := 0; j < 26; j++ {
					cc[j] += wc[i][j]
				}
			}
		}

		var got int

		for i := 0; i < 26; i++ {
			if cc[i] > cnt[i] {
				got = 0
				break
			}
			got += cc[i] * score[i]
		}

		if got > best {
			best = got
		}
	}

	return best
}

func count(letters []byte) []int {
	cnt := make([]int, 26)

	for _, letter := range letters {
		cnt[letter-'a']++
	}
	return cnt
}
