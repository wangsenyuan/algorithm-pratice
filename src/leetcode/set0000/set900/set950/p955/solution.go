package p955

func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	n := len(A)
	m := len(A[0])
	if m == 0 {
		return 0
	}
	var res int
	ss := make([]string, n)

	for j := 0; j < m; j++ {
		// check wether to delete this column or not
		i := 1
		for i < n && A[i-1][j] <= A[i][j] {
			i++
		}
		if i == n {
			// this column is already sorted
			for k := 0; k < n; k++ {
				ss[k] = ss[k] + string(A[k][j])
			}
			continue
		}
		//determine whether need to delete
		i = 1
		for i < n {
			if ss[i-1] < ss[i] || A[i-1][j] <= A[i][j] {
				i++
				continue
			}
			// ss[i-1] == ss[i] && A[i-1][j] > A[i][j]
			break
		}

		if i == n {
			for k := 0; k < n; k++ {
				ss[k] = ss[k] + string(A[k][j])
			}
			continue
		}
		res++
	}
	return res
}
