package p3307

func kthCharacter(k int64, operations []int) byte {
	// k is too large
	// op = 0, 表示增加长度
	// op = 1, 表示增加并追加
	// 需要知道k在哪个区间

	var dfs func(r int, i int, k int) int

	dfs = func(r int, i int, k int) int {
		if r == 1 {
			// 只能是a
			return 0
		}
		v := operations[i]
		nk := k
		if k >= r/2 {
			nk = k - r/2
		}
		if v == 0 || nk == k {
			return dfs(r/2, i-1, nk)
		}
		// 后半段才需要增加1
		return 1 + dfs(r/2, i-1, nk)
	}
	kk := int(k) - 1
	r := 1
	var i int
	for r <= kk {
		r *= 2
		i++
	}

	res := dfs(r, i-1, kk)
	res %= 26

	return byte('a' + res)
}
