package main

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))

}

func PredictTheWinner(nums []int) bool {
	var play1 func(i int, j int, a int64, b int64) bool
	var play2 func(i int, j int, a int64, b int64) bool

	play1 = func(i int, j int, a int64, b int64) bool {
		if i == j {
			return a+int64(nums[i]) >= b
		}

		if play2(i+1, j, a+int64(nums[i]), b) && play2(i, j-1, a+int64(nums[j]), b) {
			return false
		}
		return true
	}

	play2 = func(i int, j int, a int64, b int64) bool {
		if i == j {
			return b+int64(nums[i]) > a
		}
		if play1(i+1, j, a, b+int64(nums[i])) && play1(i, j-1, a, b+int64(nums[j])) {
			return false
		}
		return true
	}

	return play1(0, len(nums)-1, 0, 0)
}
