package p957

func prisonAfterNDays(cells []int, N int) []int {
	seen := make([]int, 1<<8)

	for i := 0; i < 1<<8; i++ {
		seen[i] = -1
	}

	bak := make([]int, len(cells))

	for N > 0 {
		state := toNum(cells)
		if seen[state] >= 0 {
			N %= seen[state] - N
		}
		seen[state] = N
		if N >= 1 {
			N--
			next(cells, bak)
		}
	}

	return cells
}

func next(cells []int, bak []int) {
	copy(bak, cells)
	cells[0] = 0
	for i := 1; i < len(cells)-1; i++ {
		if bak[i-1] == bak[i+1] {
			cells[i] = 1
		} else {
			cells[i] = 0
		}
	}
	cells[len(cells)-1] = 0
}

func toNum(cells []int) int {
	var num int
	for i := 0; i < len(cells); i++ {
		if cells[i] == 1 {
			num |= 1 << uint(i)
		}
	}
	return num
}
