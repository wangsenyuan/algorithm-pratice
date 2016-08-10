package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &nums[j])
		}
		r := play(m, nums)
		if r {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func play(m int, nums []int) bool {
	n := len(nums)

	x := 1 << uint(n)

	for i := 1; i < x; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & (1 << uint(j))) > 0 {
				sum += nums[j]
			}
		}
		if sum == m {
			return true
		}
	}
	return false
}
