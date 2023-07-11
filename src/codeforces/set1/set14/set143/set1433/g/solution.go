package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	Q := make([][]int, k)
	for i := 0; i < k; i++ {
		Q[i] = readNNums(reader, 2)
	}
	fmt.Println(solve(n, E, Q))
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

const inf = 1 << 30

func solve(n int, E [][]int, Q [][]int) int {
	// 如果知道每条边在答案中的贡献的话，
	// 是否可行呢？
	// 如果一条边在一个路径中，那么它变成0后，仍然在路径中
	// 但是如果一条边不在路径中，那么它变成0后，也会在路径中
	// 但是这样的变化会是更优的答案吗？
	// 考虑一条边，它的weight非常大，远大于其他边，那么它一开始不存在任何路径中，但是它变成0后，就可以使得很多的路径改道
	// dist[a][b] 表示未做任何改变时，a到b的最短距离
	// 这个可以在 n * m 时间内完成
	// 假设要建某条边变成0，如果dist[a][b] = dist[a][u] + weight[u][v] + dist[v][b]
	//  边在这条路径中，贡献时 -weight[u][v]
	// 如果 dist[a][b] < ... 但是dist[a][u] + dist[v][b] < dist[a][b], 那么它的贡献就是 dist[a][b] - dist[a][u] - dist[v][b]

	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dist := make([][]*Item, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			dist[i][j] = new(Item)
			dist[i][j].id = j
			dist[i][j].index = j
			dist[i][j].cost = inf
		}
		dist[i][i].cost = 0
	}
	pq := make(PriorityQueue, 0, n)

	bfs := func(x int) {
		for i := 0; i < n; i++ {
			heap.Push(&pq, dist[x][i])
		}
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[x][v].cost > dist[x][u].cost+g.val[i] {
					pq.update(dist[x][v], dist[x][u].cost+g.val[i])
				}
			}
		}
	}

	for u := 0; u < n; u++ {
		bfs(u)
	}

	var best int
	for _, cur := range Q {
		a, b := cur[0]-1, cur[1]-1
		best += dist[a][b].cost
	}

	tot := best

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		var save int
		for _, cur := range Q {
			a, b := cur[0]-1, cur[1]-1
			if dist[a][b].cost == dist[a][u].cost+w+dist[b][v].cost {
				save += w
			} else if dist[a][b].cost == dist[a][v].cost+w+dist[b][u].cost {
				save += w
			} else {
				tmp := min(dist[a][v].cost+dist[b][u].cost, dist[a][u].cost+dist[b][v].cost)
				if dist[a][b].cost > tmp {
					save += dist[a][b].cost - tmp
				}
			}
		}
		if tot-save < best {
			best = tot - save
		}
	}

	return best
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

// An Item is something we manage in a priority queue.
type Item struct {
	id   int // The value of the item; arbitrary.
	cost int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, cost int) {
	item.cost = cost
	heap.Fix(pq, item.index)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
