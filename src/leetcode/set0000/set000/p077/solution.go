package p077

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}
	if k == 0 {
		return [][]int{nil}
	}
	a := combine(n-1, k)
	b := combine(n-1, k-1)
	res := make([][]int, len(a), len(a)+len(b))
	copy(res, a)

	for _, x := range b {
		y := make([]int, len(x)+1)
		copy(y, x)
		y[len(y)-1] = n
		res = append(res, y)
	}

	return res
}
