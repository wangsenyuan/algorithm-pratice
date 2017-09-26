package main

import "fmt"

func judgePoint24(nums []int) bool {

	var process func(nums []float64) bool

	cal := func(nums []float64, i, j int, op func(a, b float64) float64) bool {
		xs := make([]float64, len(nums)-1)
		xs[0] = op(nums[i], nums[j])
		x := 1
		for k := 0; k < len(nums); k++ {
			if i == k || j == k {
				continue
			}
			xs[x] = nums[k]
			x++
		}
		return process(xs)
	}

	add := func(nums []float64, i, j int) bool {
		return cal(nums, i, j, func(a, b float64) float64 {
			return a + b
		})
	}

	mul := func(nums []float64, i, j int) bool {
		return cal(nums, i, j, func(a, b float64) float64 {
			return a * b
		})
	}

	sub := func(nums []float64, i, j int) bool {
		return cal(nums, i, j, func(a, b float64) float64 {
			return a - b
		})
	}

	div := func(nums []float64, i, j int) bool {
		if nums[j] == 0.0 {
			return false
		}
		return cal(nums, i, j, func(a, b float64) float64 {
			return a / b
		})
	}
	process = func(nums []float64) bool {
		if len(nums) == 0 {
			return false
		}

		if len(nums) == 1 {
			return nums[0] <= 24.0+1e-6 && nums[0] >= 24.0-1e-6
		}

		n := len(nums)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}

				if i > j {
					if add(nums, i, j) {
						return true
					}
					if mul(nums, i, j) {
						return true
					}
				}
				if sub(nums, i, j) {
					return true
				}

				if div(nums, i, j) {
					return true
				}
			}
		}
		return false
	}

	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		res[i] = float64(nums[i])
	}

	return process(res)
}

func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
	fmt.Println(judgePoint24([]int{1, 2, 1, 2}))
	fmt.Println(judgePoint24([]int{1, 2, 3, 4}))
	fmt.Println(judgePoint24([]int{3, 3, 8, 8}))

}
