package main

import "fmt"

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	var res int

	for num, x := range m {
		if k != 0 {
			if _, exists := m[num+k]; exists {
				res++
			}
		} else if x > 1 {
			res++
		}
	}

	return res
}
