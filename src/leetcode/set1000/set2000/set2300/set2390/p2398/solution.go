package p2398

func maximumRows(mat [][]int, cols int) int {
	// n := len(mat)
	m := len(mat[0])

	check := func(mask int) int {
		var res int
		for _, row := range mat {
			ok := true
			for j := 0; j < m; j++ {
				if row[j] == 1 {
					if (mask>>j)&1 == 0 {
						ok = false
						break
					}
				}
			}
			if ok {
				res++
			}
		}
		return res
	}

	var best int

	for mask := 1; mask < 1<<m; mask++ {
		cnt := digitCount(mask)
		if cnt == cols {
			best = max(best, check(mask))
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
