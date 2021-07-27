package p1944

func canSeePersonsCount(heights []int) []int {
	//min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1])
	// if height[j] >= height[i],
	// then j is the last person i could see
	n := len(heights)

	ans := make([]int, n)

	stack := make([]int, n)
	var p int
	for i := n - 1; i >= 0; i-- {
		for p > 0 && stack[p-1] < heights[i] {
			ans[i]++
			p--
		}
		if p > 0 {
			ans[i]++
			if stack[p-1] == heights[i] {
				p--
			}
		}
		// stack is in decreasing order, and all > heights[i]
		stack[p] = heights[i]
		p++
	}

	return ans
}
