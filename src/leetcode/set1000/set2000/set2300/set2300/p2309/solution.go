package p2309

func minimumNumbers(num int, k int) int {

	dp := make([]int, num+1)

	for i := 0; i <= num; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	que := make([]int, num+1)
	var front, end int
	que[end] = 0
	end++
	for front < end {
		cur := que[front]
		front++
		for x := 0; cur+x*10+k <= num; x++ {
			next := cur + x*10 + k
			if dp[next] < 0 {
				dp[next] = dp[cur] + 1
				que[end] = next
				end++
			}
		}
	}

	return dp[num]
}
