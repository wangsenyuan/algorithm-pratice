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
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		C := make([][]int, k)
		for i := 0; i < k; i++ {
			C[i] = readNNums(reader, 2)
		}
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, m, k, E, C)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

const INF = math.MaxInt64

func solve(N, M, K int, E [][]int, C [][]int) []int64 {
	g := NewGraph(N, 2*M)
	for _, e := range E {
		u, v, c := e[0], e[1], e[2]
		u--
		v--
		g.Add(u, v, int64(c))
		g.Add(v, u, int64(c))
	}

	cities := make([]*City, N)
	pq := make(PQ, N)

	for i := 0; i < N; i++ {
		cur := new(City)
		cur.index = i
		cur.pos = i
		cur.cost = INF
		cities[i] = cur
		pq[i] = cur
	}
	for i := 0; i < K; i++ {
		j, c := C[i][0], C[i][1]
		j--
		cities[j].cost = int64(c)
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*City)
		u := cur.pos
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if cities[v].cost > cur.cost+g.cost[i] {
				pq.update(cities[v], cur.cost+g.cost[i])
			}
		}
	}
	res := make([]int64, N)
	for i := 0; i < N; i++ {
		res[i] = cities[i].cost
	}
	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int64
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cost = make([]int64, e+2)
	g.cur = 0
	return g
}

func (g *Graph) Add(u, v int, c int64) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = c
}

type City struct {
	index int
	cost  int64
	pos   int
}

type PQ []*City

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].cost < this[j].cost
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *PQ) Push(x interface{}) {
	n := len(*this)
	city := x.(*City)
	city.index = n
	*this = append(*this, city)
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	res.index = -1
	*this = old[:n-1]
	return res
}

func (this *PQ) update(city *City, cost int64) {
	city.cost = cost
	heap.Fix(this, city.index)
}
