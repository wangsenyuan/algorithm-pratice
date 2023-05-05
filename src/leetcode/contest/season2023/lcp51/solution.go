package lcp51

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	// brute force
	// len(mat) == 5
	n := len(cookbooks)
	N := 1 << n
	best := -1
	cnt := make([]int, 5)
	for state := 1; state < N; state++ {
		for i := 0; i < 5; i++ {
			cnt[i] = 0
		}
		var a, b int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				a += attribute[i][0]
				b += attribute[i][1]
				for j := 0; j < 5; j++ {
					cnt[j] += cookbooks[i][j]
				}
			}
		}

		if b < limit {
			continue
		}
		// b >= limit
		ok := true
		for j := 0; j < 5 && ok; j++ {
			ok = materials[j] >= cnt[j]
		}
		if !ok {
			continue
		}
		if best < 0 || best < a {
			best = a
		}
	}

	return best
}
