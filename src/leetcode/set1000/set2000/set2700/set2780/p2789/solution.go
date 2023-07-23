package p2789

func maxArrayValue(nums []int) int64 {
	var res int64

	n := len(nums)
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		arr[i] = int64(nums[i])
	}

	for i := n - 1; i > 0; i-- {
		if arr[i-1] <= arr[i] {
			arr[i-1] += arr[i]
		}
		if res < arr[i] {
			res = arr[i]
		}
	}
	if res < arr[0] {
		res = arr[0]
	}
	return res
}
