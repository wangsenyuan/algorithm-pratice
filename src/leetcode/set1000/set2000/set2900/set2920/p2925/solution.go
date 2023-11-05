package p2925

func findChampion(n int, edges [][]int) int {
	deg := make([]int, n)

	for _, edge := range edges {
		b := edge[1]
		deg[b]++
	}

	winer := -1
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			if winer >= 0 {
				return -1
			}
			winer = i
		}
	}
	return winer
}
