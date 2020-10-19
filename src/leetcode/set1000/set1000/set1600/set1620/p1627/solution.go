package p1627

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	n := len(ages)
	players := make([]Player, n)

	for i := 0; i < n; i++ {
		players[i] = Player{scores[i], ages[i], scores[i]}
	}

	sort.Sort(Players(players))

	var best int

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = players[i].score
		for j := 0; j < i; j++ {
			if players[j].score <= players[i].score {
				dp[i] = max(dp[i], dp[j]+players[i].score)
			}
		}
		best = max(best, dp[i])
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Player struct {
	score int
	age   int
	cnt   int
}

type Players []Player

func (ps Players) Len() int {
	return len(ps)
}

func (ps Players) Less(i, j int) bool {
	return ps[i].age < ps[j].age || ps[i].age == ps[j].age && ps[i].score < ps[j].score
}

func (ps Players) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
