package p2113

func numberOfBeams(bank []string) int {
	m := len(bank)
	var res int
	var prev int

	for i := 0; i < m; i++ {
		var cnt int
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				cnt++
			}
		}
		if cnt > 0 {
			res += prev * cnt
			prev = cnt
		}
	}

	return res
}
