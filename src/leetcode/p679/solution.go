package main

import "fmt"

func judgePoint24(nums []int) bool {

	var process func(count int, flag int, res []float64) bool

	process = func(count int, flag int, res []float64) bool {
		if count == 1 {
			for i, one := range res {
				if flag&(1<<uint(i)) > 0 && one == 24.0 {
					return true
				}
			}
			return false
		}

		for i := 0; i < 4; i++ {
			if flag&(1<<uint(i)) == 0 {
				continue
			}

			a := res[i]

			for j := 0; j < 4; j++ {
				if i == j || flag&(1<<uint(j)) == 0 {
					continue
				}

				if i < j {
					res[i] += res[j]
					if process(count-1, flag^(1<<uint(j)), res) {
						return true
					}
					res[i] = a
					res[i] *= res[j]
					if process(count-1, flag^(1<<uint(j)), res) {
						return true
					}
					res[i] = a
				}

				res[i] -= res[j]
				if process(count-1, flag^(1<<uint(j)), res) {
					return true
				}
				res[i] = a

				if res[j] != 0 {
					res[i] /= res[j]
					if process(count-1, flag^(1<<uint(j)), res) {
						return true
					}
					res[i] = a
				}
			}
		}
		return false
	}

	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		res[i] = float64(nums[i])
	}

	return process(4, 1<<4-1, res)
}

func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
	fmt.Println(judgePoint24([]int{1, 2, 1, 2}))
	fmt.Println(judgePoint24([]int{1, 2, 3, 4}))
}
