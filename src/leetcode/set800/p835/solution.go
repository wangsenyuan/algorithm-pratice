package p835

func largestOverlap(A [][]int, B [][]int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	m := len(A[0])

	if m == 0 {
		return 0
	}

	X := convert(A)
	Y := convert(B)

	var best int
	for i := -n + 1; i < n; i++ {
		for j := -m + 1; j < m; j++ {
			Z := transform(X, i, j, n, m)
			tmp := compare(Z, Y)
			if tmp > best {
				best = tmp
			}
		}
	}
	return best
}

func compare(X []int, Y []int) int {
	var cnt int

	for i := 0; i < len(X); i++ {
		tmp := X[i] & Y[i]
		for tmp > 0 {
			if tmp&1 == 1 {
				cnt++
			}
			tmp >>= 1
		}
	}

	return cnt
}

func convert(A [][]int) []int {
	n := len(A)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		for j := 0; j < len(A[i]); j++ {
			tmp = tmp*2 + A[i][j]
		}
		res[i] = tmp
	}
	return res
}

func transform(X []int, i, j, n, m int) []int {
	Y := make([]int, n)
	for k := 0; k < n; k++ {
		if k+i >= 0 && k+i < n {
			if j >= 0 {
				Y[k] = X[k+i] >> uint(j)
			} else {
				Y[k] = X[k+i] << uint(-j)
			}
		}
	}
	return Y
}
