package p3396

func minimumOperations(nums []int) int {
	n := len(nums)

	freq := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		freq[nums[i]]++
		if freq[nums[i]] > 1 {
			j := min((i+3)/3*3, n)
			return (j + 2) / 3
		}
	}
	return 0
}
