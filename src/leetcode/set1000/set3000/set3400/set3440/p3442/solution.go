package p3442

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)

	// 不需要保留相对顺序, pref[i] 表示i前面最长的free的空间
	pref := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			pref[i] = startTime[i]
		} else {
			pref[i] = max(pref[i-1], startTime[i]-endTime[i-1])
		}
	}
	var res int
	var suf int
	for i := n - 1; i >= 0; i-- {
		end := eventTime
		if i+1 < n {
			end = startTime[i+1]
		}
		open := 0
		if i > 0 {
			open = endTime[i-1]
		}
		need := endTime[i] - startTime[i]
		if suf >= need || (i > 0 && pref[i-1] >= need) {
			res = max(res, end-open)
		} else {
			res = max(res, end-open-need)
		}
		suf = max(suf, end-endTime[i])
	}

	return res
}
