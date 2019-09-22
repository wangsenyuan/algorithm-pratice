package p1203

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	beforeGroups := make([][]int, m+1)
	groupsItems := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		groupsItems[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		before := beforeItems[i]
		// make group to [0...m]
		x := group[i] + 1

		groupsItems[x] = append(groupsItems[x], i)

		if len(beforeGroups[x]) == 0 {
			beforeGroups[x] = make([]int, 0, 10)
		}

		set := make(map[int]bool)

		for _, j := range before {
			y := group[j] + 1
			if x != y && !set[y] {
				beforeGroups[x] = append(beforeGroups[x], y)
				set[y] = true
			}
		}
	}

	var visitGroup func(g int)
	var visitItem func(i int)
	vis := make([]bool, m+1)

	visitGroup = func(g int) {
		if vis[g] {
			return
		}
		vis[g] = true

		for _, p := range beforeGroups[g] {
			visitGroup(p)
		}

		items := groupsItems[g]

		for _, item := range items {
			visitItem(item)
		}
	}

	var res []int
	items := make([]bool, n)
	visitItem = func(i int) {
		if items[i] {
			return
		}
		items[i] = true

		before := beforeItems[i]

		for _, j := range before {
			if group[j] == group[i] {
				visitItem(j)
			}
		}
		res = append(res, i)
	}

	for i := 0; i <= m; i++ {
		if !vis[i] {
			visitGroup(i)
		}
	}

	set := make(map[int]bool)

	for i := 0; i < n; i++ {
		set[res[i]] = true

		for _, j := range beforeItems[res[i]] {
			if !set[j] {
				return []int{}
			}
		}
	}

	return res
}
