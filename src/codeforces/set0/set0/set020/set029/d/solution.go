package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	getOrd := func(m int) []int {
		return readNNums(reader, m)
	}
	res := solve(n, edges, getOrd)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
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

func solve(n int, edges [][]int, getOrder func(int) []int) []int {
	g := NewGraph(n, n*2)
	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		var res int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				res += dfs(u, v)
			}
		}
		return max(res, 1)
	}

	cnt := dfs(-1, 0)

	ord := getOrder(cnt)
	pos := make([]int, n)
	first := make([]int, n)
	last := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
		first[i] = n
		last[i] = -1
	}

	for i, u := range ord {
		pos[u-1] = i
	}

	type pair struct {
		first  int
		second int
	}

	children := make([][]pair, n)

	var dfs2 func(p int, u int) bool
	dfs2 = func(p int, u int) bool {
		if pos[u] >= 0 {
			// a leaf
			first[u] = pos[u]
			last[u] = pos[u]
			return true
		}
		var arr []pair
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs2(u, v)
				if !tmp {
					return false
				}
				arr = append(arr, pair{first[v], v})
			}
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].first < arr[j].first
		})

		for i := 1; i < len(arr); i++ {
			v := arr[i-1].second
			w := arr[i].second
			if last[v]+1 != first[w] {
				return false
			}
		}
		first[u] = first[arr[0].second]
		last[u] = last[arr[len(arr)-1].second]
		children[u] = arr
		return true
	}

	if !dfs2(-1, 0) {
		return nil
	}
	var res []int
	var dfs3 func(u int)

	dfs3 = func(u int) {
		res = append(res, u+1)

		for _, x := range children[u] {
			v := x.second
			dfs3(v)
			res = append(res, u+1)
		}
	}

	dfs3(0)

	return res
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
