package p2239

func findClosestNumber(nums []int) int {
	res := nums[0]

	for _, num := range nums {
		if abs(num) < abs(res) || abs(num) == abs(res) && num > 0 {
			res = num
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
