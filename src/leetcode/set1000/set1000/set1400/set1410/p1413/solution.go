package p1413

func minStartValue(nums []int) int {
	var delta int

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			delta += -sum
			sum = 0
		}
	}
	return delta + 1
}
