package p2216

func minDeletion(nums []int) int {
	n := len(nums)
	// n <= 1e5
	// a, a, a, a, a.... m个a
	// 要么删除m-1个a，要么删除前面的数字，且删除m-2个a
	// 如果连续的a处在偶数下标，至少需要删除m-1个数字
	var res int

	for i := 0; i < n; {
		if (res+i)%2 == 1 {
			// safe
			i++
			continue
		}
		j := i
		// even
		for i < n && nums[i] == nums[j] {
			i++
		}
		res += i - j - 1
	}
	if (n-res)%2 == 1 {
		res++
	}
	return res
}
