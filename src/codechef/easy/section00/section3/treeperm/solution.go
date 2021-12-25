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
		n, S := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, A, B, S, E)
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

const MOD = 1000000007
const MAX_X = 1000010

var freq [MAX_X]int

func solve(n int, A, B []int, S int, E [][]int) int {
	g := NewGraph(n+1, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u, v, 1)
		g.AddEdge(v, u, 1)
	}

	deg := make([]int, n+1)
	parent := make([]int, n+1)
	color := make([]int, n+1)

	for i := 1; i <= n; i++ {
		deg[i] = -1
		color[i] = -1
	}

	var dfs func(u int)

	dfs = func(u int) {
		deg[u] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if deg[v] < 0 {
				parent[v] = u
				deg[u]++
				dfs(v)
			}
		}
	}

	dfs(1)

	leaves := make([]int, n)
	var end int
	for i := 1; i <= n; i++ {
		if deg[i] == 0 {
			leaves[end] = i
			end++
		}
	}

	var nonZero int

	update := func(x int, delta int) {
		if freq[x] == 0 {
			nonZero++
		}
		freq[x] += delta
		if freq[x] == 0 {
			nonZero--
		}
	}

	var front int
	for front < end {
		u := leaves[front]
		id := u
		front++
		var good bool
		for {
			color[u] = id
			update(A[u-1], 1)
			update(B[u-1], -1)
			deg[parent[u]]--
			if nonZero == 0 {
				good = true
				if deg[parent[u]] == 0 && color[parent[u]] < 0 {
					leaves[end] = parent[u]
					end++
				}
				break
			}
			if color[parent[u]] < 0 {
				u = parent[u]
			} else {
				break
			}
		}

		if !good {
			return 0
		}
	}

	if S == 1 {
		return 1
	}

	F := make([][]int, n+1)
	roots := make([]int, 0, n)

	for i := 0; i <= n; i++ {
		F[i] = make([]int, 0, 3)
	}

	for i := 1; i <= n; i++ {
		if color[i] != color[parent[i]] {
			if i > 1 && color[parent[i]] == parent[i] {
				F[parent[i]] = append(F[parent[i]], color[i])
			} else {
				roots = append(roots, color[i])
			}
		}
	}

	var dp func(u int) int64

	dp = func(u int) int64 {
		var ret = 1 + int64(len(F[u]))
		for _, v := range F[u] {
			ret *= dp(v)
			ret %= MOD
		}
		return ret
	}

	var ans int64 = 1

	for _, u := range roots {
		ans = (ans * dp(u)) % MOD
	}

	for i := 0; i < n; i++ {
		freq[A[i]] = 0
		freq[B[i]] = 0
	}

	return int(ans)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
