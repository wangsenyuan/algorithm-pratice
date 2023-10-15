package p2899

import "sort"

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func countSubMultisets(nums []int, l int, r int) int {
	sort.Ints(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}

	if r > sum {
		r = sum
	}

	dp := make([]int, r+1)
	dp[0] = 1
	n := len(nums)
	for i := 0; i < n; {
		j := i
		for i < n && nums[i] == nums[j] {
			i++
		}
		v := i - j
		if nums[j] == 0 {
			dp[0] += v
			continue
		}
		// 无限制下的多重背包
		for x := nums[j]; x <= r; x++ {
			dp[x] = add(dp[x], dp[x-nums[j]])
		}
		// 无法选到 v+1的 nums[j]
		tmp := (v + 1) * nums[j]
		for x := r; x >= tmp; x-- {
			dp[x] = sub(dp[x], dp[x-tmp])
		}
	}

	var ans int
	for x := l; x <= r && x <= sum; x++ {
		ans = add(ans, dp[x])
	}

	return ans
}
func countSubMultisets1(nums []int, l int, r int) int {
	sort.Ints(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}

	dp := make([]int, sum+1)
	fp := make([]int, sum+1)
	dp[0] = 1

	n := len(nums)
	sum = 0
	for i := 0; i < n; {
		if nums[i] == 0 {
			dp[0]++
			i++
			continue
		}
		copy(fp, dp)
		j := i
		var cur int
		for i < n && nums[i] == nums[j] {
			sum += nums[i]
			cur += nums[i]
			for x := sum; x >= cur; x-- {
				dp[x] = add(dp[x], fp[x-cur])
			}
			i++
		}
	}

	var ans int
	for x := l; x <= r && x <= sum; x++ {
		ans = add(ans, dp[x])
	}

	return ans
}
