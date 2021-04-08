package p081

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(search([]int{1, 3, 1, 1, 1}, 3))
}

func search1(nums []int, target int) bool {
	p := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			p = i
			break
		}
	}

	index := -1
	if p < 0 {
		index = sort.SearchInts(nums, target)
	} else if nums[0] <= target {
		index = sort.SearchInts(nums[0:p+1], target)
	} else {
		index = p + 1 + sort.SearchInts(nums[p+1:], target)
	}

	if index < 0 || index >= len(nums) || nums[index] != target {
		return false
	}
	return true
}

func search(nums []int, target int) bool {
	if nums[0] == target {
		return true
	}
	n := len(nums)
	var p = n
	if nums[0] >= nums[n-1] {
		for p > 0 && nums[p-1] == nums[0] {
			p--
		}

		if p == 0 {
			return false
		}

		p = sort.Search(p, func(j int) bool {
			return nums[j] < nums[0]
		})

		if target < nums[0] {
			j := sort.Search(n-p, func(j int) bool {
				return nums[j+p] >= target
			})
			j += p
			return j < n && nums[j] == target
		}
	}
	j := sort.Search(p, func(j int) bool {
		return nums[j] >= target
	})

	return j < p && nums[j] == target
}
