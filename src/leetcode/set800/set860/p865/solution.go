package p865

func shortestPathAllKeys(grid []string) int {
	dd := []int{-1, 0, 1, 0, -1}
	n := len(grid)
	m := len(grid[0])
	start := 0
	ks := make(map[byte]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				ks[grid[i][j]] = true
			}
			if grid[i][j] == '@' {
				start = i*m + j
			}
		}
	}
	kks := make([]byte, len(ks))
	var idx int
	for k := range ks {
		kks[idx] = k
		idx++
	}

	var gks func(i int, flag int)
	res := make([]string, 0, 100)
	buf := make([]byte, idx)
	gks = func(i int, flag int) {
		if i == idx {
			res = append(res, string(buf))
			return
		}

		for j := 0; j < idx; j++ {
			if flag&(1<<uint(j)) == 0 {
				buf[i] = byte('a' + j)
				gks(i+1, flag|(1<<uint(j)))
			}
		}
	}
	gks(0, 0)

	que := make([]int, n*m)
	visited := make([]int, n*m)
	dist := make([]int, n*m)
	bfs := func(cur int, expect byte, flag int, has map[byte]bool) (int, int) {
		front, end := 0, 0
		visited[cur] = flag
		que[end] = cur
		end++
		dist[cur] = 0
		for front < end {
			u := que[front]
			front++
			i, j := u/m, u%m
			for k := 0; k < 4; k++ {
				x, y := i+dd[k], j+dd[k+1]
				if x >= 0 && x < n && y >= 0 && y < m && visited[x*m+y] != flag {
					if grid[x][y] == '#' {
						continue
					}
					if grid[x][y] >= 'A' && grid[x][y] <= 'F' && has[grid[x][y]-'A'+'a'] == false {
						continue
					}
					dist[x*m+y] = dist[u] + 1
					que[end] = x*m + y
					end++
					visited[x*m+y] = flag
					if grid[x][y] == expect {
						return dist[x*m+y], x*m + y
					}
				}
			}
		}
		return -1, 0
	}

	var dfs func(cur int, key string, has map[byte]bool, idx int, flag *int, ans int) int
	best := -1
	dfs = func(cur int, key string, has map[byte]bool, idx int, flag *int, ans int) int {
		if best > 0 && ans > best {
			return -1
		}
		if idx == len(key) {
			return ans
		}
		*flag++
		dist, pos := bfs(cur, key[idx], *flag, has)
		if dist < 0 {
			return -1
		}
		has[key[idx]] = true
		return dfs(pos, key, has, idx+1, flag, ans+dist)
	}
	flag := new(int)

	for _, k := range res {
		tmp := dfs(start, k, make(map[byte]bool), 0, flag, 0)
		if tmp > 0 && (best < 0 || best > tmp) {
			best = tmp
		}
	}
	return best
}
