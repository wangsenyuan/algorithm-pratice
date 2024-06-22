package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, k, edges)
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
func solve(n int, k int, edges [][]int) int {
	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int, expect int) int

	var cnt int

	dfs = func(p int, u int, expect int) int {
		var sz int = 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v, expect)
				if tmp >= expect {
					cnt++
					continue
				}
				sz += tmp
			}
		}
		return sz
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		cnt = 0
		sz := dfs(0, 1, mid)
		if cnt == k && sz >= mid || cnt > k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func solve1(n int, k int, edges [][]int) int {
	g := NewGraph(n+1, 2*n)
	deg := make([]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	dg := make([]int, n+1)
	items := make([]*Item, n+1)
	for i := 1; i <= n; i++ {
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = 1
	}

	check := func(expect int) bool {
		pq := make(PQ, 0, n)
		for i := 1; i <= n; i++ {
			items[i].priority = 1
			if deg[i] == 1 {
				heap.Push(&pq, items[i])
			}
		}

		copy(dg, deg)
		var cnt int

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			add := it.priority
			if it.priority >= expect {
				// 它的上面截断
				cnt++
				add = 0
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				dg[v]--
				items[v].priority += add
				if dg[v] == 1 {
					heap.Push(&pq, items[v])
				}
			}
		}

		return cnt > k
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2

		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(item interface{}) {
	cur := item.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	arr := *pq
	n := len(arr)
	res := arr[n-1]
	*pq = arr[:n-1]
	res.index = -1
	return res
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
