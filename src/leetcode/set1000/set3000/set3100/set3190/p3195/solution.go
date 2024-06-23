package p3195

func minimumArea(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	t, l := n, m
	b, r := 0, 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				t = min(t, i)
				l = min(l, j)
				b = max(b, i)
				r = max(r, j)
			}
		}
	}
	if t <= b {
		return (b - t + 1) * (r - l + 1)
	}
	return 0
}
