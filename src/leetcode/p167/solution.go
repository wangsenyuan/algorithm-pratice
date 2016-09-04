package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 9, 56, 90}
	fmt.Println(twoSum(nums, 8))
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		left := target - numbers[i]
		for k := i + 1; k <= j; {
			mid := (k + j) / 2
			if numbers[mid] == left {
				return []int{i + 1, mid + 1}
			}

			if numbers[mid] > left {
				j = mid - 1
			} else {
				k = mid + 1
			}
		}

		i++
	}

	return []int{i + 1, j + 1}
}
