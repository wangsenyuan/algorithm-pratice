package p2736

func findNonMinOrMax(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return -1
	}
	a := min(nums[0], nums[1], nums[2])
	b := max(nums[0], nums[1], nums[2])
	return nums[0] + nums[1] + nums[2] - a - b
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
