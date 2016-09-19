package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{1, 2, 1, 2, 3, 3, 3}))
}

func majorityElement(nums []int) []int {
	a, b := 0, 1
	ac, bc := 0, 0

	for _, num := range nums {
		if num == a {
			ac++
		} else if num == b {
			bc++
		} else if ac == 0 {
			a = num
			ac++
		} else if bc == 0 {
			b = num
			bc++
		} else {
			ac--
			bc--
		}
	}

	ac, bc = 0, 0
	for _, num := range nums {
		if a == num {
			ac++
		} else if b == num {
			bc++
		}
	}

	var res []int
	if ac*3 > len(nums) {
		res = append(res, a)
	}

	if bc*3 > len(nums) {
		res = append(res, b)
	}

	return res
}
