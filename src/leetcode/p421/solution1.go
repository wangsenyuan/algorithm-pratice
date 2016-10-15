package main

import "fmt"

func main() {
	//fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{10, 23, 20, 18, 28}))
}

func findMaximumXOR(nums []int) int {
	mst := findMostSiginificantBit(nums)
	if mst < 0 {
		return 0
	}

	a, b := make([]int, 0, len(nums)), make([]int, 0, len(nums))

	for _, num := range nums {
		if num&(1<<uint(mst)) > 0 {
			a = append(a, num)
		} else {
			b = append(b, num)
		}
	}

	return xor(a, b, mst-1)
}

func xor(a, b []int, mst int) int {
	if len(a) == 1 && len(b) == 1 {
		return a[0] ^ b[0]
	}

	x11, x10, x01, x00 := make([]int, 0, len(a)), make([]int, 0, len(a)), make([]int, 0, len(b)), make([]int, 0, len(b))

	for _, num := range a {
		if num&(1<<uint(mst)) > 0 {
			x11 = append(x11, num)
		} else {
			x10 = append(x10, num)
		}
	}

	for _, num := range b {
		if num&(1<<uint(mst)) > 0 {
			x01 = append(x01, num)
		} else {
			x00 = append(x00, num)
		}
	}

	res := 0
	if len(x11) > 0 && len(x00) > 0 {
		res = xor(x11, x00, mst-1)
	}

	if len(x10) > 0 && len(x01) > 0 {
		y := xor(x10, x01, mst-1)
		if y > res {
			res = y
		}
	}
	return res
}

func findMostSiginificantBit(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	hb := -1
	for max != 0 {
		max >>= 1
		hb++
	}
	return hb
}
