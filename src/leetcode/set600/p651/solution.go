package main

import "fmt"

func maxA(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = i
		for j := 0; j <= i-3; j++ {
			if dp[j]*(i-j-1) > dp[i] {
				dp[i] = dp[j] * (i - j - 1)
			}
		}
	}

	return dp[n]
}

func main() {
	//fmt.Println(maxA(3))
	fmt.Println(maxA(7))
	fmt.Println(maxA(10))
	fmt.Println(maxA(11))

}
