package lcp74

import "sort"

func fieldOfGreatestBlessing(forceField [][]int) int {
	// 找出被覆盖最多的区域
	// 如果只考虑x轴，把左右坐标投影到x轴上，那么对于任何一个点
	// 覆盖区域，可以认为就在某条边上， 因为这个区域肯定有条边， 这些边肯定属于某个rect
	// n <= 100, 那么左右两条边的 2 * n <= 200
	// 200 * 200 * 200 * 200 ~ 10e9, 不行
	// 第四条边可以不用
	type Line struct {
		pos int
		id  int
		tp  int
	}

	var xs []Line

	for i, field := range forceField {
		x, w := field[0], field[2]
		x *= 2
		xs = append(xs, Line{x - w, i, -1})
		xs = append(xs, Line{x + w, i, 1})
	}

	// 先enter， 再leave
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].pos < xs[j].pos || xs[i].pos == xs[j].pos && xs[i].tp < xs[j].tp
	})

	addLine := func(ys []Line, id int) []Line {
		field := forceField[id]
		y, w := field[1], field[2]
		y *= 2
		ys = append(ys, Line{y - w, id, -1})
		ys = append(ys, Line{y + w, id, 1})
		return ys
	}

	removeLine := func(ys []Line, id int) []Line {
		n := len(ys)
		for i := n - 1; i >= 0; i-- {
			// 有两个
			if ys[i].id == id {
				copy(ys[i:], ys[i+1:])
			}
		}

		return ys[:n-2]
	}

	var best int
	var ys []Line
	occ := make(map[int]int)
	for _, x := range xs {
		occ[x.id]++
		if occ[x.id] == 1 {
			// enters
			ys = addLine(ys, x.id)
			sort.Slice(ys, func(i, j int) bool {
				return ys[i].pos < ys[j].pos || ys[i].pos == ys[j].pos && ys[i].tp < ys[j].tp
			})
			cnt := make(map[int]int)
			for _, y := range ys {
				cnt[y.id]++
				if cnt[y.id] == 2 {
					delete(cnt, y.id)
				}
				best = max(best, len(cnt))
			}
		} else {
			// leaves
			ys = removeLine(ys, x.id)
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
