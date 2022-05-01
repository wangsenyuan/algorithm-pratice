package p2256

import "math"

func minimumAverageDifference(nums []int) int {
	var sum int64

	for _, num := range nums {
		sum += int64(num)
	}

	n := len(nums)
	var best int64 = math.MaxInt64
	at := -1
	var pref int64
	for i := 0; i < n; i++ {
		pref += int64(nums[i])
		a := pref / int64(i+1)
		sum -= int64(nums[i])
		var b int64
		if n-i-1 > 0 {
			b = sum / int64(n-i-1)
		}
		c := abs(a - b)
		if c < best {
			best = c
			at = i
		}
	}
	return at
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
