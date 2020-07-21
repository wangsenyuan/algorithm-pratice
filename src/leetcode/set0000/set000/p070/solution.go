package p070

func climbStairs(n int) int {
	a, b := 1, 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
