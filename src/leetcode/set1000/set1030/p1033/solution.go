package p1033

import "sort"

func numMovesStones(a int, b int, c int) []int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	x := nums[0]
	y := nums[1]
	z := nums[2]
	var res1 int

	if x+1 == y {
		if y+1 == z {
			res1 = 0
		} else {
			res1 = 1
		}
	} else if y+1 == z {
		res1 = 1
	} else if x+2 == y {
		res1 = 1
	} else if y+2 == z {
		res1 = 1
	} else {
		res1 = 2
	}

	// res1 could be at max 2

	res2 := z - x - 1 - 1

	return []int{res1, res2}
}
