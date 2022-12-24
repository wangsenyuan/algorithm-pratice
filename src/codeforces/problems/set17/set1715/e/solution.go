package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, m, k := readThreeNums(reader)
	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	res := solve(n, k, E)

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1e15

func solve(n int, k int, E [][]int) []int64 {

	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dist := make([]int64, n)

	for i := 1; i < n; i++ {
		dist[i] = INF
	}
	vis := make([]bool, n)
	items := make([]*Item, n)
	pq := make(PQ, 0, n)

	dijkstra := func() {
		for i := 0; i < n; i++ {
			items[i] = new(Item)
			items[i].id = i
			items[i].priority = dist[i]
			items[i].index = i
			vis[i] = false
		}

		pq = append(pq, items...)

		heap.Init(&pq)

		for pq.Len() > 0 {
			u := heap.Pop(&pq).(*Item)
			if vis[u.id] {
				continue
			}
			vis[u.id] = true
			for i := g.node[u.id]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if dist[v] > dist[u.id]+int64(w) {
					dist[v] = dist[u.id] + int64(w)
					pq.update(items[v], dist[v])
				}
			}
		}
	}

	dijkstra()

	for k > 0 {
		k--
		cht := NewCHT(NewLine(0, 0))
		for i := 1; i < n; i++ {
			cht.AddLine(NewLine(-2*int64(i), dist[i]+int64(i)*int64(i)))
		}
		for i := 0; i < n; i++ {
			dist[i] = cht.Query(int64(i)) + int64(i)*int64(i)
		}

		dijkstra()
	}

	return dist
}

type Line struct {
	k int64
	b int64
}

func NewLine(k int64, b int64) Line {
	return Line{k, b}
}

func (this Line) intersect(that Line) float64 {
	db := float64(that.b - this.b)
	dk := float64(this.k - that.k)
	return db / dk
}

func (this Line) apply(x int64) int64 {
	return int64(this.k)*x + int64(this.b)
}

type CHT struct {
	x     []float64
	lines []Line
}

func NewCHT(line Line) *CHT {
	var x []float64
	var lines []Line
	x = append(x, -math.MaxFloat32)
	lines = append(lines, line)
	return &CHT{x, lines}
}

func (this *CHT) AddLine(l Line) {
	for len(this.lines) >= 2 && l.intersect(this.lines[len(this.lines)-2]) <= back(this.x) {
		this.x = this.x[:len(this.x)-1]
		this.lines = this.lines[:len(this.lines)-1]
	}
	if len(this.lines) > 0 {
		this.x = append(this.x, l.intersect(back(this.lines)))
	}
	this.lines = append(this.lines, l)
}

func (this *CHT) Query(qx int64) int64 {
	l, r := 0, len(this.x)

	for l < r {
		mid := (l + r) / 2
		if this.x[mid] > float64(qx) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	r--
	return this.lines[r].apply(qx)
}

type Number interface {
	float64 | Line
}

func back[T Number](arr []T) T {
	n := len(arr)
	return arr[n-1]
}

type Item struct {
	id       int
	priority int64
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
	*pq = old[:n-1]
	return rt
}

func (pq *PQ) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
