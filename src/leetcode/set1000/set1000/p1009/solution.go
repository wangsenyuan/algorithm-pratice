package p1009

const INF = -1

var dd = []int{-1, 0, 1, 0, -1}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	n := R * C
	que := make([]int, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[r0*C+c0] = 0
	var tail int
	que[tail] = r0*C + c0
	tail++

	var front int
	for front < tail {
		cur := que[front]
		front++

		x, y := cur/C, cur%C

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < R && v >= 0 && v < C && dist[u*C+v] == INF {
				dist[u*C+v] = dist[cur] + 1
				que[tail] = u*C + v
				tail++
			}
		}
	}

	res := make([][]int, n)

	for i := 0; i < tail; i++ {
		cur := que[i]
		x, y := cur/C, cur%C
		res[i] = []int{x, y}
	}

	return res
}
