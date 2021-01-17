package p1725

import "sort"

func tupleSameProduct(nums []int) int {
	// a * b = c * d
	sort.Ints(nums)

	var res int
	cnt := make(map[Pair]int)
	// a * d == b * c
	// a / b = c / d
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			b := nums[i]
			a := nums[j]
			g := gcd(a, b)
			a /= g
			b /= g
			p := Pair{a, b}
			if cnt[p] > 0 {
				res += cnt[p]
			}
			cnt[p]++
		}
	}
	cnt2 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cnt2[nums[i]*nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			x := nums[i] * nums[j]
			if cnt2[x] > 0 {
				res--
			}
		}
	}
	// a/b = c / d
	// a, d, b, c
	// d, a, c, b
	// b, c, a, d
	// c, b, d, a
	return res * 4
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
