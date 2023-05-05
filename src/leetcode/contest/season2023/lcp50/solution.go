package lcp50

func giveGem(gem []int, operations [][]int) int {
	for _, op := range operations {
		x, y := op[0], op[1]
		half := gem[x] / 2
		gem[x] -= half
		gem[y] += half
	}

	a := 0
	b := 1 << 30

	for _, v := range gem {
		if v > a {
			a = v
		}
		if v < b {
			b = v
		}
	}

	return a - b
}
