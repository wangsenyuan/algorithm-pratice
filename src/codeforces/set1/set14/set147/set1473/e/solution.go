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

	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

const inf = 1e18

func solve(n int, E [][]int) []int64 {
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dp := make([][][]*Item, n)
	pq := make(PriorityQueue, n*2*2)

	for i := 0; i < n; i++ {
		dp[i] = make([][]*Item, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]*Item, 2)
			for k := 0; k < 2; k++ {
				it := new(Item)
				it.id = i
				it.mx = j
				it.mn = k
				it.priority = inf
				dp[i][j][k] = it
				it.index = i*4 + j*2 + k
				pq[i*4+j*2+k] = it
			}
		}
	}

	pop := func() (id int, mx int, mn int, v int64) {
		it := heap.Pop(&pq).(*Item)
		id = it.id
		mx = it.mx
		mn = it.mn
		v = it.priority
		return
	}

	dp[0][0][0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		u, mx, mn, we := pop()
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := int64(g.val[i])
			for a := 0; a <= 1-mx; a++ {
				for b := 0; b <= 1-mn; b++ {
					if dp[v][a|mx][b|mn].priority > we+int64(1-a+b)*w {
						pq.update(dp[v][a|mx][b|mn], we+int64(1-a+b)*w)
					}
				}
			}
		}
	}

	ans := make([]int64, n-1)

	for i := 1; i < n; i++ {
		ans[i-1] = dp[i][1][1].priority
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
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
	mx       int
	mn       int
	priority int64
	index    int
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
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(it *Item, v int64) {
	it.priority = v
	heap.Fix(pq, it.index)
}
