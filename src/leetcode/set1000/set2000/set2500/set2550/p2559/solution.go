package p2559

import "sort"

func maxCount(banned []int, n int, maxSum int) int {
	// better to choose small values=
	sort.Ints(banned)
	var sum int
	var res int
	for i, j := 1, 0; i <= n; {
		if j < len(banned) && banned[j] <= i {
			if banned[j] == i {
				i++
			}
			j++
			continue
		}

		if sum+i > maxSum {
			break
		}
		sum += i
		res++
		i++
	}
	return res
}
