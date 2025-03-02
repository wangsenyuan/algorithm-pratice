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

func process(reader *bufio.Reader) int {
	n, m, h := readThreeNums(reader)
	a := readNNums(reader, h)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, a, edges)
}

const inf = 1 << 60

func solve(n int, a []int, edges [][]int) int {
	g := NewGraph(n, 2*len(edges)+1)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	horse := make([]bool, n)
	for _, i := range a {
		horse[i-1] = true
	}

	bfs := func(s int) []int {
		items := make([][]*Item, n)
		pq := make(PQ, n*2)
		for i := 0; i < n; i++ {
			items[i] = make([]*Item, 2)
			for j := 0; j < 2; j++ {
				it := new(Item)
				it.id = i
				it.ride = j
				it.priority = inf
				it.index = i*2 + j
				items[i][j] = it
				pq[it.index] = it
			}
		}

		items[s][0].priority = 0
		heap.Init(&pq)
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			r := it.ride
			if r == 1 || horse[u] {
				// 接下来都可以骑马了
				r = 1
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v][r].priority > it.priority+g.val[i]/(r+1) {
					pq.update(items[v][r], it.priority+g.val[i]/(r+1))
				}
			}
		}
		res := make([]int, n)
		for i := range n {
			res[i] = min(items[i][0].priority, items[i][1].priority)
		}
		return res
	}

	d1 := bfs(0)

	if d1[n-1] == inf {
		// not connected
		return -1
	}

	d2 := bfs(n - 1)

	res := inf
	for i := range n {
		res = min(res, max(d1[i], d2[i]))
	}

	return res
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
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
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
	ride     int
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

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	it.index = -1
	*pq = old[:n-1]
	return it
}

func (pq *PQ) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
