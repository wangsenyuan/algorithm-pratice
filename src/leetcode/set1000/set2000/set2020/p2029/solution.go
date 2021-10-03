package p2029

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	n := len(s)
	cnt := make([][]int, n+1)
	cnt[n] = make([]int, 26)
	pos := make([][]int, n+1)
	pos[n] = make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[n][i] = -1
	}
	var lk int
	for i := n - 1; i >= 0; i-- {
		cnt[i] = make([]int, 26)
		copy(cnt[i], cnt[i+1])
		if s[i] == letter {
			lk++
		}
		j := int(s[i] - 'a')
		cnt[i][j] = lk

		pos[i] = make([]int, 26)
		copy(pos[i], pos[i+1])
		pos[i][j] = i
	}

	res := make([]byte, k)
	var p int
	r := repetition

	for i := 0; i < n && p < k; i++ {
		x := int(s[i] - 'a')
		// test whether we can put x at current position
		ok := true

		for y := 0; y < x; y++ {
			if pos[i][y] >= 0 && (r == 0 || cnt[i][y] >= r) && n-pos[i][y] >= k-p {
				ok = false
				break
			}
		}

		if ok && (r == 0 || (s[i] != letter && k-p > r || s[i] == letter && k-p >= r)) {
			res[p] = s[i]
			p++
			if s[i] == letter {
				r--
			}
		}
	}

	return string(res)
}
