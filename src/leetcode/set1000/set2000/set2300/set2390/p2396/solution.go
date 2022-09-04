package p2396

func findSubarrays(nums []int) bool {
	mem := make(map[int]bool)

	n := len(nums)
	for i := 0; i+1 < n; i++ {
		cur := nums[i] + nums[i+1]
		if mem[cur] {
			return true
		}
		mem[cur] = true
	}
	return false
}
