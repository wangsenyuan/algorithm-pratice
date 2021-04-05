package b

import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0] || logs[i][0] == logs[j][0] && logs[i][1] < logs[j][1]
	})

	cnt := make(map[int]map[int]int)

	for i := 0; i < len(logs); i++ {
		log := logs[i]
		id, t := log[0], log[1]
		if _, found := cnt[id]; !found {
			cnt[id] = make(map[int]int)
		}
		cnt[id][t]++
	}

	res := make([]int, k)

	for _, v := range cnt {
		if len(v) <= k {
			res[len(v)-1]++
		}
	}

	return res
}
