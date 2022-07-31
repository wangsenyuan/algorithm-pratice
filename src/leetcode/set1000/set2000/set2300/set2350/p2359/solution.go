package p2359

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)

	d1 := make([]int, n)
	d2 := make([]int, n)
	for i := 0; i < n; i++ {
		d1[i] = -1
		d2[i] = -1
	}

	travel := func(x int, arr []int) {
		arr[x] = 0
		cur := x
		for edges[cur] >= 0 && arr[edges[cur]] < 0 {
			arr[edges[cur]] = arr[cur] + 1
			cur = edges[cur]
		}
	}

	travel(node1, d1)
	travel(node2, d2)

	best := -1
	dist := -1

	for i := 0; i < n; i++ {
		if d1[i] >= 0 && d2[i] >= 0 {
			if best < 0 || dist > max(d1[i], d2[i]) {
				best = i
				dist = max(d1[i], d2[i])
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
