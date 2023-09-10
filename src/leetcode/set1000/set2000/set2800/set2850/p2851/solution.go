package p2851

import "sort"

func numberOfPoints(nums [][]int) int {
	diff := make([]int, 102)
	for _, cur := range nums {
		diff[cur[0]]++
		diff[cur[1]+1]--
	}
	var ans int
	for i := 0; i <= 100; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		if diff[i] > 0 {
			ans++
		}
	}
	return ans
}

func numberOfPoints1(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0] || nums[i][0] == nums[j][0] && nums[i][1] > nums[j][1]
	})

	// 需要把被包含的部分给省略掉
	n := len(nums)
	//arr := make([][]int, n)
	var prev []int
	l, r := -10, -11
	var res int
	for i := 0; i < n; i++ {
		if len(prev) == 0 || nums[i][1] > prev[1] {
			// a new range added
			if r+1 < nums[i][0] {
				res += r - l + 1
				l = nums[i][0]
			}
			prev = nums[i]
			r = nums[i][1]
		}
	}

	res += r - l + 1

	return res
}
