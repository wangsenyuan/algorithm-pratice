package p2475

import "sort"

func unequalTriplets(nums []int) int {
	sort.Ints(nums)

	cnt1 := make(map[int]int)
	for _, num := range nums {
		cnt1[num]++
	}

	cnt2 := make(map[int]int)

	var sum = cnt1[nums[0]]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			cnt2[nums[i]] = sum * cnt1[nums[i]]
			sum += cnt1[nums[i]]
		}
	}

	var res int

	var sum2 int

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			// count only change
			y := cnt1[nums[i]]
			res += sum2 * y
			sum2 += cnt2[nums[i]]
		}
	}
	return res
}
