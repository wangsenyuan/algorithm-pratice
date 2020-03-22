package p1386

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := make([]int, 0, len(reservedSeats))

	ii := make(map[int]int)

	for _, seat := range reservedSeats {
		r, c := seat[0]-1, seat[1]-1
		if j, found := ii[r]; found {
			rows[j] |= 1 << c
		} else {
			rows = append(rows, 1<<c)
			ii[r] = len(rows) - 1
		}
	}

	var res int

	for _, row := range rows {
		res += check(row)
		n--
	}

	return res + n*2
}

func check(row int) int {
	var res int

	process := func(i int) int {
		if i+4 > 10 {
			return 10
		}
		j := i
		for j < i+4 && row&(1<<j) == 0 {
			j++
		}
		if j == i+4 {
			res++
		}
		return j
	}

	i := process(1)

	if i < 5 {
		i = process(3)
	}
	if i < 7 {
		process(5)
	}

	return res
}
