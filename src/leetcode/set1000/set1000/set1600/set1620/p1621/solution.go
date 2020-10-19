package p1621


func areConnected(n int, threshold int, queries [][]int) []bool {
	// if X % y == 0, then (a * X) % y == 0
	// a0 -> a1 -> ... -> b1 -> b0
	// 2 * n
	// 3 * n
	// 5 * n
	arr := make([]int, n+1)
	cnt := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	union := func(a, b int) {
		pa := find(a)
		pb := find(b)
		if pa == pb {
			return
		}
		if cnt[pa] < cnt[pb] {
			pa, pb = pb, pa
		}
		cnt[pa] += cnt[pb]
		arr[pb] = pa
	}

	for j := threshold + 1; j <= n; j++ {
		for i := j; i <= n; i += j {
			union(i, j)
		}
	}

	ans := make([]bool, len(queries))

	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		ans[i] = find(a) == find(b)
	}

	return ans
}
