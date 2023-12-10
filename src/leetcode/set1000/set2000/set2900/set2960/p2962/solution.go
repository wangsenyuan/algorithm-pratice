package p2962

func countSubarrays(nums []int, k int) int64 {
	x := nums[0]

	for _, cur := range nums {
		x = max(x, cur)
	}
	var arr []int
	for i, cur := range nums {
		if x == cur {
			arr = append(arr, i)
		}
	}
	if len(arr) < k {
		return 0
	}
	var res int64
	for i := k - 1; i < len(arr); i++ {
		j := i - k + 1
		before := arr[j] + 1
		// 最后一个最大值是是 arr[i]
		next := len(nums) - arr[i]
		if i+1 < len(arr) {
			next = arr[i+1] - arr[i]
		}
		res += int64(before) * int64(next)
	}

	return res
}
