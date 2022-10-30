package p2456

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	cnt := make(map[string]int)

	for i := 0; i < len(creators); i++ {
		cnt[creators[i]] += views[i]
	}

	var most_view int

	for _, v := range cnt {
		most_view = max(most_view, v)
	}

	most_view_id := make(map[string]int)
	n := len(creators)
	for i := 0; i < n; i++ {
		if cnt[creators[i]] == most_view {
			if j, ok := most_view_id[creators[i]]; !ok {
				most_view_id[creators[i]] = i
			} else if views[i] > views[j] || views[i] == views[j] && ids[i] < ids[j] {
				most_view_id[creators[i]] = i
			}
		}
	}

	var res [][]string

	for k, i := range most_view_id {
		res = append(res, []string{k, ids[i]})
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
