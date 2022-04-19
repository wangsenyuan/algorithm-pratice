package p2244

import "sort"

func minimumRounds(tasks []int) int {
	sort.Ints(tasks)

	var res int

	n := len(tasks)

	for i := 0; i < n; {
		j := i
		for i < n && tasks[i] == tasks[j] {
			i++
		}
		cnt := i - j
		// cnt = 2 * x + 3 * y
		x := cnt / 3
		r := cnt % 3
		if r == 0 {
			res += x
		} else if r == 2 {
			res += x + 1
		} else {
			// r == 1
			if x > 0 {
				res += x + 1
			} else {
				return -1
			}
		}
	}
	return res
}
