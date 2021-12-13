package p2017

func gridGame(grid [][]int) int64 {
	//如果第一个人在第i列的时候，向下，那么对于第二个人来说，只有两个策略;
	// 第一个策略时一开始就到第二行，并只获得第0列到i-1的分数;
	// 或者第一行i+1列到最后一列的值
	var sum0 int64

	for i := 0; i < len(grid[0]); i++ {
		sum0 += int64(grid[0][i])
	}

	var sum1 int64
	var best int64 = sum0

	for i := 0; i < len(grid[0]); i++ {
		// if turn at i
		sum0 -= int64(grid[0][i])

		best = min(best, max(sum1, sum0))
		sum1 += int64(grid[1][i])
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
