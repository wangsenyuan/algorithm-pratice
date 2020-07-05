package main

import "sort"

func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	var x int

	if len(left) > 0 {
		x = left[len(left)-1]
	}

	if len(right) > 0 && n-right[0] > x {
		x = n - right[0]
	}

	return x
}
