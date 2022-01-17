package p2141

func maxRunTime(n int, batteries []int) int64 {
	// 假设可以工作x分钟

	var sum int64

	for _, num := range batteries {
		sum += int64(num)
	}

	N := int64(n)

	check := func(x int64) bool {
		var tot int64

		for i := 0; i < len(batteries); i++ {
			tot += min(x, int64(batteries[i]))
		}

		return tot >= N*x
	}

	right := (sum / N) + 1
	var left int64

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
