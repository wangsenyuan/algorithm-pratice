package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		m := readNum(reader)
		P := readNNums(reader, m)
		res := solve(n, E, P)
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

func solve(n int, E [][]int, P []int) int {
	// k = prod_of(P)
	// 如果 len(P) < n - 1, 那么一种可能的安排是把P[i]分布到边上
	// 这时候的1的个数是最小的
	// 如果 len(P) >= n,
	// 这时候1的个数是0
	// max(f(i, j)) for i in [0...n), j in [0...n)
	// 已知 a * b >= a + b,
	// 所以尽可能的数出现的次数越多越好
	// 那就需要计算每天边的贡献度了，
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	contrib := make([]int64, 0, n-1)

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sz += dfs(u, v)
			}
		}
		// edge into u
		if p >= 0 {
			contrib = append(contrib, int64(n-sz)*int64(sz))
		}

		return sz
	}

	dfs(-1, 0)

	sort.Slice(contrib, func(i, j int) bool {
		return contrib[i] > contrib[j]
	})
	sort.Ints(P)
	var res int

	it := len(P) - 1

	if len(P) >= n-1 {
		cur := 1
		for i := len(P) - 1; i >= n-2; i-- {
			cur = modMul(cur, P[i])
		}
		it = n - 2
		P[it] = cur
	}
	for i := 0; i < len(contrib); i++ {
		x := int(contrib[i] % MOD)
		if it < 0 {
			res = modAdd(res, x)
		} else {
			res = modAdd(res, modMul(P[it], x))
			it--
		}
	}

	return res
}

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
