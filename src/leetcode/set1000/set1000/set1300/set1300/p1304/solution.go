package p1304

func sumZero(n int) []int {
	if n == 0 {
		return nil
	}

	res := make([]int, n)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i] = i + 1
		res[j] = -(i + 1)
	}
	return res
}
