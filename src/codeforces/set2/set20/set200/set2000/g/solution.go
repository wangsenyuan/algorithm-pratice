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
	n, m := readTwoNums(reader)
	tt := readNNums(reader, 3)
	streets := make([][]int, m)
	for i := range m {
		streets[i] = readNNums(reader, 4)
	}
	return solve(n, streets, tt)
}

const inf = 1 << 60

func solve(n int, streets [][]int, tt []int) int {
	m := len(streets)
	g := NewGraph(n, m*2)
	for i, cur := range streets {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	// 二分最晚出发时间late, 必须保证在时刻tt[1]时，是在节点,可以下车，然后下个走路到达下个路口

	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = i
		items[i] = it
	}

	travel := func(late int) bool {
		pq := make(PriorityQueue, n)
		for i := range n {
			it := items[i]
			it.priority = inf
			it.index = i
			pq[i] = it
		}
		pq[0].priority = late
		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				eid := g.val[i]
				l1 := streets[eid][2]
				l2 := streets[eid][3]
				nextStop := inf
				if it.priority+l1 <= tt[1] || it.priority >= tt[2] {
					// 在会议前后，都可以乘车
					nextStop = it.priority + l1
				} else {
					// 在这站就下车，并且走路过去
					// 等在这站，直到会议结束，并乘车到下站
					nextStop = min(it.priority+l2, tt[2]+l1)
				}
				if items[v].priority > nextStop {
					pq.update(items[v], nextStop)
				}
			}
		}
		return items[n-1].priority <= tt[0]
	}

	l, r := 0, tt[0]
	for l < r {
		mid := (l + r) / 2
		if travel(mid) {
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

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int // The value of the item; arbitrary.
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, v int) {
	item.priority = v
	heap.Fix(pq, item.index)
}
