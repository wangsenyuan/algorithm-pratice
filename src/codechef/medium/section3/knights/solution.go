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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)
		G := make([][]byte, m)
		for i := 0; i < m; i++ {
			G[i], _ = reader.ReadBytes('\n')
		}
		res := solve(m, n, G)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(m, n int, G [][]byte) int {
	N := m * n
	g := NewGraph(N, N*N)

	connect := func(u, v int) {
		a, b := u/n, u%n
		if (a^b)&1 == 1 {
			g.AddEdge(u, v)
		} else {
			g.AddEdge(v, u)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] == '#' {
				continue
			}
			if i+1 < m && j+2 < n && G[i+1][j+2] == '.' {
				u := i*n + j
				v := (i+1)*n + j + 2
				connect(u, v)
			}
			if i+1 < m && j-2 >= 0 && G[i+1][j-2] == '.' {
				u := i*n + j
				v := (i+1)*n + j - 2
				connect(u, v)
			}
			if i+2 < m && j+1 < n && G[i+2][j+1] == '.' {
				u := i*n + j
				v := (i+2)*n + j + 1
				connect(u, v)
			}
			if i+2 < m && j-1 >= 0 && G[i+2][j-1] == '.' {
				u := i*n + j
				v := (i+2)*n + j - 1
				connect(u, v)
			}
		}
	}
	pair := make([]int, N)
	for i := 0; i < N; i++ {
		pair[i] = -1
	}
	seen := make([]int, N)
	var dfs func(u int, cur int) bool
	dfs = func(u int, cur int) bool {
		if seen[u] == cur {
			return false
		}
		seen[u] = cur
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if pair[v] < 0 || dfs(pair[v], cur) {
				pair[v] = u
				return true
			}
		}
		return false
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] == '.' {
				res++
				if (i^j)&1 == 1 && pair[i*n+j] < 0 && dfs(i*n+j, i*n+j+1) {
					res--
				}
			}

		}
	}

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
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
