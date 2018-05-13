package p832

func flipAndInvertImage(A [][]int) [][]int {
	B := make([][]int, len(A))
	for i, row := range A {
		B[i] = make([]int, len(row))
		copy(B[i], row)
		for j, k := 0, len(row)-1; j < k; j, k = j+1, k-1 {
			B[i][j], B[i][k] = B[i][k], B[i][j]
		}

		for j := 0; j < len(row); j++ {
			B[i][j] = 1 - B[i][j]
		}
	}
	return B
}
