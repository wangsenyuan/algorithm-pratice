package p2096

func validArrangement(pairs [][]int) [][]int {
	// no starti < endi
	// 给定一个pair， 假设 [x, y], 那么必然有一些 [., x], 有一些是 [y, .]
	deg := make(map[int]int)
	out := make(map[int][]int)
	for i, p := range pairs {
		deg[p[0]]++
		deg[p[1]]--
		if out[p[0]] == nil {
			out[p[0]] = make([]int, 0, 2)
		}
		out[p[0]] = append(out[p[0]], i)
	}
	// find odd one
	var start int = -1
	for k, v := range deg {
		if v == 1 {
			start = k
			break
		}
	}

	if start < 0 {
		// all even
		for k := range deg {
			if len(out[k]) > 0 {
				start = k
				break
			}
		}
	}
	var stack []int
	pos := make(map[int]int)
	stack = append(stack, start)
	path := make([]int, 0, len(out))

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		if pos[u] == len(out[u]) {
			path = append(path, u)
			stack = stack[:len(stack)-1]
		} else {
			i := out[u][pos[u]]
			v := pairs[i][1]
			pos[u]++
			stack = append(stack, v)
		}
	}

	var res [][]int

	for i := len(path) - 2; i >= 0; i-- {
		res = append(res, []int{path[i+1], path[i]})
	}

	return res
}
