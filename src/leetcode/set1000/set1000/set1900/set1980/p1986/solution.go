package p1986

const INF = 1 << 28

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	N := 1 << n
	valid := make([]bool, N)
	for state := 0; state < N; state++ {
		var need int
		var cnt int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				need += tasks[i]
				cnt++
			}
		}
		if need <= sessionTime {
			valid[state] = true
		}
	}
	dp := make([]int, N)
	dp[0] = 0

	for cur := 1; cur < N; cur++ {
		dp[cur] = INF
		for sub := cur; sub > 0; sub = (sub - 1) & cur {
			if valid[sub] {
				dp[cur] = min(dp[cur], 1+dp[cur^sub])
			}
		}
	}
	return dp[N-1]
}

func minSessions1(tasks []int, sessionTime int) int {
	n := len(tasks)
	// n <= 14
	// sessionTime <= 15,
	// max(tasks[i]) <= sessionTime
	// dp[state][time]
	N := 1 << n
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, sessionTime+1)
		for j := 0; j <= sessionTime; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][sessionTime] = 1

	for cur := 0; cur < N; cur++ {
		for i := 0; i < n; i++ {
			if (cur>>i)&1 == 1 {
				continue
			}
			for time := tasks[i]; time <= sessionTime; time++ {
				if dp[cur][time] < INF {
					// still in the previous work session
					dp[cur|(1<<i)][time-tasks[i]] = min(dp[cur|(1<<i)][time-tasks[i]], dp[cur][time])
				}
			}
			// start a new session
			dp[cur|(1<<i)][sessionTime-tasks[i]] = min(dp[cur|(1<<i)][sessionTime-tasks[i]], dp[cur][0]+1)
			for time := sessionTime - 1; time >= 0; time-- {
				dp[cur|(1<<i)][time] = min(dp[cur|(1<<i)][time], dp[cur|(1<<i)][time+1])
			}
		}
	}
	return dp[N-1][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
