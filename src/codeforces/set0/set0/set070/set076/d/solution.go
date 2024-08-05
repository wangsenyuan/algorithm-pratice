package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var a, b uint

	fmt.Scan(&a, &b)

	res := solve(a, b)

	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(res[0], res[1])
	}
}

func solve(a uint, b uint) []uint {
	if a < b {
		return nil
	}
	if (a-b)%2 == 1 {
		return nil
	}
	x := (a - b) / 2
	y := x + b
	return []uint{x, y}
}

func solve1(a uint, b uint) []uint {
	ah := 63 - bits.LeadingZeros64(uint64(a))
	bh := 63 - bits.LeadingZeros64(uint64(b))
	if bh > ah {
		return nil
	}
	var x, y uint

	for i := ah; i >= 0; i-- {
		u := (a >> i) & 1
		v := (b >> i) & 1
		if u == 1 && v == 1 {
			y |= 1 << i
			continue
		}
		if u == 0 && v == 0 {
			continue
		}
		if u == 0 && v == 1 {
			return nil
		}
		// u == 1, v == 0
		j := i - 1
		for j >= 0 && (a>>j)&1 == 0 && (b>>j)&1 == 1 {
			y |= 1 << j
			j--
		}

		if j < 0 && (a>>j)&1 == 1 {
			return nil
		}
		// a >> j & 1 == 0 and b >> j & 1 == 0
		x |= 1 << j
		y |= 1 << j
		i = j + 1
	}

	if x+y != a || x^y != b {
		return nil
	}
	return []uint{x, y}
}
