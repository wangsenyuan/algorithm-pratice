package p2482

func bestClosingTime(customers string) int {
	n := len(customers)
	var cnt_y int

	for i := n - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			cnt_y++
		}
	}

	best := cnt_y
	var ans int
	var cnt_n int
	for i := 0; i < n; i++ {
		if customers[i] == 'Y' {
			cnt_y--
		} else {
			cnt_n++
		}
		if cnt_n+cnt_y < best {
			best = cnt_n + cnt_y
			ans = i + 1
		}
	}
	return ans
}
