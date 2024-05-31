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

	n, m := readTwoNums(reader)
	trains := make([][]int, m)
	for i := 0; i < m; i++ {
		trains[i] = readNNums(reader, 6)
	}

	res := solve(n, trains)

	var buf bytes.Buffer
	for _, x := range res {
		if x >= 0 {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		} else {
			buf.WriteString("Unreachable\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, trains [][]int) []int {
	m := len(trains)

	g := NewGraph(n+1, m)

	for i, cur := range trains {
		a, b := cur[4], cur[5]
		g.AddEdge(b, a, i)
	}

	items := make([]*Item, n+1)
	pq := make(PQ, n)

	for i := 1; i <= n; i++ {
		it := new(Item)
		it.id = i
		it.priority = -inf
		it.index = i - 1
		items[i] = it
		pq[i-1] = it
	}

	heap.Init(&pq)

	for i := g.nodes[n]; i > 0; i = g.next[i] {
		v := g.to[i]
		w := g.val[i]
		tr := trains[w]
		l, d, k := tr[0], tr[1], tr[2]
		pq.update(items[v], l+(k-1)*d)
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority < 0 {
			break
		}
		u := it.id
		tu := it.priority

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			tr := trains[w]
			l, d, k, c := tr[0], tr[1], tr[2], tr[3]
			if tu-c >= l {
				x := min((tu-c-l)/d, k-1)
				tv := l + d*x
				if items[v].priority < tv {
					pq.update(items[v], tv)
				}
			}
		}
	}

	ans := make([]int, n-1)

	for i := 1; i < n; i++ {
		ans[i-1] = items[i].priority
		if ans[i-1] < 0 {
			ans[i-1] = -1
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(item interface{}) {
	cur := item.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
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
