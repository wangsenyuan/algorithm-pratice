package main

import (
	"fmt"
	"sort"
)

func checkPossibility(nums []int) bool {
	n := len(nums)
	i := 0
	for i < n-1 {
		if nums[i] <= nums[i+1] {
			i++
			continue
		}

		if i == n-1 {
			return true
		}

		if !sort.IntsAreSorted(nums[i+1:]) {
			return false
		}
		a := nums[i+1]

		j := sort.Search(i+1, func(j int) bool {
			return nums[j] > a
		})

		if j == i {
			// we can just decrease nums[i] to a
			return true
		}
		xs := nums[i+1:]
		k := sort.Search(len(xs), func(k int) bool {
			return xs[k] > nums[i]
		})

		// then we can increase nums[i+1] to nums[i]
		return k == 1
	}

	return true
}

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{2, 3, 3, 2, 4}))

}
