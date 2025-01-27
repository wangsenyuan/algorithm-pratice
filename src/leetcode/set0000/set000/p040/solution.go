package p040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	n := len(candidates)
	arr := make([]int, n)
	var dfs func(i int, j int, sum int)

	dfs = func(i int, j int, sum int) {
		if sum > target {
			return
		}
		if i == n {
			if sum == target {
				tmp := make([]int, j)
				copy(tmp, arr[:j])
				res = append(res, tmp)
			}
			return
		}

		ni := i
		for ni < n && candidates[ni] == candidates[i] {
			ni++
		}
		dfs(ni, j, sum)
		for ii := i; ii < ni; ii++ {
			arr[j] = candidates[i]
			j++
			sum += candidates[i]
			dfs(ni, j, sum)
		}
	}

	dfs(0, 0, 0)

	return res
}
