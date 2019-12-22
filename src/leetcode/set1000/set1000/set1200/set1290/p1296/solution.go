package p1296

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if cnt[num] > 0 {
			for j := 1; j < k; j++ {
				if cnt[num+j] == 0 {
					return false
				}
				cnt[num+j]--
			}
			cnt[num]--
		}
	}
	return true
}
