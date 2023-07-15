package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, Q)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

const H = 18

func solve(n int, Q [][]int) int {
	// nodes that in comp u

	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
		for j := 0; j < H; j++ {
			P[i][j] = i
		}
	}

	depth := make([]int, n)

	g := NewGraph(n, 2*len(Q))

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}
	}

	dia := make([][]int, n)
	res := make([]int, n)
	var sum int
	sz := make([]int, n)

	for i := 0; i < n; i++ {
		dia[i] = []int{i, i}
		res[i] = 0
		sz[i]++
	}

	lca := func(u int, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if depth[u]-(1<<i) >= depth[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}

		return P[u][0]
	}

	distance := func(u int, v int) int {
		p := lca(u, v)
		return depth[u] + depth[v] - 2*depth[p]
	}

	merge := func(u int, v int) {
		u--
		v--
		// merge tree v into u
		ru := P[u][H-1]
		rv := P[v][H-1]

		if sz[ru] < sz[rv] {
			u, v = v, u
			ru, rv = rv, ru
		}

		sum -= res[ru]
		sum -= res[rv]
		sz[ru] += sz[rv]

		depth[v] = depth[u] + 1
		dfs(u, v)

		arr := []int{dia[ru][0], dia[ru][1], dia[rv][0], dia[rv][1]}

		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				tmp := distance(arr[i], arr[j])
				if res[ru] < tmp {
					res[ru] = tmp
					dia[ru][0] = arr[i]
					dia[ru][1] = arr[j]
				}
			}
		}

		g.AddEdge(u, v)
		g.AddEdge(v, u)

		sum += res[ru]
	}
	var score int
	for _, cur := range Q {
		u, v := cur[0], cur[1]
		u ^= score
		v ^= score
		merge(u, v)
		score ^= sum
	}

	return score
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
