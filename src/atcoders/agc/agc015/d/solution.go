package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var low, high uint
	fmt.Scan(&low, &high)
	res := solve(low, high)
	fmt.Println(res)
}

func solve(low uint, high uint) uint {
	if low == high {
		return 1
	}
	ans := high - low + 1
	mask := uint(1)<<(bits.Len(high^low)-1) - 1
	high &= mask
	low &= mask
	nh := bits.Len(high)
	if bits.Len(low) <= nh {
		ans += mask - high
	} else {
		ans += mask - low + 1<<nh - high
	}

	return ans
}
