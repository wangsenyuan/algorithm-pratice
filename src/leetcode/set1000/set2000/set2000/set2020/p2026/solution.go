package p2026

func minimumMoves(s string) int {
	var res int

	var i int

	for i < len(s) {
		if s[i] == 'X' {
			res++
			i += 3
		} else {
			i++
		}
	}

	return res
}
