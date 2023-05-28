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
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	ask := func(arr []int) int {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNum(reader)
	}

	res := solve(n, E, ask)
	fmt.Printf("! %d %d\n", res[0], res[1])
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

func solve(n int, E [][]int, ask func([]int) int) []int {

	g := NewGraph(n+1, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	ord := make([]int, 0, 2*n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		ord = append(ord, u)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				ord = append(ord, u)
			}
		}

	}

	dfs(0, 1)

	query := func(arr []int) (bool, int) {
		mem := make(map[int]bool)
		var tmp []int
		for _, x := range arr {
			if !mem[x] {
				tmp = append(tmp, x)
			}
			mem[x] = true
		}
		if len(mem) < 2 {
			return false, 0
		}
		return true, ask(tmp)
	}

	_, ans := query(ord)

	l, r := 0, len(ord)-1

	for l < r {
		mid := (l + r) / 2
		if mid == l {
			break
		}
		ok, tmp := query(ord[l : mid+1])
		if !ok {
			break
		}
		if tmp == ans {
			r = mid
		} else {
			l = mid
		}
	}

	return ord[l : l+2]
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
