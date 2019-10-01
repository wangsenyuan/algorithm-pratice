package contest

func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int)

	for _, num := range arr {
		occ[num]++
	}

	set := make(map[int]int)

	for _, c := range occ {
		if _, found := set[c]; found {
			return false
		}
		set[c] = 1
	}
	return true
}
