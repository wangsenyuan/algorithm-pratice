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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n-1)
		A := readNNums(reader, n-1)
		res := solve(n, P, A)
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

const N_INF = -(1 << 59)

func solve(n int, P []int, A []int) int64 {

	pp := make([]int, n)
	copy(pp[1:], P)

	g := NewGraph(n, n+10)

	for i := 1; i < n; i++ {
		pp[i]--
		g.AddEdge(pp[i], i)
	}

	D := make([]int, n)
	var mx int
	var dfs func(u int)
	dfs = func(u int) {
		mx = max(mx, D[u])
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			D[v] = D[u] + 1
			dfs(v)
		}
	}
	dfs(0)
	mx++
	C := make([]int, mx)
	for i := 0; i < n; i++ {
		C[D[i]]++
	}
	L := make([][]Pair, mx)
	for i := 0; i < mx; i++ {
		L[i] = make([]Pair, 0, C[i])
	}
	for i := 1; i < n; i++ {
		L[D[i]] = append(L[D[i]], Pair{A[i-1], i})
	}

	for i := 0; i < mx; i++ {
		sort.Sort(Pairs(L[i]))
	}
	dp := make([]int64, n)

	for i := 1; i < mx; i++ {
		cur := L[i]
		// if place blue at j'th node in this level, find another best red
		// because all nodes before j < A[j], so the best is A[j] - A[k] + dp[P[k]]
		// let best is dp[P[k]] - A[k]
		var best int64 = N_INF
		for _, p := range cur {
			j := p.second
			dp[j] = max2(dp[j], best+int64(A[j-1]))
			best = max2(best, dp[pp[j]]-int64(A[j-1]))
		}
		// A[k] - A[j] + dp[pp[k]]
		best = N_INF
		for k := len(cur) - 1; k >= 0; k-- {
			p := cur[k]
			j := p.second
			dp[j] = max2(dp[j], best-int64(A[j-1]))
			best = max2(best, int64(A[j-1])+dp[pp[j]])
		}
		// a0 or an is blue
		a0 := cur[0].first
		an := cur[len(cur)-1].first
		for _, p := range cur {
			j := p.second
			//if make j as red, then it's parent has to be red
			dp[j] = max2(dp[j], int64(p.first-a0)+dp[pp[j]])
			dp[j] = max2(dp[j], int64(an-p.first)+dp[pp[j]])
		}
	}

	var res int64

	for _, p := range L[mx-1] {
		res = max2(res, dp[p.second])
	}

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
