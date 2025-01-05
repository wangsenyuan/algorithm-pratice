package p3411

import "slices"

func maxLength(nums []int) int {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}
	var ans int
	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
