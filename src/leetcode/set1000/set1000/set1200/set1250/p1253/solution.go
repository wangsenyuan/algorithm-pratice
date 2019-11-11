package p1253

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)

	matrix := make([][]int, 2)
	matrix[0] = make([]int, n)
	matrix[1] = make([]int, n)

	var diff int

	for i := n - 1; i >= 0; i-- {
		if colsum[i] == 0 {
			// both zeros
		} else if colsum[i] == 2 {
			// both ones
			matrix[0][i] = 1
			matrix[1][i] = 1
			upper--
			lower--
		} else {
			// only one one
			matrix[0][i] = -1
			matrix[1][i] = -1
			diff++
		}
	}

	if upper < 0 || lower < 0 {
		return nil
	}

	if upper == 0 && lower == 0 {
		if diff > 0 {
			return nil
		}
		return matrix
	}

	if diff == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] < 0 {
			if upper > lower {
				matrix[0][i] = 1
				matrix[1][i] = 0
				upper--
			} else {
				matrix[0][i] = 0
				matrix[1][i] = 1
				lower--
			}
			diff--
			if lower < 0 {
				return nil
			}
		}
	}

	if diff != 0 || upper != 0 || lower != 0 {
		return nil
	}

	return matrix
}
