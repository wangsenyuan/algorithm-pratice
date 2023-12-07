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
		n, m, p := readThreeNums(reader)
		W := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(p, n, E, W)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

const inf = 1 << 60

func solve(p int, n int, E [][]int, W []int) int {
	g := NewGraph(n, len(E))
	for _, e := range E {
		u, v, s := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, s)
	}

	dp := make([][]*State, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]*State, n)
		for j := 0; j < n; j++ {
			dp[i][j] = new(State)
			dp[i][j].id = i
			dp[i][j].best = j
			dp[i][j].numShows = inf
			dp[i][j].money = 0
		}
	}

	pq := make(PriorityQueue, 0, n*n)

	push := func(i int, j int, x int, m int) {
		if dp[i][j].numShows < x || dp[i][j].numShows == x && dp[i][j].money > m {
			// have a better solution already
			return
		}
		dp[i][j].numShows = x
		dp[i][j].money = m
		heap.Push(&pq, dp[i][j])
	}

	push(0, 0, 0, p)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*State)
		if !dp[cur.id][cur.best].Equals(cur) {
			continue
		}
		u := cur.id
		best := cur.best
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			cost := g.val[i]
			// go from u to v
			nextBest := best
			if W[best] <= W[v] {
				nextBest = v
			}
			// just do lest number of shows
			// cur.money + j * W[best] >= cost
			j := max(0, (cost-cur.money+W[best]-1)/W[best])
			push(v, nextBest, cur.numShows+j, cur.money+j*W[best]-cost)
		}
	}

	ans := inf
	for i := 0; i < n; i++ {
		ans = min(ans, dp[n-1][i].numShows)
	}
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type State struct {
	id       int
	best     int
	numShows int
	money    int
	index    int
}

func (this *State) Equals(that *State) bool {
	return this.numShows == that.numShows && this.money == that.money
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].numShows < pq[j].numShows {
		return true
	}
	if pq[i].numShows == pq[j].numShows && pq[i].money > pq[j].money {
		return true
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(state *State, num int, money int) {
	state.numShows = num
	state.money = money
	heap.Fix(pq, state.index)
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
