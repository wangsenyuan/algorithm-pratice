package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	detachments := make([][]int, n)

	for i := 0; i < n; i++ {
		detachments[i] = readNNums(reader, 2)
	}
	fmt.Println(solve(n, detachments))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const INF = 2000000001

func solve(n int, detachments [][]int) int {

	sameX := make(map[int][]int)
	sameY := make(map[int][]int)

	for i := 0; i < n; i++ {
		cur := detachments[i]
		x, y := cur[0], cur[1]
		if _, found := sameX[x]; !found {
			sameX[x] = make([]int, 0, 10)
		}
		sameX[x] = append(sameX[x], i)

		if _, found := sameY[y]; !found {
			sameY[y] = make([]int, 0, 10)
		}
		sameY[y] = append(sameY[y], i)
	}

	var left, right = 0, INF

	g := NewGraph(n)
	segs := make([]Pair, n*n/2)
	segx := make([]Pair, n*n/2)
	segy := make([]Pair, n*n/2)

	for left < right {
		mid := (left + right) / 2

		g.Reset()
		segs = segs[:0]
		segx = segx[:0]
		segy = segy[:0]

		for _, v := range sameX {
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					if abs(detachments[v[i]][1]-detachments[v[j]][1]) <= mid {
						g.Connect(v[i], v[j])
					}
				}
			}
		}

		for _, v := range sameY {
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					if abs(detachments[v[i]][0]-detachments[v[j]][0]) <= mid {
						g.Connect(v[i], v[j])
					}
				}
			}
		}

		cnt := 1

		for i := 0; i < n; i++ {
			if g.GetColor(i) == 0 {
				g.Visit(i, cnt)
				cnt++
			}
		}
		cnt--

		if cnt > 4 {
			// bad
			left = mid + 1
			continue
		}

		if cnt == 1 {
			// good
			right = mid
			continue
		}
		var ok bool

		if cnt == 2 {
			for i := 0; i < n && !ok; i++ {
				for j := i + 1; j < n && !ok; j++ {
					if g.GetColor(i) == g.GetColor(j) {
						continue
					}
					if detachments[i][0] == detachments[j][0] {
						ok = abs(detachments[i][1]-detachments[j][1]) <= 2*mid
					} else if detachments[i][1] == detachments[j][1] {
						ok = abs(detachments[i][0]-detachments[j][0]) <= 2*mid
					} else {
						ok = abs(detachments[i][0]-detachments[j][0]) <= mid && abs(detachments[i][1]-detachments[j][1]) <= mid
					}
				}
			}
		} else if cnt == 3 {
			for _, v := range sameX {
				for i := 1; i < len(v); i++ {
					if g.GetColor(v[i]) != g.GetColor(v[i-1]) {
						segs = append(segs, Pair{v[i], v[i-1]})
					}
				}
			}
			for _, v := range sameY {
				for i := 1; i < len(v); i++ {
					if g.GetColor(v[i]) != g.GetColor(v[i-1]) {
						segs = append(segs, Pair{v[i], v[i-1]})
					}
				}
			}

			for k := 0; k < n; k++ {
				for _, p := range segs {
					i, j := p.fi, p.se
					if g.GetColor(i) == g.GetColor(k) || g.GetColor(j) == g.GetColor(k) {
						continue
					}
					if detachments[i][0] == detachments[j][0] {
						// same x
						if min(detachments[i][1], detachments[j][1]) >= detachments[k][1] || max(detachments[i][1], detachments[j][1]) <= detachments[k][1] {
							continue
						}
						if abs(detachments[i][0]-detachments[k][0]) > mid {
							continue
						}
						if abs(detachments[i][1]-detachments[k][1]) <= mid && abs(detachments[j][1]-detachments[k][1]) <= mid {
							ok = true
						}
					} else {
						// same y
						if min(detachments[i][0], detachments[j][0]) >= detachments[k][0] || max(detachments[i][0], detachments[j][0]) <= detachments[k][0] {
							continue
						}
						if abs(detachments[i][1]-detachments[k][1]) > mid {
							continue
						}
						if abs(detachments[i][0]-detachments[k][0]) <= mid && abs(detachments[j][0]-detachments[k][0]) <= mid {
							ok = true
						}
					}
				}
			}

		} else {
			// cnt == 4
			for _, v := range sameX {
				for i := 1; i < len(v); i++ {
					if g.GetColor(v[i]) != g.GetColor(v[i-1]) {
						segx = append(segx, Pair{v[i], v[i-1]})
					}
				}
			}
			for _, v := range sameY {
				for i := 1; i < len(v); i++ {
					if g.GetColor(v[i]) != g.GetColor(v[i-1]) {
						segy = append(segy, Pair{v[i], v[i-1]})
					}
				}
			}

			for _, a := range segx {
				for _, b := range segy {
					i, j := a.fi, a.se
					k, p := b.fi, b.se

					if g.GetColor(i) == g.GetColor(k) || g.GetColor(i) == g.GetColor(p) || g.GetColor(j) == g.GetColor(k) || g.GetColor(j) == g.GetColor(p) {
						continue
					}
					if min(detachments[i][1], detachments[j][1]) >= detachments[k][1] {
						continue
					}
					if max(detachments[i][1], detachments[j][1]) <= detachments[k][1] {
						continue
					}
					if min(detachments[k][0], detachments[p][0]) >= detachments[i][0] {
						continue
					}
					if max(detachments[k][0], detachments[p][0]) <= detachments[i][0] {
						continue
					}
					x0, y0 := detachments[i][0], detachments[k][1]
					if abs(detachments[i][1]-y0) <= mid && abs(detachments[j][1]-y0) <= mid && abs(detachments[k][0]-x0) <= mid && abs(detachments[p][0]-x0) <= mid {
						ok = true
					}
				}
			}
		}

		if ok {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == INF {
		return -1
	}
	return right
}

type Pair struct {
	fi, se int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	n   int
	adj [][]int
	vis []int
}

func NewGraph(n int) Graph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, (n+9)/10)
	}
	vis := make([]int, n)
	return Graph{n, adj, vis}
}

func (g *Graph) Reset() {
	for i := 0; i < g.n; i++ {
		g.adj[i] = g.adj[i][:0]
		g.vis[i] = 0
	}
}

func (g *Graph) Connect(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) Visit(s int, color int) {
	var dfs func(u int)

	dfs = func(u int) {
		g.vis[u] = color
		for _, v := range g.adj[u] {
			if g.vis[v] != color {
				dfs(v)
			}
		}
	}

	dfs(s)
}

func (g Graph) GetColor(u int) int {
	return g.vis[u]
}
