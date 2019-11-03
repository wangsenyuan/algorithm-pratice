package p1250

func isGoodArray(nums []int) bool {
	g := nums[0]

	for i := 1; i < len(nums); i++ {
		g = gcd(g, nums[i])
		if g == 1 {
			break
		}
	}

	return g == 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
