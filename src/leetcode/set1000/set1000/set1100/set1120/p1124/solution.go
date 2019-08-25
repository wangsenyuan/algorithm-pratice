package p1124

func longestWPI(hours []int) int {
	n := len(hours)
	var sum int
	var ans int
	seen := make(map[int]int)

	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			sum++
		} else {
			sum--
		}
		if sum > 0 {
			ans = i + 1
		} else {
			if _, found := seen[sum]; !found {
				seen[sum] = i
			}
			if j, found := seen[sum-1]; found {
				ans = max(ans, i-j)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
