package p2260

func minimumCardPickup(cards []int) int {
	pos := make(map[int]int)
	best := -1
	for i, card := range cards {
		if j, ok := pos[card]; ok {
			if best < 0 || best > i-j+1 {
				best = i - j + 1
			}
		}
		pos[card] = i
	}
	return best
}
