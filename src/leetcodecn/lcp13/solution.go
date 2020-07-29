package lcp13

var dd = []int{-1, 0, 1, 0, -1}

const INF = 1 << 30

func minimalSteps(maze []string) int {
	// m <= 16
	start := findPosition(maze, 'S')[0]
	end := findPosition(maze, 'T')[0]
	M := findPosition(maze, 'M')
	O := findPosition(maze, 'O')

	rows := len(maze)
	cols := len(maze[0])
	N := rows * cols
	que := make([]Pair, N)
	dist := make([]int, N)
	bfs := func(p1, p2 Pair) int {
		for i := 0; i < N; i++ {
			dist[i] = -1
		}
		var front, end int
		dist[p1.Pos(cols)] = 0
		que[end] = p1
		end++
		for front < end {
			cur := que[front]
			if cur == p2 {
				break
			}
			front++
			for i := 0; i < 4; i++ {
				u, v := cur.first+dd[i], cur.second+dd[i+1]
				if u < 0 || u >= rows || v < 0 || v >= cols || maze[u][v] == '#' || dist[u*cols+v] >= 0 {
					continue
				}
				dist[u*cols+v] = dist[cur.Pos(cols)] + 1
				que[end] = Pair{u, v}
				end++
			}
		}
		return dist[p2.Pos(cols)]
	}

	// len(M) <= 16
	m := len(M)
	if m == 0 {
		return bfs(start, end)
	}

	startRockDist := make([]int, len(O))

	for i := 0; i < len(O); i++ {
		startRockDist[i] = bfs(start, O[i])
	}

	startDist := make([]int, m)
	endDist := make([]int, m)

	rockDist := make([][]int, m)
	for i := 0; i < m; i++ {
		rockDist[i] = make([]int, len(O))
		startDist[i] = -1
		for j := 0; j < len(O); j++ {
			rockDist[i][j] = bfs(M[i], O[j])
			if startRockDist[j] > 0 && rockDist[i][j] > 0 {
				if startDist[i] < 0 || startRockDist[j]+rockDist[i][j] < startDist[i] {
					startDist[i] = startRockDist[j] + rockDist[i][j]
				}
			}
		}
		endDist[i] = bfs(end, M[i])
		if startDist[i] < 0 || endDist[i] < 0 {
			return -1
		}
	}
	// is the distance bridge by a pile of rocks between too Ms
	pairDist := make([][]int, m)
	for i := 0; i < m; i++ {
		pairDist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == j {
				continue
			}
			var best = -1
			for k := 0; k < len(O); k++ {
				if rockDist[i][k] < 0 || rockDist[j][k] < 0 {
					continue
				}
				if best < 0 || rockDist[i][k]+rockDist[j][k] < best {
					best = rockDist[i][k] + rockDist[j][k]
				}
			}
			if best < 0 {
				return -1
			}
			pairDist[i][j] = best
		}
	}

	X := 1 << m
	// 处理了mask（表示的机关）,结束位置在i，花费的最短步数
	dp := make([][]int, X)
	for i := 0; i < X; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = INF

		}
	}
	for i := 0; i < m; i++ {
		dp[1<<i][i] = startDist[i]
	}

	for mask := 1; mask < X; mask++ {
		for i := 0; i < m; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			// dp[mask][i]
			for j := 0; j < m; j++ {
				if mask&(1<<j) > 0 {
					continue
				}
				next := mask | (1 << j)
				dp[next][j] = min(dp[next][j], dp[mask][i]+pairDist[i][j])
			}
		}
	}
	var best = -1

	for j := 0; j < m; j++ {
		tmp := dp[X-1][j] + endDist[j]
		if best < 0 || tmp < best {
			best = tmp
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findPosition(maze []string, x byte) []Pair {
	var res []Pair

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == x {
				res = append(res, Pair{i, j})
			}
		}
	}
	return res
}

type Pair struct {
	first, second int
}

func (pair Pair) Pos(n int) int {
	return pair.second + pair.first*n
}
