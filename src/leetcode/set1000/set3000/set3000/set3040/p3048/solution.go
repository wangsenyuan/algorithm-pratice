package p3048

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	var res int

	find := func(i, j int) int {
		x0 := max(bottomLeft[i][0], bottomLeft[j][0])
		x1 := min(topRight[i][0], topRight[j][0])
		y0 := max(bottomLeft[i][1], bottomLeft[j][1])
		y1 := min(topRight[i][1], topRight[j][1])
		if x1 <= x0 || y1 <= y0 {
			return 0
		}
		return min(x1-x0, y1-y0)
	}

	n := len(bottomLeft)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = max(res, find(i, j))
		}
	}

	return int64(res) * int64(res)
}
