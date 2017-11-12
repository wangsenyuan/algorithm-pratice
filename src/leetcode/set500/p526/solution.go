package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(countArrangement(i))
	}
}

func countArrangement(N int) int {
	nums := make([]int, N)

	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}

	var check func(n int) int

	check = func(n int) int {
		if n == 0 {
			return 1
		}
		res := 0

		for i := 0; i < n; i++ {
			if nums[i]%n == 0 || n%nums[i] == 0 {
				nums[i], nums[n - 1] = nums[n - 1], nums[i]
				res += check(n - 1)
				nums[i], nums[n - 1] = nums[n - 1], nums[i]
			}
		}
		return res
	}

	return check(N)
}
