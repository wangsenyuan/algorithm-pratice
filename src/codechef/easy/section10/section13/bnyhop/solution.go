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
		P := readNNums(reader, n-1)
		A := readNNums(reader, n)
		res := solve(n, A, P)
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

const M = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= M {
		a -= M
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, M-b)
}

func solve(n int, A []int, P []int) int {
	g := NewGraph(n, n-1)

	for i := 1; i < n; i++ {
		g.AddEdge(P[i-1]-1, i)
	}

	var p int
	st := make([]int, n)
	en := make([]int, n)

	var eularTour func(u int)
	eularTour = func(u int) {
		p++
		st[u] = p - 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			eularTour(g.to[i])
		}
		p++
		en[u] = p - 1
	}

	eularTour(0)

	pp := make([]Pair, n)

	for i := 0; i < n; i++ {
		pp[i] = Pair{A[i], i}
	}

	sort.Sort(Pairs(pp))

	dpTree := NewBit(2 * n)
	dpSubTree := NewBit(2 * n)

	var ans int

	for i := 0; i < n; i++ {
		ix := pp[i].second
		// 获取子树中比A[ix]大的节点的数量，使用的是点更新，获取区间值的方式
		cur := modAdd(1, dpTree.GetRange(st[ix], en[ix]))
		// 获取root到ix路径上所有节点对ix的影响，使用的是range update， point get 的方式；
		cur = modAdd(cur, dpSubTree.Get(st[ix]))
		ans = modAdd(ans, cur)

		dpSubTree.Add(st[ix], cur)
		dpSubTree.Add(en[ix], M-cur)
		dpTree.Add(st[ix], cur)
	}

	ans = modSub(ans, n)

	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
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
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Bit struct {
	arr []int
	n   int
}

func NewBit(n int) *Bit {
	bit := new(Bit)
	bit.arr = make([]int, n+1)
	bit.n = n
	return bit
}

func (b *Bit) Add(p int, v int) {
	p++
	for p <= b.n {
		b.arr[p] = modAdd(b.arr[p], v)
		p += p & -p
	}
}

func (b *Bit) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res = modAdd(res, b.arr[p])
		p -= p & -p
	}
	return res
}

func (b *Bit) GetRange(l, r int) int {
	if l > r {
		return 0
	}
	res := b.Get(r)
	res = modSub(res, b.Get(l-1))
	return res
}
