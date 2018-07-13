package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var s string
		fmt.Fscanln(f, &s)
		fmt.Printf("Case #%d: %d\n", i, solve(s))
	}
}

func solve(s string) int64 {
	n := len(s)

	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, 210)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		sgn := -1
		if i == 0 {
			sgn = 1
		}
		for sgn <= 1 {
			var cur int
			for j := i; j < n; j++ {
				cur = (cur*10 + int(s[j]-'0')) % 210
				for x := 0; x < 210; x++ {
					y := (x + sgn*cur + 210) % 210
					dp[j+1][y] += dp[i][x]
				}
			}
			sgn += 2
		}
	}

	var ans int64

	for x := 0; x < 210; x++ {
		if x%2 == 0 || x%3 == 0 || x%5 == 0 || x%7 == 0 {
			ans += dp[n][x]
		}
	}

	return ans
}
