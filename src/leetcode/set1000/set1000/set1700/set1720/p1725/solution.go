package p1725

import "sort"

func tupleSameProduct(nums []int) int {
	// a * b = c * d
	sort.Ints(nums)

	var res int
	cnt := make(map[Pair]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			a, b := nums[i], nums[j]
			g := gcd(a, b)
			a /= g
			b /= g
			p := Pair{a, b}
			res += cnt[p]
		}
		for j := i - 1; j >= 0; j-- {
			a, b := nums[j], nums[i]
			g := gcd(a, b)
			a /= g
			b /= g
			p := Pair{a, b}
			cnt[p]++
		}
	}

	return res * 8
}

type Pair struct {
	first, second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
