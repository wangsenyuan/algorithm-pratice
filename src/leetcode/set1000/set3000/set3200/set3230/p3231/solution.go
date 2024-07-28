package p3231

func numberOfSubstrings(s string) (ans int) {
	a := []int{}
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}

	n := len(s)
	tot1 := n - len(a)
	a = append(a, n) // 哨兵

	for left, b := range s {
		if b == '1' {
			ans += a[0] - left
		}
		for k, j := range a[:len(a)-1] {
			cnt0 := k + 1
			if cnt0*cnt0 > tot1 {
				break
			}
			cnt1 := j - left - k
			ans += max(a[k+1]-j-max(cnt0*cnt0-cnt1, 0), 0)
		}
		if b == '0' {
			a = a[1:]
		}
	}
	return
}
