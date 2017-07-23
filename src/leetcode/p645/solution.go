package main

import "fmt"

func findErrorNums(nums []int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	a, b := -1, -1

	for i := 1; i <= len(nums); i++ {
		if count[i] == 0 {
			b = i
		} else if count[i] == 2 {
			a = i
		}
	}

	return []int{a, b}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
