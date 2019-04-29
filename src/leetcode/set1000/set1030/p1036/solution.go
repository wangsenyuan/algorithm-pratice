package p1036

var dd = []int64{-1, 0, 1, 0, -1}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	const COL = 1e6
	N := len(blocked)

	set := make(map[int64]bool)

	for _, block := range blocked {
		x, y := block[0], block[1]
		set[int64(x)*COL+int64(y)] = true
	}

	bfs := func(start []int, end []int) bool {
		vis := make(map[int64]bool)

		var front int
		var que []int64
		que = append(que, int64(start[0])*COL+int64(start[1]))

		vis[int64(start[0])*COL+int64(start[1])] = true

		var level int
		for front < len(que) {
			tmp := len(que)
			for front < tmp {
				cur := que[front]
				front++
				x, y := cur/COL, cur%COL
				if x == int64(end[0]) && y == int64(end[1]) {
					return true
				}

				for k := 0; k < 4; k++ {
					u, v := x+dd[k], y+dd[k+1]
					if u >= 0 && u < COL && v >= 0 && v < COL && !vis[int64(u)*COL+int64(v)] && !set[int64(u)*COL+int64(v)] {
						que = append(que, int64(u)*COL+int64(v))
						vis[int64(u)*COL+int64(v)] = true
					}
				}
			}
			level++
			if level > 2*N {
				return true
			}
		}

		return false
	}

	return bfs(source, target) && bfs(target, source)
}
