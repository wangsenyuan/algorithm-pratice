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

	for {
		n := readNum(reader)
		if n == 0 {
			break
		}

		V := readNNums(reader, n)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(reader, n)
		}
		res := solve(n, grid, V)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const INF = 1 << 30

const X = 420
const D = 1 << 6

var valid [X][D]bool

func init() {
	for d := 0; d < D; d++ {
		for x := 0; x < X; x++ {
			valid[x][d] = true
			for i := 0; i < 6 && valid[x][d]; i++ {
				if (d>>i)&1 == 1 && (x%(i+2)) != 0 {
					valid[x][d] = false
				}
			}
		}
	}
}

func solve(n int, grid [][]int, V []int) int {

	g := NewGraph(n, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				g.AddEdge(i, j)
			}
		}
	}

	pq := make(PQ, n*D*X)

	dist := make([][][]*Item, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][]*Item, D)
		for j := 0; j < D; j++ {
			dist[i][j] = make([]*Item, X)
			for k := 0; k < X; k++ {
				item := new(Item)
				item.id = i
				item.dig = j
				item.rem = k
				item.value = INF
				item.index = i*D*X + j*X + k
				dist[i][j][k] = item
				pq[i*D*X+j*X+k] = item
			}
		}
	}
	update := func(i int, d int, r int, v int) {
		pq.update(dist[i][d][r], v)
	}

	if V[0] > 1 {
		update(0, 1<<(V[0]-2), V[0], V[0])
	} else {
		update(0, 0, V[0], V[0])
	}

	res := INF

	check := func(d int, x int) bool {
		return valid[x][d]
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)

		if cur.id == n-1 && check(cur.dig, cur.rem) {
			res = cur.value
			break
		}

		for i := g.node[cur.id]; i > 0; i = g.next[i] {
			u := g.to[i]
			w := (cur.rem*10 + V[u]) % X
			tmp := cur.value + V[u]
			d := cur.dig
			if V[u] > 1 {
				d |= (1 << (V[u] - 2))
			}
			if dist[u][d][w].value > tmp {
				update(u, d, w, tmp)
			}
		}

	}

	if res < INF {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	id    int
	dig   int
	rem   int
	value int
	index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
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
	re := old[n-1]
	re.index = -1
	*pq = old[:n-1]
	return re
}

func (pq *PQ) update(item *Item, v int) {
	item.value = v
	heap.Fix(pq, item.index)
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
