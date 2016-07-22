package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	bs := make([]int, 0, len(nums))
	//bs[0] = math.MaxInt32
	//cs := make([]int, 1, len(nums)+1)
	//cs[0] = math.MinInt32
	for _, a := range nums {
		bs = appendIfWiggle(bs, a)
	}

	return len(bs)
}

func appendIfWiggle(xs []int, x int) []int {
	if len(xs) == 0 {
		return append(xs, x)
	}

	if len(xs) == 1 && xs[0] == x {
		return xs
	}

	if len(xs) == 1 {
		return append(xs, x)
	}

	a, b := xs[len(xs)-2], xs[len(xs)-1]

	if (a > b && b < x) || (a < b && b > x) {
		return append(xs, x)
	}

	if (a > b && b > x) || (a < b && b < x) {
		return append(xs[:len(xs)-1], x)
	}

	return xs
}
