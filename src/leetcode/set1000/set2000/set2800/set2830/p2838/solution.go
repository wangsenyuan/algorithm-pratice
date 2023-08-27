package p2838

func minimumPossibleSum(n int, target int) int64 {
	vis := make(map[int]bool)
	var sum int64
	cur := 1
	for i := 0; i < n; i++ {
		for vis[target-cur] {
			cur++
		}
		vis[cur] = true
		sum += int64(cur)
		cur++
	}

	return sum
}
