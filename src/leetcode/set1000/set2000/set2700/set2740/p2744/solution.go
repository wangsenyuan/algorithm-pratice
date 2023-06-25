package p2744

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	m := len(queries)

	cnt := make(map[int]int)

	add := func(svr int) {
		cnt[svr]++
	}

	rem := func(svr int) {
		cnt[svr]--
		if cnt[svr] == 0 {
			delete(cnt, svr)
		}
	}
	qs := make([]Query, m)
	for i := 0; i < m; i++ {
		qs[i] = Query{queries[i], i}
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].time < qs[j].time
	})

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	ans := make([]int, m)

	for l, r, j := 0, 0, 0; j < m; j++ {
		for r < len(logs) && logs[r][1] <= qs[j].time {
			add(logs[r][0])
			r++
		}
		for l < len(logs) && logs[l][1] < qs[j].time-x {
			rem(logs[l][0])
			l++
		}
		ans[qs[j].id] = n - len(cnt)
	}

	return ans
}

type Query struct {
	time int
	id   int
}
