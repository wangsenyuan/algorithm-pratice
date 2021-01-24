package p1732

func largestAltitude(gain []int) int {
	var cur int
	var best int

	for i := 0; i < len(gain); i++ {
		cur += gain[i]
		if best < cur {
			best = cur
		}
	}

	return best
}
