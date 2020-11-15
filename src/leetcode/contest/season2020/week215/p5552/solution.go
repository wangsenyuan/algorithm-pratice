package p5552

const INF = 100000

func minimumJumps(forbidden []int, a int, b int, x int) int {
	//sort.Ints(forbidden)
	dp := make(map[int]int)
	dp[0] = 0
	vis := make(map[int]bool)
	for i := 0; i < len(forbidden); i++ {
		vis[forbidden[i]] = true
	}
	var best = INF
	var dfs func(cur int, cnt int, flag int)

	dfs = func(cur int, cnt int, flag int) {
		if cur == 0 {
			best = min(best, cnt)
			return
		}
		if cur-a >= 0 && !vis[cur-a] {
			vis[cur-a] = true
			dfs(cur-a, cnt+1, 0)
		}
		if cur+b < INF && !vis[cur+b] && flag == 0 {
			vis[cur+b] = true
			dfs(cur+b, cnt+1, 1)
			vis[cur+b] = false
		}
	}
	dfs(x, 0, 0)
	if best == INF {
		return -1
	}
	return best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
