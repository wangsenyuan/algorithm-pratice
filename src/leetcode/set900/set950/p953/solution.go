package p953

func isAlienSorted(words []string, order string) bool {
	if len(words) <= 1 {
		return true
	}
	ii := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		ii[order[i]] = i
	}

	less := func(a, b string) bool {
		var i int

		for i < len(a) && i < len(b) {
			if ii[a[i]] < ii[b[i]] {
				return true
			} else if ii[a[i]] > ii[b[i]] {
				return false
			}
			i++
		}
		return i == len(a)
	}

	for i := 1; i < len(words); i++ {
		a := words[i-1]
		b := words[i]
		if !less(a, b) {
			return false
		}
	}
	return true
}
