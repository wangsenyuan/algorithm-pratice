package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = readNNums(reader, n)
	}

	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res, color := solve(n, edges, c)

	if res < 0 {
		fmt.Println(res)
		return
	}
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", res))
	for _, cur := range color {
		buf.WriteString(fmt.Sprintf("%d ", cur))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func solve(n int, edges [][]int, c [][]int) (int, []int) {
	deg := make([]int, n)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	var first int

	for i := 0; i < n; i++ {
		if deg[i] > 2 {
			return -1, nil
		}
		if deg[i] == 1 {
			first = i
		}
	}

	arr := make([]int, n)
	arr[0] = first

	for i := 1; i < n; i++ {
		for j := g.nodes[first]; j > 0; j = g.next[j] {
			v := g.to[j]
			if i == 1 || v != arr[i-2] {
				arr[i] = v
				first = v
				break
			}
		}
	}

	check := func(color []int) int {
		var res int

		for i := 0; i < n; i++ {
			res += c[color[i%3]][arr[i]]
		}

		return res
	}
	best := 1 << 60
	var ans []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			k := 3 - i - j
			cur := []int{i, j, k}
			tmp := check(cur)
			if tmp < best {
				best = tmp
				ans = cur
			}
		}
	}

	color := make([]int, n)

	for i := 0; i < n; i++ {
		color[arr[i]] = ans[i%3] + 1
	}

	return best, color
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
