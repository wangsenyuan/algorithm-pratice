package main

import "fmt"

func main() {

	var tc int

	fmt.Scanf("%d", &tc)

	for j := 0; j < tc; j++ {

		var n int
		fmt.Scanf("%d", &n)

		p := make([]float64, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%f", &p[i])
		}

		F := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &F[i])
		}
		var s float64
		fmt.Scanf("%f", &s)

		fmt.Printf("Case %d: %d\n", j, solve(n, p, F, s))
	}

}

func solve(N int, p []float64, F []int, s float64) int {
	S := int(s * 100)

	dp := make(map[int]int)
	dp[0] = 0
	//dp[0]

	for i := 0; i < N; i++ {
		x := int(p[i] * 100)
		y := 100 - x
		fp := make(map[int]int)
		for t := range dp {
			// win
			if t+y < S*(i+1) {
				fp[t+y] = max(fp[t+y], dp[t]+F[i])
			}
			//lose
			fp[t-x] = max(fp[t-x], dp[t])
		}
		dp = fp
	}

	var best int

	for t := range dp {
		if t < S*N {
			best = max(best, dp[t])
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
