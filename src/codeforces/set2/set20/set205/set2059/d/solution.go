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
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n, s1, s2 := readThreeNums(reader)
	m1 := readNum(reader)
	edges1 := make([][]int, m1)
	for i := range m1 {
		edges1[i] = readNNums(reader, 2)
	}
	m2 := readNum(reader)
	edges2 := make([][]int, m2)
	for i := range m2 {
		edges2[i] = readNNums(reader, 2)
	}
	return solve(n, edges1, edges2, s1, s2)
}

const inf = 1 << 60

func solve(n int, edges1 [][]int, edges2 [][]int, s1 int, s2 int) int {
	// 假设下一个点是 (r2, c2) 如果 r2 > c2, 就把它放在左下的一个ds中
	// 如果 r2 < c2, 就把它放在右上的一个ds中。
	// 然后分别从两个结构里面去查最小值
	// 循环两种边，这样子，可以构造出(a, b) -> (c, d) 的边
	// 这个复杂性是 m1 * m2
	// 然后从 (s1, s2) bfs
	// (a, b) -> (c, d) 的边的cost = abs(d - c)
	// n * n 个点 + m1 * m2 条边
	// 可以搞
	m1 := len(edges1)
	m2 := len(edges2)
	g := NewGraph(n*n, m1*m2*4)

	add := func(r int, c int, u int, v int) {
		g.AddEdge(r*n+c, u*n+v, abs(u-v))
		g.AddEdge(u*n+v, r*n+c, abs(r-c))
	}

	for _, e1 := range edges1 {
		u, v := e1[0]-1, e1[1]-1
		for _, e2 := range edges2 {
			r, c := e2[0]-1, e2[1]-1
			add(u, r, v, c)
			add(u, c, v, r)
		}
	}

	items := make([]*Item, n*n)
	pq := make(PriorityQueue, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			it := new(Item)
			it.id = i*n + j
			it.priority = inf
			it.index = i*n + j
			items[it.id] = it
			pq[it.id] = it
		}
	}

	s1--
	s2--

	items[s1*n+s2].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority == inf {
			break
		}
		r := it.id / n
		c := it.id % n

		for i := g.nodes[it.id]; i > 0; i = g.next[i] {
			x, y := g.to[i]/n, g.to[i]%n
			if r == c && x == y {
				return it.priority
				// 没必要继续到(x, y)
			}
			if items[x*n+y].index >= 0 && items[x*n+y].priority > it.priority+g.val[i] {
				pq.update(items[x*n+y], it.priority+g.val[i])
			}
		}
	}

	return -1
}

func abs(num int) int {
	return max(num, -num)
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
