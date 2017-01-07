package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, k, x int
		fmt.Scanf("%d %d", &n, &k)

		dp := make([][]int, n + 1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, 1024)
		}

		dp[0][0] = 1

		for i := 1; i <= n; i++ {
			fmt.Scanf("%d", &x)
			for j := 0; j < 1024; j++ {
				dp[i][j] = dp[i - 1][j] | dp[i - 1][j ^ x]
			}
		}
		res := 0
		for j := 0; j < 1024; j++ {
			if dp[n][j]*(j^k) > res {
				res = j ^ k
			}
		}

		fmt.Println(res)

		t--
	}
}
