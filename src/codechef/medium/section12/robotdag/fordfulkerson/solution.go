package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, m, k, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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
func solve(n, m, k int, E [][]int) int {
	//G := NewGraph(n, m+10)
	//R := NewGraph(n, m+10)
	// maxT == 110
	maxT := n + 10

	var tt int

	ind := make([][]int, n)
	for i := 0; i < n; i++ {
		ind[i] = make([]int, maxT+1)
		for j := 0; j <= maxT; j++ {
			ind[i][j] = tt
			tt++
		}
	}
	// tt = 100 * 110 ~ 10000
	// so we can't have a grid
	G := NewGraph(n*(maxT+1)+10, 2*(maxT+1)*m+10)
	for i := 0; i < maxT; i++ {
		for j := 0; j < m; j++ {
			u, v := E[j][0], E[j][1]
			u--
			v--
			G.AddEdge(ind[u][i], ind[v][i+1], 1)
			G.AddEdge(ind[v][i+1], ind[u][i], 0)
		}
	}

	que := make([]int, tt)
	vis := make([]int, tt)
	time := 1
	pick := make([]int, tt)
	parent := make([]int, tt)
	bfs := func(s, t int) bool {
		var front, end int
		que[end] = s
		end++
		vis[s] = time

		for front < end {
			u := que[front]
			front++

			for i := G.nodes[u]; i > 0; i = G.next[i] {
				v := G.to[i]
				if vis[v] != time && G.edge[i] > 0 {
					parent[v] = u
					pick[v] = i
					que[end] = v
					end++
					vis[v] = time
				}
			}
		}
		res := vis[t] == time
		time++
		return res
	}

	foldfulkerson := func(s, t int) int {
		var maxFlow int
		for bfs(s, t) {
			maxFlow++
			for v := t; v != s; v = parent[v] {
				G.edge[pick[v]]--
				G.edge[pick[v]^1]++
			}
		}
		return maxFlow
	}

	var ans int

	for i := 0; i < maxT; i++ {
		ans += foldfulkerson(0, ind[n-1][i])
		if ans >= k {
			return i
		}
	}
	return -1
}

//
type Graph struct {
	nodes []int
	next  []int
	to    []int
	edge  []int
	cur   int
}

func NewGraph(nodeCount int, edgeCap int) *Graph {
	nodes := make([]int, nodeCount)
	next := make([]int, edgeCap)
	to := make([]int, edgeCap)
	edge := make([]int, edgeCap)
	// edge index from 1, then pair forward & backward by even & event + 1
	return &Graph{nodes, next, to, edge, 1}
}

func (g *Graph) AddEdge(u, v, c int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.edge[g.cur] = c
}
