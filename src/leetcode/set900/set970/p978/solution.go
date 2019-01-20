package p978

func maxTurbulenceSize(A []int) int {
	n := len(A)
	if n == 1 {
		return 1
	}
	if n == 2 {
		if A[0] != A[1] {
			return 2
		}
		return 1
	}
	B := make([]int, n)
	for i := 0; i < n-1; i++ {
		if A[i] < A[i+1] {
			B[i] = 1
		} else if A[i] > A[i+1] {
			B[i] = -1
		}
	}
	var best = 1
	var j = 0
	for j < n-1 {
		if B[j] != 0 {
			var i = j
			for i < n-1 {
				if (i-j)%2 == 0 && B[i] != B[j] {
					break
				}

				if (i-j)%2 == 1 && B[i] != -B[j] {
					break
				}

				i++
			}
			best = max(best, i-j+1)
			j = i
		} else {
			j++
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
