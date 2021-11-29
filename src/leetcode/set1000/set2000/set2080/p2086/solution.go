package p2086

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	// m, n <= 10^^5
	// we can't explore all cells
	var res int
	if startPos[0] < homePos[0] {
		for r := startPos[0]; r < homePos[0]; r++ {
			res += rowCosts[r+1]
		}
	} else if startPos[0] > homePos[0] {
		for r := startPos[0]; r > homePos[0]; r-- {
			res += rowCosts[r-1]
		}
	}

	if startPos[1] < homePos[1] {
		for c := startPos[1]; c < homePos[1]; c++ {
			res += colCosts[c+1]
		}
	} else {
		for c := startPos[1]; c > homePos[1]; c-- {
			res += colCosts[c-1]
		}
	}

	return res
}
