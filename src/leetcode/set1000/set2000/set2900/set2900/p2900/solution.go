package p2900

import "strconv"

func lastVisitedIntegers(words []string) []int {
	var nums []int
	var ans []int

	var k int

	for _, word := range words {
		if word == "prev" {
			k++
			if len(nums)-k >= 0 {
				ans = append(ans, nums[len(nums)-k])
			} else {
				ans = append(ans, -1)
			}
		} else {
			num, _ := strconv.Atoi(word)
			nums = append(nums, num)
			k = 0
		}
	}

	return ans
}
