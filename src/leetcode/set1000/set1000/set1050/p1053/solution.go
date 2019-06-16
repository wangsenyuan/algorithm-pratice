package p1053

func prevPermOpt1(A []int) []int {
	/**
		*  next perm gt cur is, from right, find the index j, that A[j] < A[j+1]
		*  find the smallest number at index k after j, A[k] > A[j]
		*	 swap j & k,
		*  then swap [j+1 ... n]
		*
		*  first this problem, find the index j, that A[j] > A[j+1]
		*  then swap [j+1 ... n]
		* then find the largest number at index k, after j, that A[k] < A[j]
		* swap j & k
	  *
	*/
	n := len(A)
	j := n - 2
	for j >= 0 && A[j] <= A[j+1] {
		j--
	}

	if j < 0 {
		return A
	}

	B := make([]int, n)
	copy(B, A)

	k := j + 1
	for i := j + 2; i < n; i++ {
		if B[i] < B[j] && B[i] > B[k] {
			k = i
		}
	}
	B[j], B[k] = B[k], B[j]

	return B
}
