package main

import (
	"sort"
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	i := sort.Search(n, func(i int) bool {
		return arr[i] >= x
	})

	if i == n {
		return arr[n-k:]
	}

	if i == 0 {
		return arr[:k]
	}

	j := i - 1

	res := make([]int, 0, k)

	cnt := 0

	for cnt < k {
		if j < 0 || (i < n && arr[i]-x < x-arr[j]) {
			res = append(res, arr[i])
			i++
		} else if i == n || arr[i]-x >= x-arr[j] {
			res = append(res, arr[j])
			j--
		}

		cnt++
	}

	sort.Ints(res)

	return res
}

func main() {
	//fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	//fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{0, 1, 2, 2, 2, 3, 6, 8, 8, 9}, 5, 9))
}
