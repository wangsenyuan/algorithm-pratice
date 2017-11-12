package main

import "fmt"

func kInversePairs(n int, k int) int {
	dp := make([]int64, k+1)
	dp[0] = 1
	tmp := make([]int64, k+1)
	mod := int64(1000000007)
	for i := 1; i <= n; i++ {
		tmp[0] = 1
		for j := 1; j <= k; j++ {
			tmp[j] = (tmp[j-1] + dp[j]) % mod
			if j >= i {
				tmp[j] = tmp[j] + mod - dp[j-i]
			}
			tmp[j] %= mod

		}
		dp, tmp = tmp, dp
	}
	return int(dp[k])
}

func main() {
	fmt.Println(kInversePairs(3, 0))
	fmt.Println(kInversePairs(3, 1))
	fmt.Println(kInversePairs(3, 3))


}
