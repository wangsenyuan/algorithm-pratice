package p2800

func countCompleteSubarrays(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	expect := len(cnt)
	var ans int
	n := len(nums)
	cnt = make(map[int]int)
	for l, r := 0, 0; r < n; r++ {
		cnt[nums[r]]++
		if len(cnt) < expect {
			continue
		}
		// len(cnt)== expect
		for len(cnt) == expect {
			cnt[nums[l]]--
			if cnt[nums[l]] == 0 {
				delete(cnt, nums[l])
			}
			l++
		}
		ans += l
		l--
		cnt[nums[l]]++
	}

	return ans
}
