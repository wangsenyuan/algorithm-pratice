package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, s := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E, s)
	if len(res) == 0 {
		fmt.Println("Impossible")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Possible\n")
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", len(res[i])))
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

func solve(n int, E [][]int, s int) [][]int {
	g := NewGraph(n, len(E))

	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
	}

	s--

	que := make([]int, n)
	var front, end int
	pushBack := func(u int) {
		que[end] = u
		end++
	}
	popFront := func() int {
		u := que[front]
		front++
		return u
	}

	fa := make([]int, n)
	root := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = -1
		root[i] = -1
	}

	out := func(a, b int) [][]int {
		res := make([][]int, 2)
		for i := 0; i < 2; i++ {
			res[i] = make([]int, 0, n)
		}
		t := a
		for a != s {
			res[0] = append(res[0], a+1)
			a = fa[a]
		}
		res[0] = append(res[0], s+1)
		res[1] = append(res[1], t+1)
		for b != s {
			res[1] = append(res[1], b+1)
			b = fa[b]
		}
		res[1] = append(res[1], s+1)
		for i := 0; i < 2; i++ {
			reverse(res[i])
		}
		return res
	}

	for i := g.nodes[s]; i > 0; i = g.next[i] {
		v := g.to[i]
		fa[v] = s
		root[v] = v
		pushBack(v)
	}

	for front < end {
		u := popFront()

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if fa[v] != -1 && root[v] != root[u] {
				return out(v, u)
			}
			if v != s && fa[v] < 0 {
				fa[v] = u
				root[v] = root[u]
				pushBack(v)
			}
		}
	}

	return nil
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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
