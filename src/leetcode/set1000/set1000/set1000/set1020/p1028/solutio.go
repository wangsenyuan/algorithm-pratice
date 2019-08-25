package p1028

func divisorGame(N int) bool {
	dp := make([]int, N+1)

	var loop func(x int) bool

	loop = func(x int) bool {
		if x == 0 || x == 1 {
			return false
		}

		if dp[x] != 0 {
			return dp[x] == 1
		}

		for y := 1; y < x; y++ {
			if x%y == 0 && !loop(x-y) {
				dp[x] = 1
				return true
			}
		}
		dp[x] = -1
		return false
	}

	return loop(N)
}
