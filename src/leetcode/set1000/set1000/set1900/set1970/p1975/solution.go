package p1975

func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)

	// negative count
	var cnt int
	minValue := 1 << 30
	var sum int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] < 0 {
				cnt++
			}
			if minValue > abs(matrix[i][j]) {
				minValue = abs(matrix[i][j])
			}
			sum += int64(abs(matrix[i][j]))
		}
	}

	if cnt%2 == 0 {
		// can turn all into positive
		return sum
	}

	return sum - int64(2*minValue)
}

func abs(num int) int {
	if num <= 0 {
		return -num
	}
	return num
}
