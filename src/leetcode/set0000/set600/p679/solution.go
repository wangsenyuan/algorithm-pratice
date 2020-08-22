package main

import "fmt"

func judgePoint24(nums []int) bool {
	// len(nums) == 4
	N := 1 << 4
	mem := make([]map[float64]int, N)
	mem[0] = make(map[float64]int)

	var dp func(bits int) map[float64]int

	dp = func(bits int) map[float64]int {
		if mem[bits] != nil {
			return mem[bits]
		}
		mem[bits] = make(map[float64]int)
		ii := make([]int, 0, 4)

		for i := 0; i < 4; i++ {
			if bits&(1<<i) > 0 {
				ii = append(ii, i)
			}
		}
		if len(ii) == 1 {
			mem[bits][float64(nums[ii[0]])]++
			return mem[bits]
		}
		xs := split(ii)

		for _, x := range xs {
			a, b := x[0], x[1]
			as := dp(a)
			bs := dp(b)

			for ak := range as {
				for bk := range bs {
					mem[bits][ak+bk]++
					mem[bits][ak-bk]++
					mem[bits][ak*bk]++
					if bk != 0 {
						mem[bits][ak/bk]++
					}
				}
			}
		}
		return mem[bits]
	}

	ans := dp(N - 1)

	for k := range ans {
		if k <= 24.0+1e-6 && k >= 24.0-1e-6 {
			return true
		}
	}

	return false
}

func split(arr []int) [][]int {
	var res [][]int
	n := len(arr)
	N := 1 << n
	for i := 1; i < n; i++ {
		//left has i elements, right has len(arr) - i elements
		flag := (1 << i) - 1

		for flag < N {
			var a, b int
			for j := 0; j < n; j++ {
				if flag>>j&1 == 1 {
					a |= 1 << arr[j]
				} else {
					b |= 1 << arr[j]
				}
			}
			res = append(res, []int{a, b})
			flag = snoob(flag)
		}
	}

	return res
}

// this function returns next higher number with same number of set bits as x.
func snoob(x int) int {
	rightOne := x & -x
	nextHigherOneBit := x + rightOne
	rightOnesPattern := x ^ nextHigherOneBit
	rightOnesPattern = (rightOnesPattern) / rightOne
	rightOnesPattern >>= 2
	next := nextHigherOneBit | rightOnesPattern
	return next
}

func judgePoint24_x(nums []int) bool {

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
