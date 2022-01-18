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
		n := readNum(reader)
		C, _ := reader.ReadString('\n')
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, C)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Println(buf.String())
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

func getIndex(c byte) int {
	if c == 'B' {
		return 0
	} else if c == 'R' {
		return 1
	}
	return 2
}

func solve(n int, E [][]int, C string) bool {
	// B > R + G return false
	// every B is connected to an non-B
	g := NewGraph(n, len(E)*2+2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--

		if C[u] == 'G' && C[v] == 'R' || C[u] == 'R' && C[v] == 'G' {
			continue
		}

		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	var dfs func(u int) int

	dfs = func(u int) int {
		var res int
		if C[u] == 'B' {
			res++
		}
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				res += dfs(v)
			}
		}
		return res
	}

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		cnt := dfs(i)

		if cnt > 1 {
			return false
		}
	}
	return true
}

func solve1(n int, E [][]int, C string) bool {
	// B > R + G return false
	// every B is connected to an non-B
	g := NewGraph(n, len(E)*2+2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p, u int) bool

	blue := make([]bool, n)

	dfs = func(p, u int) bool {
		x := getIndex(C[u])

		if x == 0 {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				if !dfs(u, v) || blue[v] {
					return false
				}
			}
			blue[u] = true
		} else {
			// red or green
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				tmp := dfs(u, v)
				if !tmp {
					return false
				}
				y := getIndex(C[v])
				if y > 0 && x != y {
					// a safe path
					continue
				}

				if blue[v] {
					if blue[u] {
						// some child x has blue connected to v is also blue
						return false
					}
					blue[u] = true
				}
			}
		}

		return true
	}

	return dfs(-1, 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
