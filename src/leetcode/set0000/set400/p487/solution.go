package main

import "fmt"

func main() {
	//nums := []int{1, 0, 1, 1, 0}
	nums := []int{1, 0, 1, 1, 0, 1}
	//nums := []int{1};
	//nums := []int{0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums, 0)
	j, k, best := 0, 0, 0

	for i, num := range nums {
		if num == 0 {
			if i-j > best {
				best = i - j
			}
			j, k = k, i + 1
		}
	}

	return best
}
