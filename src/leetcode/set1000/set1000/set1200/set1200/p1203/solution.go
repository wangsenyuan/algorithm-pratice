package p1203

import "sort"

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {

	vis := make([]int, n)

	out := make([]int, n)

	var dfs1 func(u int, t *int) bool

	dfs1 = func(u int, t *int) bool {
		if vis[u] == 1 {
			return false
		}

		vis[u]++
		for _, v := range beforeItems[u] {
			if vis[v] < 2 && !dfs1(v, t) {
				return false
			}
		}
		vis[u]++
		out[u] = *t
		*t++
		return true
	}
	var t int
	for i := 0; i < n; i++ {
		if vis[i] == 0 && !dfs1(i, &t) {
			//has loop
			return nil
		}
	}

	nums := make([]Pair, n)

	for i := 0; i < n; i++ {
		nums[i] = Pair{out[i], i}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].first < nums[j].first
	})
	for i := 0; i < n; i++ {
		if group[i] < 0 {
			group[i] = m
			m++
		}
	}

	cnt := make([]int, m)

	for i := 0; i < n; i++ {
		cnt[group[i]]++
	}
	gg := make([][]int, m)

	for i := 0; i < m; i++ {
		gg[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		j := nums[i].second
		gg[group[j]] = append(gg[group[j]], j)
	}

	beforeGroups := make([]map[int]bool, m)

	for i := 0; i < m; i++ {
		beforeGroups[i] = make(map[int]bool)
	}

	for i := 0; i < n; i++ {
		x := group[i]
		for _, j := range beforeItems[i] {
			y := group[j]
			if x == y {
				continue
			}
			beforeGroups[x][y] = true
		}
	}
	res := make([]int, 0, n)
	var dfs2 func(u int) bool

	vis = make([]int, m)

	dfs2 = func(u int) bool {
		if vis[u] == 1 {
			return false
		}
		vis[u]++

		for v := range beforeGroups[u] {
			if vis[v] < 2 && !dfs2(v) {
				return false
			}
		}
		vis[u]++
		for _, i := range gg[u] {
			res = append(res, i)
		}
		return true
	}

	for i := 0; i < m; i++ {
		if vis[i] == 0 && !dfs2(i) {
			return nil
		}
	}

	return res
}

type Pair struct {
	first, second int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
