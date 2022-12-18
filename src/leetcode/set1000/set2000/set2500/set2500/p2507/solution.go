package p2507

func cycleLengthQueries(n int, queries [][]int) []int {

	lca := func(u, v int) int {
		if u < v {
			u, v = v, u
		}
		// 假设p是 u, v 的公共祖先
		// u / (2 ** x) = p
		// v / (2 ** y) = p
		// u 和 v在同一层的条件是 u >=

		var h int
		for 1<<(h+1) <= v {
			h++
		}

		for u/2 >= (1 << h) {
			u /= 2
		}

		if u == v {
			return u
		}
		for u/2 != v/2 {
			u /= 2
			v /= 2
		}
		return u / 2
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		a, b := cur[0], cur[1]
		c := lca(a, b)
		x := height(a) - height(c)
		y := height(b) - height(c)
		res[i] = x + y + 1
	}

	return res
}

func height(num int) int {
	var res int
	for num > 1 {
		res++
		num /= 2
	}
	return res
}
