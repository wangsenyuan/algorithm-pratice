package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	query := func(arr []int) bool {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]+1))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNum(reader) == 1
	}

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, query)
		res++
		fmt.Printf("! %d\n", res)
	}
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

func solve(n int, E [][]int, query func([]int) bool) int {
	if n == 1 {
		return 0
	}
	g := buildGraph(n, E)
	var leaf []int
	parent := make([]int, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		ok := true
		parent[u] = p
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				ok = false
				dfs(u, v)
			}
		}
		if ok && u != 0 {
			leaf = append(leaf, u)
		}
	}

	dfs(-1, 0)

	l, r := 0, len(leaf)-1

	arr := make([]int, n)
	var node int
	for l <= r {
		mid := (l + r) / 2
		copy(arr[1:], leaf[l:mid+1])
		if query(arr[:1+(mid-l+1)]) {
			node = leaf[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	r = 0
	for node != -1 {
		arr[r] = node
		r++
		node = parent[node]
	}

	l = 0
	r--
	for l <= r {
		mid := (l + r) / 2
		if query(arr[l : mid+1]) {
			r = mid - 1
			node = arr[mid]
		} else {
			l = mid + 1
		}
	}

	return node
}

func buildGraph(n int, E [][]int) *Graph {
	g := NewGraph(n, 2*len(E)+1)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	return g
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
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
