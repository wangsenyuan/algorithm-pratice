package p1524

func numSplits(s string) int {
	n := len(s)
	cnt := make([][]int, n)

	for i := 0; i < n; i++ {
		cnt[i] = make([]int, 26)

		if i > 0 {
			for j := 0; j < 26; j++ {
				cnt[i][j] = cnt[i-1][j]
			}
		}
		cnt[i][ind(s[i])]++
	}
	tot := make([]int, 26)
	for j := 0; j < 26; j++ {
		tot[j] = cnt[n-1][j]
	}
	var res int

	for i := n - 1; i >= 0; i-- {
		var left, right int
		for j := 0; j < 26; j++ {
			if cnt[i][j] > 0 {
				left++
			}
			if cnt[i][j] < tot[j] {
				right++
			}
		}
		if left == right {
			res++
		}
	}

	return res
}

func ind(x byte) int {
	return int(x - 'a')
}
