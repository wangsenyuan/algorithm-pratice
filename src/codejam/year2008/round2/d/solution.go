package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	f, err := os.Open("./D-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var K int
		fmt.Fscanf(f, "%d", &K)
		var S string
		fmt.Fscanln(f, &S)
		res := solve(K, S)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

func solve(K int, S string) int {
	n := len(S)

	cal := func(x, y, t int) int {
		var ans int

		if x == t {
			//the last col
			for i := 0; i < len(S)-K; i += K {
				if S[i+x] != S[i+K+y] {
					ans++
				}
			}
		} else {
			for i := 0; i < n; i += K {
				if S[i+x] != S[i+y] {
					ans++
				}
			}
		}
		return ans
	}

	best := INF
	mat := make([][]int, K)
	for i := 0; i < K; i++ {
		mat[i] = make([]int, K)
	}

	for t := 0; t < K; t++ {
		// fix the last pos
		for x := 0; x < K; x++ {
			for y := 0; y < K; y++ {
				tmp := cal(x, y, t)
				mat[x][y] = tmp
			}
		}
		tmp := hamilton(K, mat)
		best = min(best, tmp)
	}

	return best + 1
}

const INF = math.MaxInt32

func hamilton(n int, mat [][]int) int {
	// dp[A][x] = The shortest path that starts from x, visits each point in A exactly once and ends up at node 0
	N := 1 << uint(n)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<uint(i)][i] = mat[i][0]
	}

	for mask := 0; mask < N; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) > 0 {
				for k := 0; k < n; k++ {
					if i != k && mask&(1<<uint(k)) > 0 {
						dp[mask][i] = min(dp[mask][i], dp[mask^(1<<uint(i))][k]+mat[i][k])
					}
				}
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
