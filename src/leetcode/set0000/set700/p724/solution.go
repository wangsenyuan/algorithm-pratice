package p724

func pivotIndex(nums []int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		right := sum - nums[i] - left
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}
