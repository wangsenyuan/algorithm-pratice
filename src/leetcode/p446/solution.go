package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}

func numberOfArithmeticSlices(A []int) int {
	dp := make([]map[int]int, len(A))
	for i := range A {
		dp[i] = make(map[int]int)
	}
	ans := 0
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			dist := A[i] - A[j]
			s := dp[j][dist] + 1
			dp[i][dist] += s
			ans += s
		}
		ans -= i
	}

	return ans
}
