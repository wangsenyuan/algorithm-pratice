package p2516

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	cnt := make([]int, 3)

	n := len(s)

	i := n

	for i > 0 {
		i--
		cnt[int(s[i]-'a')]++
		if cnt[0] >= k && cnt[1] >= k && cnt[2] >= k {
			break
		}
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1
	}
	best := n - i

	cnt2 := make([]int, 3)

	for j := 0; j < n; j++ {
		cnt2[int(s[j]-'a')]++
		for i < n {
			x := int(s[i] - 'a')
			cnt[x]--
			ok := true
			for ii := 0; ii < 3 && ok; ii++ {
				ok = cnt[ii]+cnt2[ii] >= k
			}
			if !ok {
				cnt[x]++
				break
			}
			best = min(best, j+1+(n-i-1))
			i++
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
