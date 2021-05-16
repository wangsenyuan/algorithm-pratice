package p1862

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
	}

	for c := 0; c < m; c++ {
		pos := n
		for r := n - 1; r >= 0; r-- {
			res[r][c] = '.'
			if box[m-1-c][r] == '*' {
				res[r][c] = '*'
				pos = r
				continue
			}
			if box[m-1-c][r] == '#' {
				res[pos-1][c] = '#'
				pos--
			}
		}
	}

	return res
}
