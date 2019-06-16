package p845

func longestMountain(A []int) int {
	n := len(A)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	var best int
	for i := 0; i < n; i++ {
		tmp := 1 + left[i] + right[i]
		if left[i] > 0 && right[i] > 0 && tmp > best {
			best = tmp
		}
	}

	return best
}
