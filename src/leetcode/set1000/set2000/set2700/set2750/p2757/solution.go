package p2757

func longestAlternatingSubarray(nums []int, threshold int) int {
	var res int
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 || nums[i] > threshold {
			continue
		}
		cnt := 1
		for j := i + 1; j < n && nums[j] <= threshold && nums[j]&1 != nums[j-1]&1; j++ {
			cnt++
		}
		if res < cnt {
			res = cnt
		}
	}
	return res
}
