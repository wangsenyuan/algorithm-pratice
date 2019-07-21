package p5132

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	outs := make([][][]int, 2)
	outs[0] = make([][]int, n)
	outs[1] = make([][]int, n)

	for i := 0; i < n; i++ {
		for c := 0; c < 2; c++ {
			outs[c][i] = make([]int, 0, 10)
		}
	}

	for _, red := range red_edges {
		u, v := red[0], red[1]
		outs[0][u] = append(outs[0][u], v)
	}

	for _, blue := range blue_edges {
		u, v := blue[0], blue[1]
		outs[1][u] = append(outs[1][u], v)
	}

	res := make([]int, n)

	dist := make([][]int, 2)
	dist[0] = make([]int, n)
	dist[1] = make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
		dist[0][i] = -1
		dist[1][i] = -1
	}

	que := make([][]int, 0, 2*n)
	var front int
	que = append(que, []int{0, 0})
	que = append(que, []int{0, 1})
	res[0] = 0
	dist[0][0] = 0
	dist[1][0] = 0

	for front < len(que) {
		cur := que[front]
		front++
		u, color := cur[0], cur[1]
		d := dist[color][u]

		rcolor := 1 - color

		for _, v := range outs[rcolor][u] {
			if dist[rcolor][v] < 0 {
				dist[rcolor][v] = d + 1
				que = append(que, []int{v, rcolor})
			}
		}
	}

	for i := 0; i < n; i++ {
		if dist[0][i] < 0 && dist[1][i] < 0 {
			res[i] = -1
		} else if dist[0][i] < 0 {
			res[i] = dist[1][i]
		} else if dist[1][i] < 0 {
			res[i] = dist[0][i]
		} else {
			res[i] = min(dist[0][i], dist[1][i])
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
