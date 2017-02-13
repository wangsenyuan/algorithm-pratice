package main

import (
	"fmt"
)

func main() {
	//nums := []int{2, 4, 3, 5, 1}
	//nums := []int{1, 3, 2, 3, 1}
	nums := []int{-2, -2}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	var res int
	help := make([]int, len(nums))
	var sort func(left, right int)

	sort = func(left, right int) {
		if left+1 >= right {
			return
		}

		mid := left + (right-left)/2
		sort(left, mid)
		sort(mid, right)

		i, j, k, u := left, mid, left, left

		for i < mid && j < right {
			if nums[i] < nums[j] {
				help[k] = nums[i]
				i++
			} else {
				help[k] = nums[j]

				for u < mid && nums[u] <= 2*nums[j] {
					u++
				}
				res += mid - u

				j++

			}

			k++
		}

		for i < mid {
			help[k] = nums[i]
			k++
			i++
		}
		for j < right {
			help[k] = nums[j]
			k++

			for u < mid && nums[u] <= 2*nums[j] {
				u++
			}
			res += mid - u


			j++
		}

		copy(nums[left:right], help[left:right])
	}

	sort(0, len(nums))

	return res
}
