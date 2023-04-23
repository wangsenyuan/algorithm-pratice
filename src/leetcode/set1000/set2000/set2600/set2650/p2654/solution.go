package p2654

func minOperations(nums []int) int {
	n := len(nums)

	var res int

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			res++
		}
	}
	if res > 0 {
		// operate on others
		return n - res
	}
	// res = 0
	// find min operation to get 1
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				if res == 0 || res > j-i {
					res = j - i
				}
				break
			}
		}
	}

	if res == 0 {
		return -1
	}
	// use res operation to get 1
	return n - 1 + res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
