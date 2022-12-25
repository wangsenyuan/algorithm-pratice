package p2518

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, MOD-2)
}

func countPartitions(nums []int, k int) int {
	// k <= 1000

	dp := make([]int, k)
	dp[0] = 1

	var res int

	n := len(nums)
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		for x := k - 1; x >= 0; x-- {
			if x+nums[i] >= k {
				y := sum - int64(x+nums[i])
				if y >= int64(k) {
					res = modAdd(res, modMul(dp[x], pow(2, n-(i))))
				}
			}
		}

		for x := k - 1; x >= 0; x-- {
			if x-nums[i] >= 0 {
				dp[x] = modAdd(dp[x], dp[x-nums[i]])
			}
		}
	}

	return res
}
