package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	N, K := readTwoNums(reader)
	B := readNum(reader)
	E := make([][]int, N-1)
	for i := 0; i < N-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	F := make([]int, N)
	for i := 0; i < N; i++ {
		F[i] = readNum(reader)
	}
	q := readNum(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(N, K, B, F, E, Q)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const H = 20

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve(N, K, B int, F []int, E [][]int, Q [][]int) []int {
	B--

	for i := 0; i < N; i++ {
		F[i]--
	}

	g := NewGraph(N, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sets := make([][]int, N)

	for i := 0; i < N; i++ {
		sets[i] = make([]int, K)
		for j := 0; j < K; j++ {
			sets[i][j] = N + 1
		}
	}

	P := make([][]int, N)

	for i := 0; i < N; i++ {
		P[i] = make([]int, H)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			x := P[u][i-1]
			P[u][i] = P[x][i-1]
		}

		sets[u][F[u]] = u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				for j := 0; j < K; j++ {
					sets[u][j] = min(sets[u][j], sets[v][j])
				}
			}
		}
	}

	dfs(B, B)

	find := func(a int, p int) int {
		if sets[a][p] < N {
			return sets[a][p]
		}

		if sets[B][p] >= N {
			return N
		}

		for i := H - 1; i >= 0; i-- {
			if sets[P[a][i]][p] >= N {
				a = P[a][i]
			}
		}

		return sets[P[a][0]][p]
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		a, p := cur[0], cur[1]
		a--
		p--
		r := find(a, p)
		if r >= N {
			r = -1
		} else {
			r++
		}
		ans[i] = r
	}
	return ans
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
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
