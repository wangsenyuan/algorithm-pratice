package p775

func isIdealPermutation(A []int) bool {
	if len(A) < 3 {
		return true
	}

	for i := 0; i < len(A); i++ {
		if A[i] == i {
			continue
		}
		if i-1 >= 0 && A[i-1] == i {
			continue
		}

		if i+1 < len(A) && A[i+1] == i {
			continue
		}
		return false
	}

	return true
}
