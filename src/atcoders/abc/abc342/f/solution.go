package main

import (
	"fmt"
	"math"
)

func main() {
	var n, l, d int
	fmt.Scan(&n, &l, &d)
	res := solve(n, l, d)

	fmt.Printf("%.10f\n", res)
}

func solve(n int, l int, d int) float64 {
	dp := make([]float64, n+2)
	sum := 1.0
	dp[0] = 1
	fd := float64(d)
	for i := 1; i <= n; i++ {
		dp[i] = sum / fd
		if i < l {
			sum += dp[i]
		}
		if i-d < l && i-d >= 0 {
			sum -= dp[i-d]
		}
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if i < l {
			// 不能在这里获胜
			dp[i] = 0
		}
		dp[i] += dp[i-1]
	}
	sum = 0

	fp := make([]float64, n+1)
	for i := n; i >= 0; i-- {
		if i > 0 {
			fp[i] = math.Max(1.0-(dp[n]-dp[i-1]), sum/fd)
		} else {
			fp[i] = math.Max(1.0-dp[n], sum/fd)
		}
		sum += fp[i]
		if i+d <= n {
			sum -= fp[i+d]
		}
	}
	return fp[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
