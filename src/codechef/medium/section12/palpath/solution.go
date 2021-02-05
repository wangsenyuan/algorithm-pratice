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
		line := readNNums(reader, 4)
		n, m, s, t := line[0], line[1], line[2], line[3]
		edges := make([]Edge, m)
		for i := 0; i < m; i++ {
			bb, _ := reader.ReadBytes('\n')
			edges[i] = NewEdge(bb)
		}
		res := solve(n, m, s, t, edges)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const INF = int64(1) << 60

func solve(n, m int, s, t int, edges []Edge) int64 {
	s--
	t--

	// at most n * n nodes
	dist := make([][]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
	}

	outs := make([][]int, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 10)

	}

	updateOut := func(x, i int) {
		outs[x] = append(outs[x], i)
	}

	for i, e := range edges {
		u, v := e.u, e.v
		u--
		v--
		updateOut(u, i)
		updateOut(v, i)
	}
	que := make(PriorityQueue, 0, n*n)
	dist[s][t] = 0
	dist[t][s] = 0

	heap.Push(&que, NewItem(s*n+t, dist[s][t]))

	var ans int64 = INF

	for que.Len() > 0 {
		cur := heap.Pop(&que).(*Item)
		if cur.priority == INF {
			break
		}
		u, v := cur.value/n, cur.value%n

		if cur.priority > dist[u][v] {
			continue
		}

		if u == v {
			ans = min(ans, dist[u][v])
			continue
		}

		for _, a := range outs[u] {
			uu := edges[a].GetAnotherEnd(u)

			for _, b := range outs[v] {
				vv := edges[b].GetAnotherEnd(v)

				if vv == u {
					ans = min(ans, dist[u][v]+int64(edges[b].cost))
					continue
				}

				if edges[a].label != edges[b].label {
					continue
				}

				if dist[uu][vv] > dist[u][v]+int64(edges[a].cost)+int64(edges[b].cost) {
					dist[uu][vv] = dist[u][v] + int64(edges[a].cost) + int64(edges[b].cost)
					dist[vv][uu] = dist[uu][vv]
					heap.Push(&que, NewItem(uu*n+vv, dist[uu][vv]))
				}
			}
		}
	}

	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Edge struct {
	u, v  int
	cost  int
	label byte
}

func NewEdge(buf []byte) Edge {
	var pos int
	var u, v, cost int
	pos = readInt(buf, pos, &u) + 1
	pos = readInt(buf, pos, &v) + 1
	pos = readInt(buf, pos, &cost) + 1
	label := buf[pos]
	return Edge{u, v, cost, label}
}

func (edge Edge) GetAnotherEnd(u int) int {
	u++
	return (edge.u ^ edge.v ^ u) - 1
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int   // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(value int, priority int64) *Item {
	item := new(Item)
	item.value = value
	item.priority = priority
	return item
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
