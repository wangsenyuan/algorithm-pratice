package p2344

import "sort"

func minOperations(nums []int, numsDivide []int) int {
	g := numsDivide[0]

	for i := 1; i < len(numsDivide); i++ {
		g = gcd(g, numsDivide[i])
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if g%nums[i] == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}
