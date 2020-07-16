package lcp12

func minTime(time []int, m int) int {

	check := func(t int) bool {
		var cnt int

		var i int

		for i < len(time) {
			var sum = time[i]
			var max = time[i]
			j := i + 1

			for j < len(time) && sum-max <= t {
				sum += time[j]
				if max < time[j] {
					max = time[j]
				}
				j++
			}

			// sum - max > t or j == len(time)

			cnt++
			if sum-max > t {
				j--
			}
			i = j
		}

		return cnt <= m
	}

	var left, right = 0, 1000000001

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
