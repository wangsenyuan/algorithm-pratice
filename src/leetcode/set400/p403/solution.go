package main

import "fmt"

func main() {
	//fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	//fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	fmt.Println(canCross([]int{0, 1, 3, 6, 7}))
}

func canCross(stones []int) bool {
	n := len(stones)
	if n <= 1 {
		return true
	}

	cache := make(map[int]map[int]bool)

	for i := 0; i < len(stones); i++ {
		cache[i] = make(map[int]bool)
	}

	var jump func(i int, k int) bool

	jump = func(i int, k int) bool {
		if i == n-1 {
			return true
		}

		if su, jumped := cache[i][k]; jumped {
			return su
		}

		su := false
		for j := i + 1; j < n && stones[j]-stones[i] <= k+1; j++ {
			if stones[j]-stones[i] < k-1 {
				continue
			}

			if stones[j]-stones[i] == k+1 && jump(j, k+1) {
				su = true
				break
			}

			if k > 0 && stones[j]-stones[i] == k && jump(j, k) {
				su = true
				break
			}

			if k > 1 && stones[j]-stones[i] == k-1 && jump(j, k-1) {
				su = true
				break
			}
		}

		cache[i][k] = su

		return su
	}

	return jump(0, 0)
}
