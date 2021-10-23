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

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res, C := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", C[i]))
		}
		buf.WriteByte('\n')
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func solve(n int, E [][]int) (int, []int) {
	g := NewGraph(n, len(E)*2)
	D := make([]int, n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		D[u]++
		D[v]++
	}

	items := make([]*Item, n)
	pq := make(PQ, n)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.index = i
		it.value = i
		it.priority = D[i]
		items[i] = it
		pq[i] = it
	}

	heap.Init(&pq)

	cur := n
	C := make([]int, n)
	var ans int

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		ans = max(ans, it.priority)

		u := it.value

		C[u] = cur
		cur--

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if items[v].index >= 0 {
				pq.update(items[v], items[v].priority-1)
			}
		}
	}

	return ans, C
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(n int, E [][]int) (int, []int) {
	g := NewGraph(n, len(E)*2)
	D := make([]int, n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		D[u]++
		D[v]++
	}

	items := make([]*Item, n)
	pq := make(PQ, 0, n)

	var right int

	for i := 0; i < n; i++ {
		it := new(Item)
		it.index = i
		it.value = i
		it.priority = D[i]
		items[i] = it
		if right < D[i] {
			right = D[i]
		}
	}

	C := make([]int, n)

	check := func(k int) bool {
		pq = pq[:0]
		for i := 0; i < n; i++ {
			items[i].priority = D[i]
			items[i].index = i
			pq = append(pq, items[i])
		}
		heap.Init(&pq)

		cur := n

		for pq.Len() > 0 {
			u := heap.Pop(&pq).(*Item)
			if u.priority > k {
				return false
			}
			C[u.value] = cur
			cur--
			for i := g.nodes[u.value]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v].index >= 0 {
					// still not processed
					pq.update(items[v], items[v].priority-1)
				}
			}
		}
		return true
	}

	var left int
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// fill the result
	check(right)

	return right, C
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Item struct {
	value    int
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

func (pq *PQ) Push(x interface{}) {
	i := x.(*Item)
	i.index = len(*pq)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	rt := old[n-1]
	rt.index = -1
	*pq = old[:n-1]
	return rt
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
