package p944

func minDeletionSize(A []string) int {
	var res int

	n := len(A)
	m := len(A[0])
	for j := 0; j < m; j++ {
		sorted := true
		for i := 1; i < n && sorted; i++ {
			sorted = A[i][j] >= A[i-1][j]
		}
		if !sorted {
			res++
		}
	}
	return res
}
