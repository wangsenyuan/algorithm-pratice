package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sums := make([]int, n+1)

	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	ans := -10000.0

	for i := k; i <= n; i++ {
		tmp := float64(sums[i] - sums[i-k])
		if tmp/float64(k) > ans {
			ans = tmp / float64(k)
		}
	}

	return ans
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
