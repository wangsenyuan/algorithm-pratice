package p942

func diStringMatch(S string) []int {
	n := len(S)
	res := make([]int, n+1)
	x, y := 0, n

	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			res[i] = x
			x++
		} else {
			res[i] = y
			y--
		}
	}
	res[n] = x
	return res
}
