package p845

func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	n := len(A)
	left := make([]int, n)

	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	var right int
	var best int
	for i := n - 2; i > 0; i-- {
		if A[i] > A[i+1] {
			right++
		} else {
			right = 0
		}
		tmp := right + left[i] + 1
		if right > 0 && left[i] > 0 && tmp > best {
			best = tmp
		}
	}

	if best < 3 {
		return 0
	}

	return best
}
