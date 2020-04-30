package p1422

func maxScore(s string) int {
	var cnt1 int

	for i := len(s) - 1; i >= 0; i-- {
		cnt1 += int(s[i] - '0')
	}

	var res int

	var cnt0 int

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			cnt0++
		} else {
			cnt1--
		}
		if cnt0+cnt1 > res {
			res = cnt0 + cnt1
		}
	}

	return res
}
