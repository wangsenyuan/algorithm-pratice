package p3430

func minMaxSubarraySum(nums []int, k int) int64 {
	res1 := calc(nums, k, func(a, b int) int {
		return a - b
	})

	res2 := calc(nums, k, func(a, b int) int {
		return b - a
	})

	return int64(res1 + res2)
}

func calc(nums []int, k int, cmp func(int, int) int) int {
	n := len(nums)
	stack := make([]int, n)
	var top int
	L := make([]int, n)
	for i := range n {
		for top > 0 && cmp(nums[stack[top-1]], nums[i]) <= 0 {
			top--
		}
		if top > 0 {
			L[i] = stack[top-1]
		} else {
			L[i] = -1
		}
		stack[top] = i
		top++
	}

	R := make([]int, n)
	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && cmp(nums[stack[top-1]], nums[i]) < 0 {
			top--
		}
		if top > 0 {
			R[i] = stack[top-1]
		} else {
			R[i] = n
		}
		stack[top] = i
		top++
	}

	var res int

	for i := 0; i < n; i++ {
		l := min(k, i-L[i])
		r := min(k, R[i]-i)
		t := k - r + 1
		if t > l {
			res += nums[i] * l * r
		} else {
			res += nums[i] * (t - 1) * r
			t = l - t + 1
			res += nums[i] * (r + r + 1 - t) * t / 2
		}
	}

	return res
}
