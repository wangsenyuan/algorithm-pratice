package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, E [][]int) []int {
	// 相连的两个house不能同时有老师
	// 一旦一个house里面有老师了，那么和它相连的边（house）都被选中了
	// 是不是只要联通就有答案呐？
	// 我还是没有搞清楚什么清空下，肯定有解
	// 不连通的图，肯定是不行的
	// 一棵树，肯定是有答案的，按奇偶层即可
	// 一个图删掉边后，就变成了一棵树, 但是，如果有某个节点，它的层级，即是偶数，又是奇数
	// 怎么处理呢？
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	black := NewQueue(n)
	white := NewQueue(n)

	black.PushEnd(0)
	vis[0] = true

	var res []int

	for !black.IsEmpty() || !white.IsEmpty() {
		if black.IsEmpty() {
			found := false

			for !found && !white.IsEmpty() {
				v := white.PopFront()
				for i := g.nodes[v]; i > 0; i = g.next[i] {
					w := g.to[i]
					if !vis[w] {
						vis[w] = true
						black.PushEnd(w)
						g.nodes[v] = g.next[i]
						found = true
						break
					}
				}
				if found && g.nodes[v] > 0 {
					white.PushEnd(v)
				}
			}
			if !found {
				break
			}
		}
		u := black.PopFront()
		res = append(res, u+1)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				vis[v] = true
				white.PushEnd(v)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			return nil
		}
	}

	return res
}

type Queue struct {
	arr   []int
	front int
	end   int
}

func NewQueue(n int) *Queue {
	arr := make([]int, n)
	var front, end int
	return &Queue{arr, front, end}
}

func (q *Queue) PushEnd(u int) {
	q.arr[q.end] = u
	q.end++
}

func (q *Queue) PopFront() int {
	u := q.arr[q.front]
	q.front++
	return u
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.end
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
