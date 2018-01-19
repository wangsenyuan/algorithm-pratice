package p054

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return nil
	}

	m := len(matrix[0])
	if m == 0 {
		return nil
	}
	res := make([]int, n*m)

	i, j := 0, 0
	dx, dy := 0, 1
	a, b, c, d := 0, -1, n, m
	for k := 0; k < n*m; k++ {
		res[k] = matrix[i][j]
		i, j = i+dx, j+dy
		if dy == 1 {
			if j == d {
				//turn down
				dx, dy = 1, 0
				d--
				i++
				j--
			}
		} else if dx == 1 {
			if i == c {
				//turn left
				dx, dy = 0, -1
				c--
				i--
				j--
			}
		} else if dy == -1 {
			if j == b {
				//turn up
				dx, dy = -1, 0
				b++
				i--
				j++
			}
		} else if dx == -1 {
			if i == a {
				//turn right
				dx, dy = 0, 1
				a++
				i++
				j++
			}
		}
	}

	return res
}
