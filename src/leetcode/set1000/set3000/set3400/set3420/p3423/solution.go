package p3423

func maxAdjacentDistance(nums []int) int {
	var res int
	n := len(nums)
	for i := 0; i < n; i++ {
		tmp := abs(nums[i] - nums[(i+1)%n])
		res = max(res, tmp)
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
