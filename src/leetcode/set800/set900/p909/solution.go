package p909

import "math"

func snakesAndLadders(board [][]int) int {

	n := len(board)

	bd := make([][]int, n)
	for i := 0; i < n; i++ {
		bd[i] = make([]int, n)
	}

	// mirror the board
	for i := 0; i < n; i++ {
		ii := n - 1 - i
		for j := 0; j < n; j++ {
			bd[ii][j] = board[i][j]
		}
	}

	getXY := func(idx int) (int, int) {
		idx--
		x := idx / n
		y := idx % n
		if x&1 == 1 {
			y = n - 1 - y
		}
		return x, y
	}

	que := make([]int, n*n*2)
	dist := make([]int, n*n+1)
	for i := 0; i <= n*n; i++ {
		dist[i] = math.MaxInt32
	}
	var front, tail int
	que[tail] = 1
	dist[1] = 0
	tail++
	for front < tail {
		cur := que[front]
		front++

		for k := 1; k <= 6 && cur+k <= n*n; k++ {
			next := cur + k
			u, v := getXY(next)
			if bd[u][v] == -1 && dist[next] > dist[cur]+1 {
				// just add it
				dist[next] = dist[cur] + 1
				que[tail] = next
				tail++
			} else if bd[u][v] >= 0 {
				// go to the cell
				dest := bd[u][v]
				if dist[dest] > dist[cur]+1 {
					dist[dest] = dist[cur] + 1
					que[tail] = dest
					tail++
				}
			}

		}
	}

	if dist[n*n] == math.MaxInt32 {
		return -1
	}
	return dist[n*n]
}
