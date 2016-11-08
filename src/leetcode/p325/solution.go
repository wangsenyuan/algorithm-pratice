package main

import "fmt"

func main() {
	fmt.Println(maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3))
	fmt.Println(maxSubArrayLen([]int{-2, -1, 2, 1}, 1))
}

func maxSubArrayLen(nums []int, k int) int {
	xs := make(map[int]int)
	ans := 0
	sum := 0
	xs[0] = -1
	for i, num := range nums {
		sum += num
		prev := sum - k
		if j, ok := xs[prev]; ok {
			if i-j > ans {
				ans = i - j
			}
		}

		if _, ok := xs[sum]; !ok {
			xs[sum] = i
		}
	}
	//fmt.Println(xs)
	return ans
}
