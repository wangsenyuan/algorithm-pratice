package p2418

import "sort"

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	sort.Slice(arr, func(i, j int) bool {
		return heights[arr[i]] > heights[arr[j]]
	})

	res := make([]string, n)

	for i := 0; i < n; i++ {
		res[i] = names[arr[i]]
	}

	return res
}
