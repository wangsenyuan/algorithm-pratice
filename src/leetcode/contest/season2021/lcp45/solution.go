package lcp45

import "sort"

const H = 201

var dd = []int{-1, 0, 1, 0, -1}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	m := len(terrain)
	n := len(terrain[0])
	vis := make([][][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			vis[i][j] = make([]bool, H)
		}
	}

	que := make([]int, m*n*H)
	vis[position[0]][position[1]][1] = true
	var front, end int
	que[end] = position[0]*n*H + position[1]*H + 1
	end++
	var res [][]int
	for front < end {
		cur := que[front]
		front++
		x, y, s := cur/(n*H), (cur%(n*H))/H, cur%H

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n {
				d := terrain[x][y] - (terrain[u][v] + obstacle[u][v])
				s2 := s + d
				if s2 <= 0 {
					continue
				}
				if !vis[u][v][s2] {
					if s2 == 1 {
						res = append(res, []int{u, v})
					}
					vis[u][v][s2] = true
					que[end] = u*n*H + v*H + s2
					end++
				}
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0] || res[i][0] == res[j][0] && res[i][1] < res[j][1]
	})

	return res
}
