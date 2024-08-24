package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	res, ord := solve(a, b)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", res))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", ord[i]))
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

const inf = 1 << 30

func solve(a []int, b []int) (int, []int) {
	n := len(b)
	g := NewGraph(n, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		if b[i] > 0 {
			b[i]--
		}
		if b[i] >= 0 {
			deg[b[i]]++
			g.AddEdge(i, b[i])
		}
	}

	que := make([]int, n)
	var front, end int

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[front] = i
			front++
		}
	}

	g2 := NewGraph(n, n)
	d2 := make([]int, n)

	var res int

	for end < front {
		u := que[end]
		res += a[u]
		end++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if a[u] >= 0 {
				// 先u后v
				g2.AddEdge(u, v)
				d2[v]++
			} else {
				// 先v后u
				g2.AddEdge(v, u)
				d2[u]++
			}
			a[v] += max(0, a[u])
			deg[v]--
			if deg[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	front = 0
	end = 0

	for i := 0; i < n; i++ {
		if d2[i] == 0 {
			que[front] = i
			front++
		}
	}

	for end < front {
		u := que[end]
		end++
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			d2[v]--
			if d2[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	for i := 0; i < n; i++ {
		que[i]++
	}

	return res, que
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
