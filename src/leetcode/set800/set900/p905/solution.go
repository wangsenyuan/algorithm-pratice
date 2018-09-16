package p905

func sortArrayByParity(A []int) []int {
	n := len(A)
	if n == 0 {
		return A
	}
	B := make([]int, n)

	var j int
	for i := 0; i < n; i++ {
		if A[i]&1 == 0 {
			B[j] = A[i]
			j++
		}
	}

	for i := 0; i < n; i++ {
		if A[i]&1 == 1 {
			B[j] = A[i]
			j++
		}
	}
	return B
}
