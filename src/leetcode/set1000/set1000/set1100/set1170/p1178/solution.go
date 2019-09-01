package p1178

func findNumOfValidWords(words []string, puzzles []string) []int {
	// len(puzzles[i]) <= 7
	mem := make(map[int][]Node)

	for i := 0; i < len(words); i++ {
		w := words[i]
		mask := strToIntMask(w)
		node := Node{w, mask, i}

		if tmp, found := mem[mask]; found {
			mem[mask] = append(tmp, node)
		} else {
			mem[mask] = []Node{node}
		}
	}

	res := make([]int, len(puzzles))

	for u, puzzle := range puzzles {
		var mask int
		var bits []uint
		for i := 0; i < len(puzzle); i++ {
			x := uint(puzzle[i] - 'a')
			if mask&(1<<x) == 0 {
				bits = append(bits, x)
				mask |= 1 << x
			}
		}

		n := len(bits)
		N := 1 << uint(n)

		for state := 1; state < N; state++ {
			var tmp int

			for j := 0; j < n; j++ {
				if state&(1<<uint(j)) > 0 {
					tmp |= 1 << bits[j]
				}
			}

			if nodes, found := mem[tmp]; found {
				for _, node := range nodes {
					first := uint(puzzle[0] - 'a')
					if node.mask&(1<<first) > 0 {
						// first condition meet
						res[u]++
					}
				}
			}
		}
	}

	return res
}

type Node struct {
	s     string
	mask  int
	index int
}

func strToIntMask(s string) int {
	var res int

	for i := 0; i < len(s); i++ {
		x := s[i] - 'a'
		res |= 1 << uint(x)
	}

	return res
}

// type Nodes []Node

// func (nodes Nodes) Len() int {
// 	return len(nodes)
// }

// func (nodes Nodes) Less(i, j int) bool {
// 	return nodes[i].mask < nodes[j].mask
// }

// func (nodes Nodes) Swap(i, j int) {
// 	nodes[i], nodes[j] = nodes[j], nodes[i]
// }
