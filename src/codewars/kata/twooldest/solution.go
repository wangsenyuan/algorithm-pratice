package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {
	res := [...]int{-1, -1}

	for i := 0; i < len(ages); i++ {
		if res[1] == -1 || ages[res[1]] < ages[i] {
			res[1], res[0] = i, res[1]
		} else if res[0] == -1 || ages[res[0]] < ages[i] {
			res[0] = i
		}
	}

	res[0], res[1] = ages[res[0]], ages[res[1]]
	return res
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}
