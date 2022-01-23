package p2148

func countElements(nums []int) int {
	a, b := nums[0], nums[0]

	for _, num := range nums {
		a = max(a, num)
		b = min(b, num)
	}

	var res int

	for _, num := range nums {
		if b < num && num < a {
			res++
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
