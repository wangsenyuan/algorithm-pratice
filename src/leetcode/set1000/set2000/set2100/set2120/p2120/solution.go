package p2120

import "sort"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// n := len(recipes)
	store := make(map[string]bool)
	for _, sup := range supplies {
		store[sup] = true
	}

	vis := make(map[string]int)

	pos := make(map[string]int)

	for i, cur := range recipes {
		pos[cur] = i
	}

	good := make(map[string]bool)

	var dfs func(cur string, flag int) bool

	dfs = func(cur string, flag int) bool {
		if res, found := good[cur]; found {
			return res
		}
		if vis[cur] == flag {
			return false
		}
		vis[cur] = flag
		i := pos[cur]
		var cnt int
		for _, ing := range ingredients[i] {
			if store[ing] {
				cnt++
			} else if _, ok := pos[ing]; ok && dfs(ing, flag) {
				cnt++
			}
		}
		good[cur] = cnt == len(ingredients[i])
		return good[cur]
	}

	var res []string

	for i, rec := range recipes {
		if dfs(rec, i+1) {
			res = append(res, rec)
		}
	}
	sort.Strings(res)

	return res
}
