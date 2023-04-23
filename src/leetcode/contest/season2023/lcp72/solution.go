package lcp72

func supplyWagon(supplies []int) []int {
	n := len(supplies)
	h := n / 2

	for n > h {
		var it int
		for i := 0; i+1 < n; i++ {
			if supplies[it]+supplies[it+1] > supplies[i]+supplies[i+1] {
				it = i
			}
		}
		supplies[it] += supplies[it+1]
		copy(supplies[it+1:], supplies[it+2:])
		n--
	}

	return supplies[:n]
}
