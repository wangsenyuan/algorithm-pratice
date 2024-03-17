package p3082

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func sumOfPower(nums []int, k int) int {
	// k <= 100
	//n := len(nums)
	// n <= 100
	// nums[i] <= 10000,
	// 那些很大的数，也会有贡献
	// dp[i][x] 表示到i为止，sum = x的子序列的计数
	// 这个思路是对的，但是计算有问题
	// dp[x]就是到目前为止，子序列和=x的power的和

	dp := make([]int, k+2)
	dp[0] = 1
	for _, num := range nums {
		fp := make([]int, k+2)

		for j := 0; j < k+2; j++ {
			// 将i加入前面的序列/或者不加入
			fp[j] = add(fp[j], mul(dp[j], 2))
			if j+num <= k {
				fp[j+num] = add(fp[j+num], dp[j])
			}
		}
		copy(dp, fp)
	}

	return dp[k]
}
