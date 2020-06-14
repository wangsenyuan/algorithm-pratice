package p1481

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	bloom := make([]bool, n)
	check := func(day int) bool {
		for i := 0; i < n; i++ {
			bloom[i] = day >= bloomDay[i]
		}

		var cnt int

		for i := 0; i < n && cnt < m; {

			if bloom[i] {
				j := i
				for j < n && bloom[j] {
					j++
				}

				cnt += (j - i) / k

				i = j
			} else {
				i++
			}
		}

		return cnt >= m
	}

	var left int = 1
	var right = 1000000001

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
