package p2016

func maximumDifference(nums []int) int {

	best := -1
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-cur > best && nums[i] > cur {
			best = nums[i] - cur
		}
		if nums[i] < cur {
			cur = nums[i]
		}
	}
	return best
}
