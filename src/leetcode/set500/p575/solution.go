package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
}

func distributeCandies(candies []int) int {
	set := make(map[int]bool)

	for _, candy := range candies {
		if set[candy] {
			continue
		}
		set[candy] = true
	}

	m := len(set)
	n := len(candies)
	if 2*m <= n {
		return m
	}

	return n / 2
}
