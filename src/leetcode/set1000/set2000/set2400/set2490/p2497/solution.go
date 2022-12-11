package p2497

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	// 如果使用下标0似乎是一个更好的操作
	// 假设 a[i] = b[i], 且 a[0] != a[i], 使用下标0似乎是个比较好的方案

	cnt := make(map[int]int)

	for _, num := range nums2 {
		cnt[num]++
	}

	for _, v := range cnt {
		if v*2 > n {
			return -1
		}
	}

	// 所有相同的值，都需要被交换至少一次，应该想办法减少交换的次数
	// 约的数字越需要减少
	cnt = make(map[int]int)
	same := make(map[int]bool)

	for i := 0; i < n; i++ {
		if nums1[i] == nums2[i] {
			same[i] = true
			cnt[nums2[i]]++
		}
	}
	var item int
	var mx int
	for k, v := range cnt {
		if v > mx {
			item = k
			mx = v
		}
	}

	for i := 0; i < n && mx*2 > len(same); i++ {
		if !same[i] && nums1[i] != item {
			same[i] = true
		}
	}
	var res int64

	for k := range same {
		res += int64(k)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
