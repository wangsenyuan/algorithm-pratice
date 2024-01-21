package p3012

func minimumCost(nums []int) int {

	first := 1 << 30
	second := 1 << 30

	for i := 1; i < len(nums); i++ {
		if nums[i] < first {
			second = first
			first = nums[i]
		} else if nums[i] < second {
			second = nums[i]
		}
	}

	return nums[0] + first + second
}
