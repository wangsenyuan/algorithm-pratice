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
		n, m, c := readThreeNums(reader)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, edges, c)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 62

type Pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int, c int) int {
	// c is useless in this problem
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	// 需要将那些strong component缩放成一个点
	stack := make([]int, n)
	var top int
	low := make([]int, n)
	dis := make([]int, n)
	vis := make([]bool, n)

	belong := make([]int, n)
	var comp []int

	var time int

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		dis[u] = time
		low[u] = time
		time++
		vis[u] = true

		stack[top] = u
		top++

		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if !vis[v] {
					sz += dfs(u, v)
					low[u] = min(low[u], low[v])
				} else {
					low[u] = min(low[u], dis[v])
				}
			}
		}

		if low[u] == dis[u] {
			// 这部分是一个强连通分量
			id := len(comp)
			var cnt int
			for top > 0 {
				v := stack[top-1]
				top--
				belong[v] = id
				cnt++
				if u == v {
					break
				}
			}
			comp = append(comp, cnt)
		}

		return sz
	}
	var trs []Pair

	for i := 0; i < n; i++ {
		if !vis[i] {
			sz := dfs(-1, i)
			trs = append(trs, Pair{sz, i})
		}
	}
	m := len(comp)

	if m == 1 {
		// can't split
		return -1
	}

	// g2 is a forest for those trees
	g2 := NewGraph(m, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u = belong[u]
		v = belong[v]
		if u != v {
			g2.AddEdge(u, v)
			g2.AddEdge(v, u)
		}
	}

	best := inf

	var dfs2 func(p int, u int, k int, bs *BitSet) int

	dfs2 = func(p int, u int, k int, bs *BitSet) int {
		sz := comp[u]
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if p != v {
				sz += dfs2(u, v, k, bs)
			}
		}

		for x := sz; x <= n; x++ {
			y := n - x
			if y < k-sz {
				break
			}
			a := x - sz
			b := y - (k - sz)
			if bs.IsSet(a) && bs.IsSet(b) {
				best = min(best, x*x+y*y)
			}
		}

		return sz
	}

	bs := NewBitSet(n + 1)
	k := len(trs)

	for i, tr := range trs {
		bs.Reset()
		bs.Set(0)
		for j := 0; j < k; j++ {
			if i == j {
				continue
			}
			tmp := bs.Copy()
			tmp.LeftShift(trs[j].first)
			bs.Union(tmp)
		}

		dfs2(-1, belong[tr.second], tr.first, bs)
	}
	// a = k-1
	return best + c*(k-1)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) Count() int {
	var res int
	for i := 0; i < bs.sz; i++ {
		res += countDigit(bs.arr[i])
	}
	return res
}

func countDigit(num uint64) int {
	var res int
	if (num>>63)&1 == 1 {
		res++
	}
	tmp := int64(num & ((1 << 63) - 1))

	for tmp > 0 {
		res++
		tmp -= tmp & -tmp
	}
	return res
}

func (bs *BitSet) LeftShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := 0; u+i < bs.sz; u++ {
			bs.arr[u] = bs.arr[u+i]
		}
	} else {
		for u := 0; u+i < bs.sz; u++ {
			v := u + i
			bs.arr[u] = bs.arr[v] << uint64(j)
			if v+1 < bs.sz {
				bs.arr[u] |= bs.arr[v+1] >> uint64(64-j)
			}
		}
	}
	for u := bs.sz - i; u < bs.sz; u++ {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) RightShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := bs.sz - 1; u-i >= 0; u-- {
			bs.arr[u] = bs.arr[u-i]
		}
	} else {
		for u := bs.sz - 1; u-i >= 0; u-- {
			v := u - i
			bs.arr[u] = bs.arr[v] >> uint64(j)
			if v-1 >= 0 {
				bs.arr[u] |= bs.arr[v-1] << uint64(64-j)
			}
		}
	}
	for u := i - 1; u >= 0; u-- {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) Copy() *BitSet {
	res := NewBitSet(len(bs.arr) * 64)
	copy(res.arr, bs.arr)
	return res
}

func (this *BitSet) Union(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] |= that.arr[i]
	}
}

func (this *BitSet) Reset() {
	clear(this.arr)
}
