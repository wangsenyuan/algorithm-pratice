package p3080

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	var sum int
	for _, cur := range apple {
		sum += cur
	}
	sort.Ints(capacity)
	var res int
	for i := len(capacity) - 1; i >= 0 && sum > 0; i-- {
		res++
		sum -= capacity[i]
	}

	return res
}
