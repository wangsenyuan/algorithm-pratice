package p1557

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	degree := make([]int, n)

	for _, e := range edges {
		degree[e[1]]++
	}

	res := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}
