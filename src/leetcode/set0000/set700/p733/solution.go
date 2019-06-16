package p733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	if n == 0 {
		return image
	}
	m := len(image[0])
	if m == 0 {
		return image
	}

	result := make([][]int, n, m)
	checked := make([][]bool, n, m)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		checked[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			result[i][j] = image[i][j]
		}
	}
	startingColor := image[sr][sc]
	ds := []int{-1, 0, 1, 0, -1}

	var fill func(i, j int)

	fill = func(i, j int) {
		checked[i][j] = true
		result[i][j] = newColor
		for k := 0; k < 4; k++ {
			x, y := i+ds[k], j+ds[k+1]
			if x >= 0 && x < n && y >= 0 && y < m && image[x][y] == startingColor && !checked[x][y] {
				fill(x, y)
			}
		}
		//image[i][j] = -image[i][j]
	}

	fill(sr, sc)

	return result
}
