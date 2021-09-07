package p1991

func findMiddleIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	var left int
	for i := 0; i < len(nums); i++ {
		right := sum - left - nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}
