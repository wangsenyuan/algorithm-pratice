package p1217

func minCostToMoveChips(chips []int) int {
	var a, b int

	for i := 0; i < len(chips); i++ {
		if chips[i]%2 == 0 {
			a++
		} else {
			b++
		}
	}

	if a < b {
		return a
	}
	return b
}
