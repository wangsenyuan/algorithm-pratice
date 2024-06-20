package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	branches := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		branches[i] = readNNums(reader, 2)
	}
	res := solve(n, branches)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(n int, branches [][]int) bool {
	// 始终只能从少的移动到多的地方
	// 考虑一条线？这条线的长度如果超过4，就无法聚集在一起
	//  任何一个叶子节点都不可能成为目标位置
	// 这是因为，这个叶子节点为1，从它的parent那里过来的，必然超过1
	// 或者parent移动到它这里后，那么parent就变成了0，其他ant移动不进去
	// 所以ant只能移动到它的parent节点去
	// 但是到了parent似乎也进行不下去了？除非它的parent现在有2个了？
	// 所以用优先队列？
	g := NewGraph(n+1, 2*n)
	deg := make([]int, n+1)
	for _, b := range branches {
		u, v := b[0], b[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	vis := make([]bool, n+1)

	active := make(PriorityQueue, 0, n)
	cnt := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cnt[i] = 1
		if deg[i] == 1 {
			vis[i] = true
			it := new(Item)
			it.id = i
			it.priority = 1
			heap.Push(&active, it)
		}
	}

	for active.Len() > 0 {
		// 它们只能进入parent
		it := heap.Pop(&active).(*Item)
		u := it.id
		// deg[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] {
				continue
			}
			// v is the parent
			if cnt[v] < cnt[u] {
				return false
			}
			cnt[v] += cnt[u]
			deg[v]--
			if deg[v] == 1 {
				vis[v] = true
				it := new(Item)
				it.id = v
				it.priority = cnt[v]
				heap.Push(&active, it)
			}
		}
	}

	return true
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
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
