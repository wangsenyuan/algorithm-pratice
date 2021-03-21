package p1796

func secondHighest(s string) int {
	occ := make([]bool, 10)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			occ[int(s[i]-'0')] = true
		}
	}

	for x := 9; x > 0; x-- {
		if occ[x] {
			for y := x - 1; y >= 0; y-- {
				if occ[y] {
					return y
				}
			}
			break
		}
	}
	return -1
}
