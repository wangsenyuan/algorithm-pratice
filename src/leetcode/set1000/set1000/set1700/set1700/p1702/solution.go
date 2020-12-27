package p1702

func minMoves(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	var best = 1 << 31
	que := make([]int, n)
	var end int

	var count = -1
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count++
			que[end] = i - count
			sum[end+1] = que[end] + sum[end]
			end++
		}
	}

	for i := 0; i+k <= end; i++ {
		mid := (i + i + k - 1) / 2
		q := que[mid]
		best = min(best, (2*(mid-i)-k+1)*q+(sum[i+k]-sum[mid+1])-(sum[mid]-sum[i]))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
