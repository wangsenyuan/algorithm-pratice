package p3452

func sumOfGoodNumbers(nums []int, k int) int {
	var res int
	n := len(nums)
	for i, x := range nums {
		if (i-k < 0 || x > nums[i-k]) && (i+k >= n || x > nums[i+k]) {
			res += x
		}
	}
	return res
}
