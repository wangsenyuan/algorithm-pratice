package p799

func champagneTower(poured int, query_row int, query_glass int) float64 {
	f := make([][]float64, 102)
	for i := 0; i < 102; i++ {
		f[i] = make([]float64, 102)
	}
	f[0][0] = float64((poured))
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			x := 0.5 * (f[i][j] - 1.0)
			if x > 0 {
				f[i+1][j] += x
				f[i+1][j+1] += x
			}
		}
	}
	res := f[query_row][query_glass]
	if res > 1.0 {
		res = 1.0
	}
	return res
}
