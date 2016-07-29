package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &s[j])
		}
		r := play(s)
		fmt.Printf("%d\n", r)
	}
}

func play(s []int) int {
	sort.Ints(s)
	diff := math.MaxInt32

	for i := 1; i < len(s); i++ {
		xdiff := s[i] - s[i-1]
		if xdiff < 0 {
			xdiff = -xdiff
		}
		if xdiff < diff {
			diff = xdiff
		}
	}

	return diff
}
