package p2076

import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	// more powerful worker, more effective getting pills
	// better to assign difficult task to powerful workers
	sort.Ints(tasks)
	// reverse(tasks)

	sort.Ints(workers)

	m := len(workers)

	n := len(tasks)

	que := make([]int, n)

	check := func(x int) bool {
		has := pills
		var front, end int
		var i int
		for j := m - x; j < m; j++ {
			for i < n && tasks[i] <= workers[j]+strength {
				que[end] = tasks[i]
				end++
				i++
			}
			if front == end {
				return false
			}
			if que[front] <= workers[j] {
				front++
				continue
			}
			if has == 0 {
				return false
			}
			has--
			end--
		}
		return true
	}

	l, r := 0, m+1

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}
