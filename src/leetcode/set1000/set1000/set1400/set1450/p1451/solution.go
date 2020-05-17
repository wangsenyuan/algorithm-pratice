package p1451

func peopleIndexes(favoriteCompanies [][]string) []int {
	n := len(favoriteCompanies)

	fs := make([]map[string]bool, n)

	for i := 0; i < n; i++ {
		fs[i] = make(map[string]bool)

		for _, s := range favoriteCompanies[i] {
			fs[i][s] = true
		}
	}

	check := func(i, j int) bool {
		a := fs[i]
		b := fs[j]

		for k := range a {
			if !b[k] {
				return false
			}
		}

		return true
	}
	var res []int
	for i := 0; i < n; i++ {
		var can bool = true
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if check(i, j) {
				// same
				can = false
				break
			}
		}
		if can {
			res = append(res, i)
		}
	}
	return res
}
