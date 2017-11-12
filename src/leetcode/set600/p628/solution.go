package main

import (
	"math"
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	n := len(nums)

	as, bs := make([]int, 0, n), make([]int, 0, n)

	for _, num := range nums {
		if num > 0 {
			as = append(as, num)
		} else if num < 0 {
			bs = append(bs, num)
		}
	}


	zeros := n - len(as) - len(bs)

	sort.Ints(as)
	sort.Ints(bs)

	var res = math.MinInt32
	A := len(as)

	if A >= 3 {
		tmp := as[A-3] * as[A-2] * as[A-1]
		res = tmp
	}

	B := len(bs)

	if A >= 1 && B >= 2 {
		tmp := as[A-1] * bs[0] * bs[1]
		if tmp > res {
			res = tmp
		}
	}

	if zeros > 0 {
		if res < 0 {
			res = 0
		}
	} else {
		if A == 0 {
			tmp := bs[B-3] * bs[B-2] * bs[B-1]
			if tmp > res {
				res = tmp
			}
		} else if A == 2 && B == 1 {
			tmp := as[0] * as[1] * bs[0]
			if tmp > res {
				res = tmp
			}
		}
	}

	return res
}

func main() {
	nums := []int{7, 3, 1, 0, 0, 6}
	fmt.Println(maximumProduct(nums))
}
