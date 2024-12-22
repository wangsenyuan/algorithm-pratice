package p3392

func countSubarrays(nums []int) int {
	var res int

	for i := 1; i+1 < len(nums); i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			res++
		}
	}
	return res
}
