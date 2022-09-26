package p2420

const H = 20

func goodIndices(nums []int, k int) []int {
	n := len(nums)

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		if i > 0 && nums[i] <= nums[i-1] {
			dp[i] = 1 + dp[i-1]
		}
	}

	fp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		fp[i] = 1
		if i+1 < n && nums[i] <= nums[i+1] {
			fp[i] = 1 + fp[i+1]
		}
	}

	var res []int

	for i := k; i+1 < n; i++ {
		if dp[i-1] >= k && fp[i+1] >= k {
			res = append(res, i)
		}
	}
	return res
}

func goodIndices1(nums []int, k int) []int {
	// k <= i <= n - k
	// k <= n - k, k <= n / 2
	// i >= k
	// dp[i][j] = 以i结尾，长度为 2 ** j的sub arr 是否是非递增的
	n := len(nums)
	dp := make([][]bool, n)
	fp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, H)
		dp[i][0] = true
		fp[i] = make([]bool, H)
		fp[i][0] = true
	}

	for d := 1; d < H; d++ {
		for i := (1 << d) - 1; i < n; i++ {
			j := i - (1 << (d - 1))
			if dp[i][d-1] && dp[j][d-1] {
				if nums[j] >= nums[j+1] {
					dp[i][d] = true
				}
			}
		}

		for i := 0; i+(1<<d) <= n; i++ {
			j := i + (1 << (d - 1))
			if fp[i][d-1] && fp[j][d-1] {
				if nums[j-1] <= nums[j] {
					fp[i][d] = true
				}
			}
		}
	}

	check := func(i int) bool {
		kk := k
		a := i - 1
		b := i + 1
		for j := H - 1; j >= 0; j-- {
			if (k>>j)&1 == 1 {
				if !dp[a][j] || !fp[b][j] {
					return false
				}
				kk -= 1 << j
				a -= 1 << j
				if kk > 0 && nums[a] < nums[a+1] {
					return false
				}
				b += 1 << j
				if kk > 0 && nums[b-1] > nums[b] {
					return false
				}
			}
		}
		return true
	}

	var res []int

	for i := k; i+k <= n-1; i++ {
		if check(i) {
			res = append(res, i)
		}
	}

	return res
}
