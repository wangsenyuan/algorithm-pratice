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
	n := readNum(reader)
	s := readString(reader)
	res := solve(n, s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, s string) []int {
	// = 的那部分必须一起考虑
	g := NewGraph(n, n)
	id := make([]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && s[i-1] == '=' {
			id[i] = id[i-1]
		} else {
			id[i] = i
		}
		if i > 0 {
			if s[i-1] == 'L' {
				g.AddEdge(id[i], id[i-1], 1)
				deg[id[i-1]]++
			} else if s[i-1] == 'R' {
				g.AddEdge(id[i-1], id[i], 1)
				deg[id[i]]++
			}
		}
	}
	res := make([]int, n)
	que := make([]int, n)
	var front, tail int
	for i := 0; i < n; i++ {
		if id[i] == i && deg[i] == 0 {
			que[front] = id[i]
			front++
			res[id[i]] = 1
		}
	}

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			res[v] = max(res[v], res[u]+1)
			deg[v]--
			if deg[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = res[id[i]]
	}

	return ans
}

func solve1(n int, s string) []int {
	// = 的那部分必须一起考虑
	g := NewGraph(n, n)
	deg := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && s[i-1] == '=' {
			id[i] = id[i-1]
		} else {
			id[i] = i
		}
		if i > 0 {
			if s[i-1] == 'L' {
				g.AddEdge(id[i], id[i-1], 1)
				deg[id[i-1]]++
			} else if s[i-1] == 'R' {
				g.AddEdge(id[i-1], id[i], 1)
				deg[id[i]]++
			}
		}
	}
	res := make([]int, n)
	pq := make(PQ, 0, n)

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			res[i] = 1
			it := new(Item)
			it.id = i
			it.priority = 1
			heap.Push(&pq, it)
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			res[v] = max(res[v], res[u]+w)
			deg[v]--
			if deg[v] == 0 {
				it := new(Item)
				it.id = v
				it.priority = res[v]
				heap.Push(&pq, it)
			}
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = res[id[i]]
	}

	return ans
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

func (pq *PQ) Push(item any) {
	cur := item.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() any {
	arr := *pq
	n := len(arr)
	res := arr[n-1]
	*pq = arr[:n-1]
	res.index = -1
	return res
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
