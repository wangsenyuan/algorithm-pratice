package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := readNNums(reader, n)
	deps := make([][]int, m)
	for i := 0; i < m; i++ {
		deps[i] = readNNums(reader, 2)
	}

	res := solve(n, deps, E)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, deps [][]int, E []int) int {
	// 这里要两个队列才行
	g := NewGraph(n, len(deps))
	deg := make([]int, n)
	for _, dep := range deps {
		u, v := dep[0], dep[1]
		g.AddEdge(v, u)
		deg[u]++
	}

	q0 := NewQueue(n)
	q1 := NewQueue(n)

	add := func(i int) {
		if E[i] == 0 {
			q0.PushBack(i)
		} else {
			q1.PushBack(i)
		}
	}

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			add(i)
		}
	}

	var res int
	for !q0.IsEmpty() || !q1.IsEmpty() {
		// 先执行main的
		for !q0.IsEmpty() {
			u := q0.PopFront()
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				deg[v]--
				if deg[v] == 0 {
					add(v)
				}
			}
		}

		if !q1.IsEmpty() {
			res++
			for !q1.IsEmpty() {
				u := q1.PopFront()
				for i := g.nodes[u]; i > 0; i = g.next[i] {
					v := g.to[i]
					deg[v]--
					if deg[v] == 0 {
						add(v)
					}
				}
			}
		}
	}

	return res
}

type Queue struct {
	arr   []int
	front int
	tail  int
}

func NewQueue(n int) *Queue {
	arr := make([]int, n*2)
	return &Queue{arr, n, n}
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *Queue) PushBack(u int) {
	q.arr[q.tail] = u
	q.tail++
}

func (q *Queue) PopFront() int {
	u := q.arr[q.front]
	q.front++
	return u
}

func (q *Queue) PushFront(u int) {
	q.front--
	q.arr[q.front] = u
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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

type Task struct {
	id    int
	deg   int
	kind  int
	index int
}
