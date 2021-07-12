package p1917

func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)
	copy(res, nums)
	copy(res[len(nums):], nums)
	return res
}
