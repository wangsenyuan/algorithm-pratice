package p2195

import "sort"

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)

	var res int64

	n := len(nums)

	for i := 0; i < n && k > 0; i++ {
		cur := nums[i]
		var a, b int
		if i == 0 {
			a = 1
			b = cur - 1
		} else {
			a = nums[i-1] + 1
			b = cur - 1
		}
		cnt := b - a + 1
		if cnt <= 0 {
			continue
		}
		if cnt <= k {
			res += sum(a, b)
			k -= cnt
		} else {
			res += sum(a, a+k-1)
			k = 0
		}
	}

	if k > 0 {
		a := nums[n-1] + 1
		res += sum(a, a+k-1)
	}
	return res
}

func sum(a, b int) int64 {
	A := int64(a)
	B := int64(b)
	return (A + B) * (B - A + 1) / 2
}
