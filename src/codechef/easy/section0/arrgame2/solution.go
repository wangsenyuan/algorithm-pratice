package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		as, bs := make([]int, n), make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &as[j])
		}
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &bs[j])
		}
		res := solve(n, as, bs)
		fmt.Printf("%.6f\n", res)
	}
}

func solve(n int, as []int, bs []int) float64 {
	sort.Ints(bs)

	ans := int64(0)

	one := sort.Search(n, func(j int) bool {
		return bs[j] > 1
	})

	two := sort.Search(n, func(j int) bool {
		return bs[j] > 2
	})

	four := sort.Search(n, func(j int) bool {
		return bs[j] >= 4
	})

	five := sort.Search(n, func(j int) bool {
		return bs[j] >= 5
	})

	for i := 0; i < n; i++ {
		a := as[i]
		if a == 1 {
			//you wish...
			continue
		}
		if a == 2 {
			ans += int64(one + n - five)
			continue
		}

		if a == 3 {
			ans += int64(two + n - four)
			continue
		}

		j := sort.Search(n, func(j int) bool {
			return bs[j] >= a+1
		})

		ans += int64(one + n - j)
	}

	return float64(ans) / float64(n)
}
