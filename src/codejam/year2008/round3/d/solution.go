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

var cc [][]int64

func init() {
	cc = make([][]int64, MOD)
	cc[0] = make([]int64, MOD)
	cc[0][0] = 1
	for i := 1; i < MOD; i++ {
		cc[i] = make([]int64, MOD)
		cc[i][0] = 1
		cc[i][i] = 1
		for j := 1; j < i; j++ {
			cc[i][j] = (cc[i-1][j-1] + cc[i-1][j])
			if cc[i][j] >= MOD {
				cc[i][j] -= MOD
			}
		}
	}
}

func solve(H, W int64, R int, rocks [][]int64) int {

}

func calculate(r, c int64) int64 {
	if (r+c)%3 != 2 {
		return 0
	}
	rr := r - 1 - (r+c-2)/3
	cc := c - 1 - (r+c-2)/3

}

func C(n, k int64) int64 {
	res := int64(1)

	for n > 0 && k > 0 {
		x := n % MOD
		y := k % MOD
		res = (res * choose(int(x), int(y))) % MOD
	}
	return res
}

func choose(x, y int) int64 {
	if y > x {
		return 0
	}
	return cc[x][y]
}

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
