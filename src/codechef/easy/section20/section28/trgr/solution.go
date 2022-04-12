package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	X := make([][]int, n)
	for i := 0; i < n; i++ {
		X[i] = readNNums(reader, 2)
	}

	res := solve(n, E, X)

	var buf bytes.Buffer

	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
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

func solve(n int, E [][]int, X [][]int) []int {
	//通过union-find把图连起来没有问题
	// 但如何把在离开节点的时候，将图给分开呢?
	// 目前的实现有问题
	g := NewGraph(n, 2*(n-1))

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	set := NewSet(n)

	ans := make([]int, n)

	roots := make([][]int, n)
	for i := 0; i < n; i++ {
		roots[i] = make([]int, 0, 1)
		roots[i] = append(roots[i], i)
	}

	var dfs func(p, u int)

	cc := n

	dfs = func(p, u int) {
		info := X[u]
		a, b := info[0]-1, info[1]-1
		pa := set.Find(a)
		pb := set.Find(b)
		union := pa != pb
		var buf []int
		if union {
			if set.cnt[pa] < set.cnt[pb] {
				pa, pb = pb, pa
			}
			buf = make([]int, set.cnt[pb])
			copy(buf, roots[pb])
			roots[pa] = append(roots[pa], roots[pb]...)
			set.cnt[pa] += set.cnt[pb]
			set.arr[pb] = pa
			cc--
		}

		ans[u] = cc

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}

		if union {
			for i := 0; i < len(buf); i++ {
				set.arr[buf[i]] = pb
			}
			roots[pb] = buf
			set.cnt[pa] -= set.cnt[pb]
			roots[pa] = roots[pa][:set.cnt[pa]]
			cc++
		}
	}

	dfs(0, 0)

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

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	set := new(Set)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i]++
	}
	return set
}

func (set *Set) Find(u int) int {
	p := set.arr[u]
	if p != u {
		set.arr[u] = set.Find(p)
	}
	return set.arr[u]
}

func (set *Set) Union(a, b int) {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa != pb {
		if set.cnt[pa] < set.cnt[pb] {
			pa, pb = pb, pa
		}
		set.cnt[pa] += set.cnt[pb]
		set.arr[pb] = pa
	}
}
