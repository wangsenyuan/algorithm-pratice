package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	var xs []int
	var res []int

	for i := 0; i < k; i++ {
		a := nums[i]
		for len(xs) > 0 && nums[xs[len(xs)-1]] <= a {
			xs = xs[:len(xs)-1]
		}
		xs = append(xs, i)
	}

	res = append(res, nums[xs[0]])

	for i := k; i < len(nums); i++ {
		for len(xs) > 0 && xs[0]+k <= i {
			xs = xs[1:]
		}

		a := nums[i]

		for len(xs) > 0 && nums[xs[len(xs)-1]] <= a {
			xs = xs[:len(xs)-1]
		}
		xs = append(xs, i)
		res = append(res, nums[xs[0]])
	}

	return res
}
