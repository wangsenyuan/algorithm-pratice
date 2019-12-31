package p1300

import "sort"

const INF = 1 << 30

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)

	n := len(arr)
	// |sum1 + m * value - target| = sum1 + m * value - target or target - (sum1 + m * value)
	var sum int
	dist := INF
	ans := -1

	for i := 0; i < n; i++ {
		// value >= arr[i-1]
		num := arr[i]
		// value <= num
		// x = |sum + (n - i) * value - target|

		if sum <= target {
			// x = sum - target + (n - i) * value
			// (n - i) * value - x = target - sum
			v := (target - sum) / (n - i)
			if v <= num && (i == 0 || v >= arr[i-1]) {
				x := target - sum - (n-i)*v
				if ans < 0 || abs(x) < dist {
					dist = abs(x)
					ans = v
				}
			}
			v++
			if v <= num && (i == 0 || v >= arr[i-1]) {
				x := target - sum - (n-i)*v
				if ans < 0 || abs(x) < dist {
					dist = abs(x)
					ans = v
				}
			}
		}
		sum += num
	}

	if abs(sum-target) < dist {
		ans = arr[n-1]
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
