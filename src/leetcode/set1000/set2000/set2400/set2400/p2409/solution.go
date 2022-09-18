package p2409

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	var res int

	for i, j := 0, 0; i < len(players) && j < len(trainers); i++ {
		for j < len(trainers) && trainers[j] < players[i] {
			j++
		}
		if j < len(trainers) {
			res++
			j++
		}
	}

	return res
}
