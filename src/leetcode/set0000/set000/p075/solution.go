package main

import "fmt"

func main() {
	nums := []int{1, 1, 0}
	sortColors(nums)
	fmt.Printf("%v\n", nums)
}

func sortColors(nums []int) {
	n := len(nums)
	j := n - 1
	i := 0
	k := 0
	for i < j {
		if nums[i] == 0 {
			i++
			continue
		}

		for j > i && nums[j] == 2 {
			j--
		}

		if i == j {
			break
		}

		if nums[i] != 1 || nums[j] != 1 {
			nums[i], nums[j] = nums[j], nums[i]
			continue
		}

		if k <= i {
			k = i + 1
		}

		for k < j && nums[k] == 1 {
			k++
		}
		if k == j {
			break
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}
