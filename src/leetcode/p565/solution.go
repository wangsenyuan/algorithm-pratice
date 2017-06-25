package main

import "fmt"

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}

func arrayNesting(nums []int) int {
	checked := make([]bool, len(nums))

	ans := 0
	for i := 0; i < len(nums); i++ {
		if checked[i] {
			continue
		}

		tmp := 0
		j := i
		for !checked[j] {
			tmp++
			checked[j] = true
			j = nums[j]
		}
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}
