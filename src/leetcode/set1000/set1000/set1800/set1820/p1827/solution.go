package p1827

func minOperations(nums []int) int {
	var res int
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= prev {
			res += prev + 1 - nums[i]
			prev++
		} else {
			prev = nums[i]
		}
	}
	return res
}
