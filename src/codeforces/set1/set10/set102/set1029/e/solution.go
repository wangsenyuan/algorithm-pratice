package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

const inf = 1 << 30

func solve(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([][]int, n)

	// dp[u][0] 虽然在u处没有连线，但是u已经被cover住了，不需要通过子节点上的线来处理
	// dp[u][1] 表示在u处连线
	// dp[u][2] 表示在u处没有连线，且它的父节点也没有连线

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		dp[u] = []int{0, 1, 0, inf}
		if p == 0 {
			// 0的直接子节点，不需要额外的线
			dp[u][1] = 0
		}
		best := inf
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				dp[u][1] += min(dp[v][1], dp[v][2])
				dp[u][2] += min(dp[v][1], dp[v][3])
				best = min(best, dp[v][1]-min(dp[v][1], dp[v][3]))
			}
		}

		dp[u][3] = dp[u][2] + best
	}

	var res int

	for i := g.nodes[0]; i > 0; i = g.next[i] {
		u := g.to[i]
		dfs(0, u)
		res += min(dp[u][1], dp[u][3])
	}

	return res
}

func solve1(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)
	fa := make([]int, n)
	dist := make([]int, n)

	var head, tail int
	que[head] = 0
	head++

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] == 0 && v != 0 {
				dist[v] = dist[u] + 1
				fa[v] = u
				que[head] = v
				head++
			}
		}
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.index = i
		it.priority = dist[i]
		items[i] = it
		pq[i] = it
	}

	heap.Init(&pq)
	var ans int
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority <= 2 {
			break
		}
		// dist[it.id] > 2
		p := fa[it.id]
		ans++
		pq.update(items[p], 1)
		for i := g.nodes[p]; i > 0; i = g.next[i] {
			v := g.to[i]
			if items[v].index < 0 || items[v].priority <= 2 {
				continue
			}
			pq.update(items[v], 2)
		}
	}

	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
