package p413

func numberOfArithmeticSlices(A []int) int {
	res := 0

	start := 0
	for start+2 < len(A) {
		offset := A[start+1] - A[start]

		i, count := start+2, 0
		for i < len(A) && A[i]-A[i-1] == offset {
			count += (i - 1 - start)
			i++
		}
		if i-start >= 3 {
			res += count
		}
		start = i - 1
	}

	return res
}
