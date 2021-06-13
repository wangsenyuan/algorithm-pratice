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
		n, k := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			s, _ := reader.ReadBytes('\n')
			var cnt int
			pos := readInt(s, 0, &cnt) + 1
			Q[i] = make([]int, cnt)
			for j := 0; j < cnt; j++ {
				pos = readInt(s, pos, &Q[i][j]) + 1
			}
		}
		res := solve(n, E, Q)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d", len(res[i])))
			for j := 0; j < len(res[i]); j++ {
				buf.WriteString(fmt.Sprintf(" %d", res[i][j]))
			}
			buf.WriteByte('\n')
		}
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

const M_H = 20

func solve(N int, E [][]int, Q [][]int) [][]int {
	g := NewGraph(N, len(E)*2)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, N)
	P := make([][]int, N)
	for i := 0; i < N; i++ {
		P[i] = make([]int, M_H)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < M_H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		// D[u] >= D[v]
		for i := M_H - 1; i >= 0; i-- {
			if D[u]-(1<<uint(i)) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := M_H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	ans := make([][]int, len(Q))

	for i, cur := range Q {
		first := -1
		for j := 0; j < len(cur); j++ {
			if first < 0 || D[first] < D[cur[j]-1] {
				first = cur[j] - 1
			}
		}

		if len(cur) == 1 {
			ans[i] = []int{first + 1}
			continue
		}
		var di int
		second := -1
		for j := 0; j < len(cur); j++ {
			p := lca(first, cur[j]-1)
			tmp := D[first] + D[cur[j]-1] - 2*D[p]
			if tmp > di {
				second = cur[j] - 1
				di = tmp
			}
		}

		if D[first] > D[second] {
			second = first
		}
		half := di / 2
		for j := M_H - 1; j >= 0; j-- {
			if half >= 1<<uint(j) {
				half -= 1 << uint(j)
				second = P[second][j]
			}
		}
		if di%2 == 0 {
			ans[i] = []int{second + 1}
		} else {
			ans[i] = []int{min(second+1, P[second][0]+1), max(second+1, P[second][0]+1)}
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
