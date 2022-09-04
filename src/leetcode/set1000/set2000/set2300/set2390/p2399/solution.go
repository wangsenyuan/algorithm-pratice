package p2399

func checkDistances(s string, distance []int) bool {
	pos := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
	}
	real_dist := make([]int, 26)

	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if pos[x] < 0 {
			pos[x] = i
		} else {
			real_dist[x] = i - pos[x] - 1
		}
	}

	for i := 0; i < 26; i++ {
		if pos[i] < 0 {
			continue
		}
		if distance[i] != real_dist[i] {
			return false
		}
	}

	return true
}
