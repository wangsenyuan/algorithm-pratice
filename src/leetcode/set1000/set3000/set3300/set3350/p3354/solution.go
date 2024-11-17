package p3354

func countValidSelections(nums []int) int {
	// 如果在一个位置i处开始运动，如果它左边的数是x，右边的数是y
	// 那么如果 x < y, 那么必然没法让左边的数y减小到0
	// got
	n := len(nums)

	var pref int
	for _, num := range nums {
		pref += num
	}

	var suf int
	var res int
	for i := n - 1; i >= 0; i-- {
		pref -= nums[i]
		if nums[i] == 0 {
			if pref == suf {
				res += 2
			} else if pref == suf+1 {
				res++
			} else if pref+1 == suf {
				res++
			}
		}
		suf += nums[i]
	}

	return res
}
