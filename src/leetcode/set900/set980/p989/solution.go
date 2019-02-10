package p989

func addToArrayForm(A []int, K int) []int {
	n := len(A)
	B := make([]int, n+11)
	copy(B[11:], A)
	i := n + 10
	j := n - 1
	for j >= 0 || K > 0 {
		B[i] += K
		if B[i] >= 10 {
			K = B[i] / 10
			B[i] %= 10
		} else {
			K = 0
		}
		i--
		j--
	}

	return B[i+1:]
}
