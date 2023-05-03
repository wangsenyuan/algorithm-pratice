package lcp61

func temperatureTrend(A []int, B []int) int {
	a := getTrend(A)
	b := getTrend(B)

	n := len(a)
	res := 0
	prev := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			tmp := i - prev
			res = max(res, tmp)
			prev = i + 1
		} else {
			res = max(res, i-prev+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getTrend(A []int) []int {
	n := len(A)
	a := make([]int, n-1)
	for i := 0; i+1 < n; i++ {
		if A[i+1] > A[i] {
			a[i] = 1
		} else if A[i+1] < A[i] {
			a[i] = -1
		} else {
			a[i] = 0
		}
	}
	return a
}
