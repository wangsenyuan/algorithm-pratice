package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (res []int, cmds []string, arcs [][]int) {
	n := readNum(reader)
	cmds = make([]string, n)
	for i := range n {
		cmds[i] = readString(reader)
	}
	m := readNum(reader)
	arcs = make([][]int, m)
	for i := range m {
		arcs[i] = readNNums(reader, 2)
	}
	res = solve(cmds, arcs)
	return
}

func solve(cmds []string, arcs [][]int) []int {
	n := len(cmds)

	g := NewGraph(n, len(arcs)+1)
	deg := make([]int, n)
	for _, arc := range arcs {
		u, v := arc[0]-1, arc[1]-1
		g.AddEdge(u, v)
		deg[v]++
	}

	ques := make([]*Queue, 4)
	for i := range 4 {
		ques[i] = NewQueue(n)
	}
	// 0 for unset false
	// 1 for set true
	// 2 for set false
	// 3 for unset true
	id := map[string]int{
		"unset false": 0,
		"set true":    1,
		"set false":   2,
		"unset true":  3,
	}

	add := func(i int, s string) {
		j := id[s]
		ques[j].Push(i)
	}

	var res []int

	process := func(u int) {
		res = append(res, u+1)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				add(v, cmds[v])
			}
		}
	}

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			add(i, cmds[i])
		}
	}

	for d := range 4 {
		if d&1 == 1 && ques[d].IsEmpty() {
			return nil
		}
		for !ques[d].IsEmpty() {
			u := ques[d].Pop()
			process(u)
		}
		if len(res) == n {
			break
		}
	}

	for !ques[0].IsEmpty() {
		u := ques[0].Pop()
		process(u)
	}

	if len(res) < n {
		return nil
	}

	return res
}

type Queue struct {
	arr  []int
	head int
	tail int
}

func NewQueue(n int) *Queue {
	arr := make([]int, n+1)
	return &Queue{arr, 0, 0}
}

func (q *Queue) Push(u int) {
	q.arr[q.head] = u
	q.head = (q.head + 1) % len(q.arr)
}

func (q *Queue) Pop() int {
	u := q.arr[q.tail]
	q.tail = (q.tail + 1) % len(q.arr)
	return u
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
