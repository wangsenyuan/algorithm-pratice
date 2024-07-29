package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	x, y := readTwoNums(reader)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 3)
	}
	taxis := make([][]int, n)
	for i := 0; i < n; i++ {
		taxis[i] = readNNums(reader, 2)
	}

	res := solve(n, m, x, y, roads, taxis)

	fmt.Println(res)
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

const inf = 1 << 60

func solve(n int, m int, x int, y int, roads [][]int, taxis [][]int) int {
	x--
	y--
	// 先需要知道任意两点间的最短距离
	g := NewGraph(n, 2*m)

	for _, road := range roads {
		u, v, w := road[0]-1, road[1]-1, road[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		items[i] = it
	}

	g2 := NewGraph(n, n*n*2)

	dp := make([]int, n)

	bfs := func(s int) {
		pq := make(PQ, n)
		for i := 0; i < n; i++ {
			items[i].priority = inf
			items[i].index = i
			pq[i] = items[i]
			dp[i] = inf
		}

		items[s].priority = 0

		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			if it.priority >= inf {
				break
			}
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v].priority >= inf && it.priority+g.val[i] <= taxis[s][0] {
					dp[v] = taxis[s][1]
					pq.update(items[v], it.priority+g.val[i])
				}
			}
		}

		for i := 0; i < n; i++ {
			if i == s || dp[i] >= inf {
				continue
			}
			g2.AddEdge(s, i, taxis[s][1])
		}
	}
	// n * m
	for i := 0; i < n; i++ {
		bfs(i)
	}

	pq := make(PQ, n)
	for i := 0; i < n; i++ {
		items[i].priority = inf
		items[i].index = i
		pq[i] = items[i]
	}

	items[x].priority = 0
	heap.Init(&pq)

	// n * n
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			w := g2.val[i]
			if it.priority+w < items[v].priority {
				pq.update(items[v], it.priority+w)
			}
		}
	}

	if items[y].priority >= inf {
		return -1
	}

	return items[y].priority
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
	e++
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
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
