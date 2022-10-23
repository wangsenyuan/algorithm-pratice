package p2447

func subarrayGCD(nums []int, k int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		var g int
		for j := i; j >= 0; j-- {
			g = gcd(g, nums[j])
			if g == k {
				res++
			}
			if g < k {
				break
			}
		}
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
