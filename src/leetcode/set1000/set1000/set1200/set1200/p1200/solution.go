package p1200

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	best := arr[1] - arr[0]

	i := 2

	for i < len(arr) {
		tmp := arr[i] - arr[i-1]
		if tmp < best {
			best = tmp
		}
		i++
	}

	var res [][]int

	i = 1

	for i < len(arr) {
		if arr[i]-arr[i-1] == best {
			res = append(res, []int{arr[i-1], arr[i]})
		}
		i++
	}

	return res
}
