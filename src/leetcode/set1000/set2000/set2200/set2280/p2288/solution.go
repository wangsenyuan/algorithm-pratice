package p2288

func totalSteps(nums []int) int {
	// 5, 4, 3, 2, 3, 2, 1
	// 5, 3
	// 5
	n := len(nums)

	stack := make([]int, n)
	var p int
	F := make([]int, n)
	var res int
	for i := 0; i < n; i++ {
		var cur int
		for p > 0 && nums[stack[p-1]] <= nums[i] {
			cur = max(cur, F[stack[p-1]])
			p--
		}
		if p > 0 {
			F[i] = cur + 1
			res = max(res, F[i])
		}
		stack[p] = i
		p++
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
