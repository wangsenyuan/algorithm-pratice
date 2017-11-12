package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}

func numberOfArithmeticSlices(A []int) int {
	//DP[i][d] = the number of arithmetic subsequences whose last but one number is A[i], difference is d.
	dp := make([]map[int]int, len(A))
	nums := make(map[int]bool)
	for i, a := range A {
		dp[i] = make(map[int]int)
		nums[a] = true
	}
	ans := 0
	for i := 1; i < len(A); i++ {
		for j := i - 1; j >= 0; j-- {
			dist := A[i] - A[j]
			s := dp[j][dist]
			ans += s
			if nums[A[i]+dist] {
				dp[i][dist] += 1 + s
			}
		}
	}

	return ans
}
