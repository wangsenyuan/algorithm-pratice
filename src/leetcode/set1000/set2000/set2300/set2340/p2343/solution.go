package p2343

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	m := len(nums)
	n := len(nums[0])
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m)
	}

	cnt := make([]int, 10)

	for r := 0; r < m; r++ {
		x := int(nums[r][n-1] - '0')
		cnt[x]++
	}

	for i := 1; i < 10; i++ {
		cnt[i] += cnt[i-1]
	}
	for r := m - 1; r >= 0; r-- {
		x := int(nums[r][n-1] - '0')
		cnt[x]--
		dp[1][cnt[x]] = r
	}

	for l := 2; l <= n; l++ {
		for x := 0; x < 10; x++ {
			cnt[x] = 0
		}

		for r := 0; r < m; r++ {
			x := int(nums[r][n-l] - '0')
			cnt[x]++
		}

		for x := 1; x < 10; x++ {
			cnt[x] += cnt[x-1]
		}

		for i := m - 1; i >= 0; i-- {
			j := dp[l-1][i]
			x := int(nums[j][n-l] - '0')
			cnt[x]--
			dp[l][cnt[x]] = j
		}
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		k, trm := cur[0], cur[1]
		k--
		res[i] = dp[trm][k]
	}

	return res
}
