package p3033

func countMatchingSubarrays(nums []int, pattern []int) int {
	m := len(pattern)
	lps := make([]int, m)
	for i := 1; i < m; i++ {
		j := lps[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		lps[i] = j
	}

	var res int

	for i, j := 1, 0; i < len(nums); i++ {
		var x int
		if nums[i] > nums[i-1] {
			x = 1
		} else if nums[i] < nums[i-1] {
			x = -1
		}

		for j > 0 && x != pattern[j] {
			j = lps[j-1]
		}
		if x == pattern[j] {
			j++
		}
		if j == m {
			res++
			j = lps[j-1]
		}
	}

	return res
}
