package p2157

func groupStrings(words []string) []int {
	// n <= 2 * 10000
	n := len(words)

	set := make(map[int]int)
	sz := make(map[int]int)
	var find func(x int) int
	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	var tot int

	union := func(a, b int) {
		pa := find(a)
		pb := find(b)
		if pa == pb {
			return
		}

		if sz[pa] < sz[pb] {
			pa, pb = pb, pa
		}
		sz[pa] += sz[pb]
		set[pb] = pa
		tot--
	}
	cnt := make(map[int]int)
	rep := make([]int, n)
	for i, word := range words {
		rep[i] = as_mask(word)
		sz[rep[i]]++
		set[rep[i]] = rep[i]
		cnt[rep[i]]++
	}

	tot = len(set)

	for _, cur := range rep {
		for i := 0; i < 26; i++ {
			// add or remove one
			tmp := cur ^ (1 << i)
			if cnt[tmp] > 0 {
				union(tmp, cur)
			}
			if (cur>>i)&1 == 1 {
				for j := 0; j < 26; j++ {
					if (cur>>j)&1 == 0 {
						tmp = cur ^ (1 << i) ^ (1 << j)
						if cnt[tmp] > 0 {
							union(tmp, cur)
						}
					}
				}
			}
		}
	}
	var largest int
	for i := 0; i < n; i++ {
		p := find(rep[i])
		if sz[p] > largest {
			largest = sz[p]
		}
	}

	return []int{tot, largest}
}

func as_mask(w string) int {
	var res int
	for i := 0; i < len(w); i++ {
		x := int(w[i] - 'a')
		res |= 1 << x
	}
	return res
}
