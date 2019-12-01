package p1273

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	outs := make([]int, nodes)

	for i := 1; i < nodes; i++ {
		outs[parent[i]]++
	}
	que := make([]int, nodes)
	var end int

	for i := 0; i < nodes; i++ {
		if outs[i] == 0 {
			que[end] = i
			end++
		}
	}
	dp := make([]int, nodes)
	cnt := make([]int, nodes)
	for i := 0; i < nodes; i++ {
		dp[i] = value[i]
		cnt[i] = 1
	}

	res := nodes

	var front int

	for front < end {
		cur := que[front]
		front++

		if dp[cur] == 0 {
			res -= cnt[cur]
		}

		if parent[cur] >= 0 {
			if dp[cur] != 0 {
				dp[parent[cur]] += dp[cur]
				cnt[parent[cur]] += cnt[cur]
			}
			outs[parent[cur]]--

			if outs[parent[cur]] == 0 {
				que[end] = parent[cur]
				end++
			}
		}
	}

	return res
}
