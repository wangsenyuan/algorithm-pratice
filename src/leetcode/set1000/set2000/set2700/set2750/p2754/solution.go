package p2754

func sumImbalanceNumbers(nums []int) int {
	n := len(nums)
	cnt := make([]int, n+2)

	var sz int
	add := func(x int) {
		if cnt[x] > 0 {
			return
		}
		sz++
		cnt[x]++
		if cnt[x+1] > 0 {
			sz--
		}
		if cnt[x-1] > 0 {
			sz--
		}
	}

	var res int

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			add(nums[j])
			res += sz - 1
		}
		for j := 0; j <= i; j++ {
			cnt[nums[j]] = 0
		}
		sz = 0
	}

	return res
}
