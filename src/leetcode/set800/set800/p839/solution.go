package p839

func numSimilarGroups(A []string) int {
	n := len(A)
	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if set[i] != i {
			set[i] = find(set[i])
		}
		return set[i]
	}

	union := func(i, j int) {
		pi := find(i)
		pj := find(j)
		if pi != pj {
			set[pi] = pj
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if similar(A[i], A[j]) {
				union(i, j)
			}
		}
	}

	cnt := make(map[int]bool)

	for i := 0; i < n; i++ {
		cnt[find(i)] = true
	}
	return len(cnt)
}

func similar(a, b string) bool {
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff <= 2
}
