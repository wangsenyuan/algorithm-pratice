package p2656

func maximizeSum(nums []int, k int) int {
	var x = nums[0]
	for _, num := range nums {
		if num > x {
			x = num
		}
	}
	// num * k + (k-1) * k / 2
	return x*k + (k-1)*k/2
}
