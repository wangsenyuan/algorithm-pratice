package p1522

func countOdds(low int, high int) int {
	if low == high {
		return high & 1
	}
	// 1, 2, 3
	// 2, 3
	// 3, 4
	// 4, 5, 6
	if low&1 == 1 && high&1 == 1 {
		return (high-low-1)/2 + 2
	}
	return (high - low + 1) / 2
}
