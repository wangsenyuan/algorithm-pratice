package p1936

const INF = 1 << 30

func addRungs(rungs []int, dist int) int {
	// dp[i] = min of (dp[j] + (p[j] - p[i] -1)/ dist)
	// dp[i] = min (dp[j] * dist + p[j] - p[i] - 1)
	// for any i, it need to find min of dp[j] * dist + p[j]
	// lets uses a queue
	n := len(rungs)

	que := make([]int, n)
	var p int
	que[p] = rungs[n-1]
	p++

	for i := n - 2; i >= 0; i-- {
		// min is que[0]
		x := que[0]
		cur := (x - rungs[i] - 1) / dist
		y := cur*dist + rungs[i]
		for p > 0 && que[p-1] >= y {
			p--
		}
		que[p] = y
		p++
	}
	return (que[0] - 1) / dist
}
