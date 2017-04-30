package main

import "fmt"

func main() {
	nums := [][]int{
		{1, 2},
		{3, 4},
	}

	res := matrixReshape(nums, 1, 4)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%v\n", res[i])
	}
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	if n == 0 {
		return nums
	}

	m := len(nums[0])
	if m == 0 {
		return nums
	}

	if n*m != r*c {
		return nums
	}

	res := make([][]int, r)

	k := 0
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			x := k / m
			y := k % m
			res[i][j] = nums[x][y]
			k++
		}
	}

	return res
}
