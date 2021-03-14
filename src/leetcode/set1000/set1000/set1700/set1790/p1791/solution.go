package p1791

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		deg[u-1]++
		deg[v-1]++
	}
	center := 0
	for i := 1; i < n; i++ {
		if deg[center] < deg[i] {
			center = i
		}
	}
	return center + 1
}
