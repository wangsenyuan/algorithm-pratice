package lcp38

const INF = 20010

func guardCastle(grid []string) int {
	G := make([][]byte, 2)
	for i := 0; i < 2; i++ {
		G[i] = []byte(grid[i])
		for j := 0; j < len(G[i]); j++ {
			if G[i][j] == 'P' {
				G[i][j] = 'S'
			}
		}
	}
	a := dp(G)
	for i := 0; i < 2; i++ {
		G[i] = []byte(grid[i])
		for j := 0; j < len(G[i]); j++ {
			if G[i][j] == 'P' {
				G[i][j] = 'C'
			}
		}
	}
	b := dp(G)
	a = min(a, b)
	if a >= INF {
		return -1
	}
	return a
}

var rep = map[byte]int{'.': 0, 'C': 1, 'S': 2, '#': 3}

func dp(grid [][]byte) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' || grid[i][j] == '#' {
				continue
			}
			if i > 0 && rep[grid[i-1][j]]+rep[grid[i][j]] == 3 {
				return INF
			}
			if j > 0 && rep[grid[i][j-1]]+rep[grid[i][j]] == 3 {
				return INF
			}
		}
	}

	f := make([]int, 16)
	for i := 0; i < 16; i++ {
		f[i] = INF
	}
	f[0] = 0
	g := make([]int, 16)

	update := func(v int, s1, s2, t1, t2 int, f []int) {
		if s1 == 1 || s1 == 2 {
			if s1+t1 == 3 {
				// monster and castle side by side
				return
			}
			if t1 == 0 {
				t1 = s1
			}
		}
		if s2 == 1 || s2 == 2 {
			if s2+t2 == 3 {
				return
			}
			if t2 == 0 {
				t2 = s2
			}
		}

		if t1+t2 == 3 && (t1 == 1 || t2 == 1) {
			return
		}

		if (t1 == 1 || t1 == 2) && t2 == 0 {
			t2 = t1
		}
		if (t2 == 1 || t2 == 2) && t1 == 0 {
			t1 = t2
		}
		f[t1*4+t2] = min(f[t1*4+t2], v)
	}

	for i := 0; i < len(grid[0]); i++ {
		t1 := rep[grid[0][i]]
		t2 := rep[grid[1][i]]
		for j := 0; j < 16; j++ {
			g[j] = INF
		}
		for x := 0; x < 16; x++ {
			if f[x] >= INF {
				continue
			}
			s1, s2 := x/4, x%4
			update(f[x], s1, s2, t1, t2, g)
			if grid[0][i] == '.' {
				// try put a stone
				update(f[x]+1, s1, s2, 3, t2, g)
			}
			if grid[1][i] == '.' {
				update(f[x]+1, s1, s2, t1, 3, g)
			}
			if grid[0][i] == '.' && grid[1][i] == '.' {
				update(f[x]+2, s1, s2, 3, 3, g)
			}
		}
		copy(f, g)
	}
	res := INF
	for x := 0; x < 16; x++ {
		res = min(res, f[x])
	}
	return res
}

var dd = []int{-1, 0, 1, 0, -1}

func guardCastle1(grid []string) int {
	n := len(grid[0])
	// S 是monster的位置，C是城堡的位置，P是传送点, #是石头
	var sx, sy int
	N := 4*n + 2
	g := NewGraph(N, 2*n*10)

	connect := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, 0)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				continue
			}
			id := i*n + j
			if grid[i][j] == '.' {
				connect(id*2, id*2+1, 1)
			} else {
				connect(id*2, id*2+1, INF)
			}
			if grid[i][j] == 'C' {
				sx = i
				sy = j
			}
			if grid[i][j] == 'P' {
				connect(id*2+1, 4*n, INF)
				connect(4*n, id*2, INF)
			}
			if grid[i][j] == 'S' {
				// 从monster到汇点
				connect(id*2+1, 4*n+1, INF)
			}
			for k := 0; k < 4; k++ {
				u, v := i+dd[k], j+dd[k+1]
				if u >= 0 && u < 2 && v >= 0 && v < n {
					if grid[u][v] != '#' {
						ii := u*n + v
						connect(id*2+1, ii*2, INF)
					}
				}
			}
		}
	}
	level := make([]int, N)
	que := make([]int, N)
	bfs := func(s int) {
		for i := 0; i < N; i++ {
			level[i] = -1
		}
		level[s] = 0
		var front, end int
		que[end] = s
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if level[v] < 0 && g.val[i] > 0 {
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
	}
	pos := make([]int, N)

	var dinic func(u int, f int) int

	dinic = func(u int, f int) int {
		if u == 4*n+1 {
			return f
		}
		for ; pos[u] > 0; pos[u] = g.next[pos[u]] {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.val[i] > 0 {
				res := dinic(v, min(f, g.val[i]))
				if res > 0 {
					g.val[i] -= res
					g.val[i^1] += res
					return res
				}
			}
		}
		return 0
	}
	start := (sx*n + sy) * 2
	var flow int
	for flow < INF {
		bfs(start)
		if level[N-1] < 0 {
			break
		}
		for i := 0; i < N; i++ {
			pos[i] = g.nodes[i]
		}
		for flow < INF {
			tmp := dinic(start, INF)
			if tmp == 0 {
				break
			}
			flow += tmp
		}
	}
	if flow < INF {
		return flow
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	g.val = make([]int, e+10)
	// start from 2, reverse of edge i is i ^ 1
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
