package p1856

const MOD = 1000000007

func maxSumMinProduct(nums []int) int {
	n := len(nums)
	// 以nums[i] 为最小值，所能到达的最左边
	left := make([]int, n)
	stack := make([]int, n)
	// 13212
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		if p == 0 {
			left[i] = 0
		} else {
			left[i] = stack[p-1] + 1
		}
		stack[p] = i
		p++
	}
	p = 0
	right := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		if p == 0 {
			right[i] = n - 1
		} else {
			right[i] = stack[p-1] - 1
		}
		stack[p] = i
		p++
	}

	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(nums[i])
	}

	var best int64

	for i := 0; i < n; i++ {
		l, r := left[i], right[i]
		tmp := (sum[r+1] - sum[l]) * int64(nums[i])
		if best < tmp {
			best = tmp
		}
	}
	return int(best % MOD)
}
