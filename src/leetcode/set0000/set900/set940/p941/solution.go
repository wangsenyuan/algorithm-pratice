package p941

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	n := len(A)

	var i int

	for i < n-1 && A[i] < A[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}
	for i < n-1 && A[i] > A[i+1] {
		i++
	}
	return i == n-1
}
