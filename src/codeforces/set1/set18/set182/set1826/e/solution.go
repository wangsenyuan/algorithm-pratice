package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&p[i])
	}

	arr := make([]int, m)

	res := solve(n, m, p, func() []int {
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}
		return arr
	})
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
func solve(n int, m int, p []int, f func() []int) int {
	sets := make([]*BitSet, m)
	for i := 0; i < m; i++ {
		sets[i] = NewBitSet(m)
		for j := 0; j < m; j++ {
			sets[i].Set(j)
		}
	}

	id := make([]*Pair, m)
	for i := 0; i < m; i++ {
		id[i] = new(Pair)
		id[i].first = 0
		id[i].second = i
	}
	tmp := NewBitSet(m)
	for i := 0; i < n; i++ {
		cur := f()
		for j := 0; j < m; j++ {
			id[j].first = cur[j]
			id[j].second = j
			tmp.Clear(j)
		}

		sort.Slice(id, func(i, j int) bool {
			return id[i].first < id[j].first
		})

		for j := 0; j < m; {
			k := j
			for j < m && id[j].first == id[k].first {
				sets[id[j].second].And(tmp)
				j++
			}
			for k < j {
				tmp.Set(id[k].second)
				k++
			}
		}
	}

	g := NewGraph(m, m*m)
	deg := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j && sets[i].IsSet(j) {
				g.AddEdge(j, i)
				deg[i]++
			}
		}
	}

	// 越前面的越小

	item := make([]*Item, m)
	for i := 0; i < m; i++ {
		item[i] = new(Item)
		item[i].id = i
		item[i].value = p[i]
		item[i].index = -1
	}

	pq := make(PriorityQueue, 0, m)

	var best int

	for i := 0; i < m; i++ {
		if deg[i] == 0 {
			heap.Push(&pq, item[i])
		}
		// 即使，无法找到2个人以上的队伍，也可以只安排最有价值的人参加
		best = max(best, p[i])
	}

	for pq.Len() > 0 {
		// it is dag, so no loop back
		cur := heap.Pop(&pq).(*Item)
		u := cur.id

		best = max(best, cur.value)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]

			if item[v].value < cur.value+p[v] {
				item[v].value = cur.value + p[v]
			}
			deg[v]--
			if deg[v] == 0 {
				heap.Push(&pq, item[v])
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Item struct {
	id    int
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) Clear(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	if (bs.arr[i]>>j)&1 == 1 {
		bs.arr[i] ^= 1 << uint64(j)
	}
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (this *BitSet) And(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] &= that.arr[i]
	}
}
