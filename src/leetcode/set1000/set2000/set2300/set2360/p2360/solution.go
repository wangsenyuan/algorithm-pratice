package p2360

func longestCycle(edges []int) int {
	n := len(edges)
	vis := make([]int, n)

	cycleLength := func(x int) int {
		cnt := 1
		i := x
		for edges[i] != x {
			cnt++
			i = edges[i]
		}
		return cnt
	}

	travel := func(x int, label int) int {
		vis[x] = label
		for {
			if edges[x] < 0 {
				// not a loop
				return -1
			}
			if vis[edges[x]] > 0 {
				// a cycle
				if vis[edges[x]] == label {
					return cycleLength(edges[x])
				}
				return -1
			}
			vis[edges[x]] = label
			x = edges[x]
		}
	}
	best := -1
	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			tmp := travel(i, i+1)
			if tmp > 0 && best < tmp {
				best = tmp
			}
		}
	}

	return best
}
