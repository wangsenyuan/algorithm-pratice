package p1996

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	roles := make([]Role, n)
	for i := 0; i < n; i++ {
		roles[i] = Role{properties[i][0], properties[i][1]}
	}
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].attack > roles[j].attack
	})

	var res int

	dp := make([]int, n)
	dp[0] = roles[0].defense
	before := make([]int, n)
	before[0] = -1

	for i := 1; i < n; i++ {
		if roles[i].attack == roles[i-1].attack {
			before[i] = before[i-1]
		} else {
			before[i] = i - 1
		}
		j := before[i]
		if j >= 0 && dp[j] > roles[i].defense {
			res++
		}
		dp[i] = max(dp[i-1], roles[i].defense)
	}

	return res
}

type Role struct {
	attack  int
	defense int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
