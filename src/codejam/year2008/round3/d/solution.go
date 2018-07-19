package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./D-small-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var H, W, R int
		fmt.Fscanf(f, "%d %d %d", &H, &W, &R)
		rocks := make([][]int, R)
		for j := 0; j < R; j++ {
			rocks[j] = make([]int, 2)
			fmt.Fscanf(f, "%d %d", &rocks[j][0], &rocks[j][1])
		}
		res := solveSmall(H, W, R, rocks)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

const MOD = 10007

func solveSmall(H, W, R int, rocks [][]int) int {
	dp := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		dp[i] = make([]int, W+1)
	}

	evil := make([][]bool, H+1)
	for i := 0; i <= H; i++ {
		evil[i] = make([]bool, W+1)
	}

	for _, rock := range rocks {
		i, j := rock[0], rock[1]
		evil[i][j] = true
	}

	dp[1][1] = 1

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if evil[i][j] {
				continue
			}
			if i-1 >= 1 && j-2 >= 1 {
				dp[i][j] += dp[i-1][j-2]
				if dp[i][j] >= MOD {
					dp[i][j] -= MOD
				}
			}
			if i-2 >= 1 && j-1 >= 1 {
				dp[i][j] += dp[i-2][j-1]
				if dp[i][j] >= MOD {
					dp[i][j] -= MOD
				}
			}
		}
	}

	return dp[H][W]
}
