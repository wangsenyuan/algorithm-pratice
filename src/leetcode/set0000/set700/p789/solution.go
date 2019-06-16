package p789

func escapeGhosts(ghosts [][]int, target []int) bool {
	dist := abs(target[0]) + abs(target[1])

	for _, ghost := range ghosts {
		tmp := abs(ghost[0]-target[0]) + abs(ghost[1]-target[1])
		if tmp <= dist {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
