package p778

var dd = []int{-1, 0, 1, 0, -1}

func swimInWater(grid [][]int) int {
	n := len(grid)
	que := make([]int, n*n)
	dist := make([]int, n*n)

	check := func(time int) bool {
		if time < grid[0][0] {
			return false
		}
		for i := 0; i < n*n; i++ {
			dist[i] = -1
		}
		dist[0] = min(time, grid[0][0])
		head, tail := 0, 0
		que[tail] = 0
		tail++
		for head < tail {
			tmp := tail
			for head < tmp {
				v := que[head]
				head++
				if v == n*n-1 {
					return true
				}
				r, c := v/n, v%n
				for j := 0; j < 4; j++ {
					x, y := r+dd[j], c+dd[j+1]
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] <= time && dist[x*n+y] == -1 {
						que[tail] = x*n + y
						tail++
						dist[x*n+y] = dist[v] + 1
					}
				}
			}
		}
		return false
	}

	left := 0
	right := n*n + 1
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
