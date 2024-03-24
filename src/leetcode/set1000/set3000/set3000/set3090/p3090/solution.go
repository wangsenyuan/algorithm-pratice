package p3090

func maximumLengthSubstring(s string) int {
	cnt := make([]int, 26)

	var best int

	for l, r := 0, 0; r < len(s); r++ {
		x := int(s[r] - 'a')
		cnt[x]++
		for cnt[x] > 2 {
			cnt[int(s[l]-'a')]--
			l++
		}
		best = max(best, r-l+1)
	}

	return best
}
