package p3309

import (
	"math/bits"
	"slices"
)

const H = 8

func maxGoodNumber(nums []int) int {
	// 根据高位进行排列

	get := func(a, b int) int {
		l := bits.Len(uint(b))
		return (a << l) | b
	}

	slices.SortFunc(nums, func(a, b int) int {
		return get(b, a) - get(a, b)
	})

	var res int
	for _, num := range nums {
		l := bits.Len(uint(num))
		res <<= l
		res |= num
	}

	return res
}
