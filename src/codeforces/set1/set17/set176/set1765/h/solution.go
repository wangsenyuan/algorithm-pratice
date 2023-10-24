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
	p := readNNums(reader, n)
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = readTwoNums(reader)
	}
	res := solve(p, a, b)
	var buf bytes.Buffer
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

func solve(p []int, a []int, b []int) []int {
	n := len(p)
	people := make([]*Person, n)
	for i := 0; i < n; i++ {
		people[i] = new(Person)
		people[i].deadLine = p[i]
		people[i].id = i
		people[i].index = i
		people[i].deg = 0
	}
	deg := make([]int, n)
	g := NewGraph(n, len(a))
	for i := 0; i < len(a); i++ {
		g.AddEdge(b[i]-1, a[i]-1)
		deg[a[i]-1]++
	}
	ans := make([]int, n)
	for s := 0; s < n; s++ {
		pq := make(PQ, 0, n)
		for i := 0; i < n; i++ {
			people[i].deg = deg[i]
			if deg[i] == 0 && i != s {
				heap.Push(&pq, people[i])
			}
		}
		for i := n; i > 0; i-- {
			if pq.Len() == 0 || pq[0].deadLine < i {
				if people[s].deg == 0 && i <= people[s].deadLine {
					ans[s] = i
				}
				break
			}
			u := heap.Pop(&pq).(*Person).id
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				people[v].deg--
				if v != s && people[v].deg == 0 {
					heap.Push(&pq, people[v])
				}
			}
		}
	}

	return ans
}

type Person struct {
	id       int
	deg      int
	deadLine int
	index    int
}

type PQ []*Person

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].deadLine > pq[j].deadLine
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	p := x.(*Person)
	p.index = len(*pq)
	*pq = append(*pq, p)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[:n-1]
	return old[n-1]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
