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
		n, m := readTwoNums(reader)
		options := make([][]int, m)
		for i := 0; i < m; i++ {
			options[i] = readNNums(reader, 3)
		}
		res := solve(n, options)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%.9f\n", ans))
		}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, E [][]int) []float64 {
	g := NewGraph(n, len(E))

	for _, e := range E {
		s, f, p := e[0], e[1], e[2]
		s--
		f--
		g.AddEdge(s, f, float64(p))
	}

	pq := make(PQ, 0, n)

	heap.Push(&pq, NewNode(0, 0))

	res := make([]float64, n)

	for i := 0; i < n; i++ {
		res[i] = math.MaxFloat64
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)
		u := cur.value
		if res[u] <= cur.priority {
			continue
		}
		res[u] = cur.priority
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.weight[i]
			heap.Push(&pq, NewNode(v, res[u]/2+w))
		}
	}

	for i := 0; i < n; i++ {
		if res[i] == math.MaxFloat64 {
			res[i] = -1
		}
	}

	return res
}

type Graph struct {
	nodes  []int
	next   []int
	to     []int
	weight []float64
	cur    int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.weight = make([]float64, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int, w float64) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.weight[g.cur] = w
}

type Node struct {
	value    int
	priority float64
	index    int
}

func NewNode(i int, p float64) *Node {
	node := new(Node)
	node.value = i
	node.priority = p
	return node
}

type PQ []*Node

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
	node := x.(*Node)
	node.index = len(*pq)
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}
