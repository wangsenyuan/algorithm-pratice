package p2201

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	mem := make(map[int]bool)
	for _, coord := range dig {
		r, c := coord[0], coord[1]
		key := r*n + c
		mem[key] = true
	}

	var res int

	for _, art := range artifacts {
		ok := true
		for r := art[0]; r <= art[2] && ok; r++ {
			for c := art[1]; c <= art[3] && ok; c++ {
				k := r*n + c
				ok = mem[k]
			}
		}
		if ok {
			res++
		}
	}
	return res
}
