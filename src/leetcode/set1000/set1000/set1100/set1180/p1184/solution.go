package p1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)

	var a int

	cur := start

	for cur != destination {
		a += distance[cur]
		cur = (cur + 1) % n
	}

	// cur == destination
	var b int
	for cur != start {
		b += distance[cur]
		cur = (cur + 1) % n
	}

	if a > b {
		return b
	}
	return a
}
