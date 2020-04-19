package p1419

func minNumberOfFrogs(croakOfFrogs string) int {
	n := len(croakOfFrogs)
	cnt := make([]int, 5)
	var res int
	for i := 0; i < n; i++ {
		x := croakOfFrogs[i]
		var y int
		if x == 'c' {
			y = 0
		} else if x == 'r' {
			y = 1
		} else if x == 'o' {
			y = 2
		} else if x == 'a' {
			y = 3
		} else {
			y = 4
		}
		if y == 0 {
			cnt[y]++
		} else {
			cnt[y]++
			cnt[y-1]--
		}

		if y == 4 {
			cnt[y]--
		}

		if cnt[y] < 0 || y > 0 && cnt[y-1] < 0 {
			return -1
		}

		var tmp int
		for j := 0; j < 5; j++ {
			tmp += cnt[j]
		}
		res = max(res, tmp)
	}

	for j := 0; j < 5; j++ {
		if cnt[j] > 0 {
			return -1
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
