package p1297

const MOD = 1000000000007

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	cnt1 := make(map[int64]int)

	cnt2 := make(map[int]int)

	base := int64(1)
	var hash int64
	var i int
	for i < minSize {
		x := int(s[i] - 'a')
		hash = hash*37 + int64(x)
		hash %= MOD

		base *= 37
		base %= MOD

		cnt2[x]++
		i++
	}

	var best int

	cnt1[hash] = 1

	for i <= len(s) {
		if len(cnt2) <= maxLetters {
			best = max(best, cnt1[hash])
		}

		a := int(s[i-minSize] - 'a')

		cnt2[a]--

		if cnt2[a] == 0 {
			delete(cnt2, a)
		}

		if i < len(s) {
			b := int(s[i] - 'a')
			cnt2[b]++
			hash = hash*37 + int64(b)
			hash %= MOD

			hash = hash - (int64(a)*base)%MOD
			if hash < 0 {
				hash += MOD
			}

			cnt1[hash]++
		}
		i++
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
