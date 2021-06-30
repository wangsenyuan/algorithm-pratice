package p1914

func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	arr := make([]int, 2*(m+n))

	var loop func(a, b, c, d int)

	loop = func(a, b, c, d int) {
		if a > c || b > d {
			return
		}
		var l int
		i, j := a, b
		for j <= d {
			arr[l] = grid[i][j]
			l++
			j++
		}
		j--
		i++
		for i <= c {
			arr[l] = grid[i][j]
			l++
			i++
		}
		i--
		j--
		for j >= b {
			arr[l] = grid[i][j]
			l++
			j--
		}
		j++
		i--
		for i > a {
			arr[l] = grid[i][j]
			l++
			i--
		}

		rightShift(arr[:l], k%l)

		i, j = a, b
		l = 0

		for j <= d {
			grid[i][j] = arr[l]
			l++
			j++
		}
		j--
		i++
		for i <= c {
			grid[i][j] = arr[l]
			l++
			i++
		}
		i--
		j--
		for j >= b {
			grid[i][j] = arr[l]
			l++
			j--
		}
		j++
		i--
		for i > a {
			grid[i][j] = arr[l]
			l++
			i--
		}

		loop(a+1, b+1, c-1, d-1)
	}

	loop(0, 0, m-1, n-1)

	return grid
}

func rightShift(arr []int, k int) {
	rotate(arr[:k])
	rotate(arr[k:])
	rotate(arr)
}

func rotate(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
