package p1018

func prefixesDivBy5(A []int) []bool {
	n := len(A)
	res := make([]bool, n)

	var cur int

	for i := 0; i < n; i++ {
		cur = cur*2 + A[i]
		cur %= 5
		if cur == 0 {
			res[i] = true
		}
	}

	return res
}
