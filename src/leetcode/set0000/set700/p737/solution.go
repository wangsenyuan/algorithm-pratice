package p737

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}
	dict := make(map[string]int)
	index := 0

	set := make(map[int]int)

	var find func(x int) int

	find = func(x int) int {
		if px, found := set[x]; !found {
			set[x] = x
		} else if px != x {
			set[x] = find(px)
		}

		return set[x]
	}

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if _, found := dict[a]; !found {
			dict[a] = index
			index++
		}
		if _, found := dict[b]; !found {
			dict[b] = index
			index++
		}

		i, j := dict[a], dict[b]
		pi, pj := find(i), find(j)
		if pi != pj {
			set[pi] = pj
		}
	}

	for i := 0; i < len(words1); i++ {
		a, b := words1[i], words2[i]
		if a == b {
			continue
		}

		x, founda := dict[a]
		if !founda {
			return false
		}
		y, foundb := dict[b]
		if !foundb {
			return false
		}

		if find(x) != find(y) {
			return false
		}
	}
	return true
}
