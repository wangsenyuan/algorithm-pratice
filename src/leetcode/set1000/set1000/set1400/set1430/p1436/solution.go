package p1436

func destCity(paths [][]string) string {
	mem := make(map[string]int)

	for _, path := range paths {
		a := path[0]
		mem[a]++
	}

	for _, path := range paths {
		b := path[1]
		if mem[b] == 0 {
			return b
		}
	}

	return ""
}
