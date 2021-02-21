package p1766

var dd = []int{-1, 0, 1, 0, -1}

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	height := make([][]int, m)
	vis := make([][]bool, m)

	que := make([]int, m*n)

	var end int

	for i := 0; i < m; i++ {
		height[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				que[end] = i*n + j
				end++
				vis[i][j] = true
				height[i][j] = 0
			} else {
				height[i][j] = m + n
			}
		}
	}

	var front int
	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && !vis[u][v] {
				height[u][v] = min(height[u][v], height[x][y]+1)
				que[end] = u*n + v
				end++
				vis[u][v] = true
			}
		}
	}

	return height
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
