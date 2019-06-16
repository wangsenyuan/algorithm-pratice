package p896

func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return true
	}
	return isIncreasing(A) || isDescreasing(A)
}

func isIncreasing(A []int) bool {
	n := len(A)
	for i := 1; i < n; i++ {
		if A[i] < A[i-1] {
			return false
		}
	}
	return true
}

func isDescreasing(A []int) bool {
	n := len(A)
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			return false
		}
	}
	return true
}
