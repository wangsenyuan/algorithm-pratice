package p3447

import "slices"

func assignElements(groups []int, elements []int) []int {
	n := len(groups)
	mem := make(map[int][]int)

	for i, x := range groups {
		mem[x] = append(mem[x], i)
	}
	//m := len(elements)
	ans := make([]int, n)
	for i := range n {
		ans[i] = -1
	}

	mx := slices.Max(groups)

	pos := make(map[int]int)

	vis := make(map[int]bool)

	for i, e := range elements {
		if vis[e] {
			continue
		}
		vis[e] = true
		for u := e; u <= mx; u += e {
			for pos[u] < len(mem[u]) {
				j := mem[u][pos[u]]
				ans[j] = i
				pos[u]++
			}
		}
		if e == 1 {
			break
		}
	}

	return ans
}
