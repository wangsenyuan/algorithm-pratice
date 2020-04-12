package p1409

func processQueries1(queries []int, m int) []int {
	n := len(queries)
	res := make([]int, n)

	nums := make([]int, m)

	for i := 0; i < m; i++ {
		nums[i] = i + 1
	}

	for i := 0; i < n; i++ {
		x := queries[i]
		var j int
		for nums[j] != x {
			j++
		}
		res[i] = j
		copy(nums[1:j+1], nums[0:j])
		nums[0] = x
	}
	return res
}

func processQueries(queries []int, m int) []int {
	M := 2*m + 1
	arr := make([]int, M+1)

	update := func(pos int, v int) {
		pos++
		for pos <= M {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}
		return res
	}

	ii := make(map[int]int)

	for i := 0; i < m; i++ {
		update(i+m+1, 1)
		ii[i+1] = i + m + 1
	}

	// arr 保存了每个数字移动的次数
	res := make([]int, len(queries))
	pos := m

	for i := 0; i < len(queries); i++ {
		x := queries[i]
		j := ii[x]

		rank := get(j - 1)

		res[i] = rank
		ii[x] = pos

		update(j, -1)
		update(pos, 1)

		pos--
	}

	return res
}
