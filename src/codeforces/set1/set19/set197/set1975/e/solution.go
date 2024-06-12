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
		c := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		queries := make([]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNum(reader)
		}
		res := solve(n, c, edges, queries)

		for _, x := range res {
			if x {
				buf.WriteString("Yes\n")
			} else {
				buf.WriteString("No\n")
			}
		}
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

type Pair struct {
	first  int
	second int
}

func solve(n int, color []int, edges [][]int, queries []int) []bool {
	g := NewGraph(n+1, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	c := make([]int, n+1)
	copy(c[1:], color)
	color = c

	var faw, sum_two, sum_more, black_cnt, xor_two int

	fa := make([]int, n+1)
	black_child_cnt := make([]int, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		if color[u] == 1 {
			black_cnt++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if color[v] == 1 {
					black_child_cnt[u]++
				}
			}
		}

		if color[p] == 0 && color[u] == 1 {
			faw++
		}
		if color[u] == 1 {
			if black_child_cnt[u] == 2 {
				sum_two++
				xor_two ^= u
			} else if black_child_cnt[u] > 2 {
				sum_more++
			}
		}
	}

	dfs(0, 1)

	flip := func(x int) {
		color[x] ^= 1

		var d int

		if color[x] == 0 {
			d = -1
		} else {
			d = 1
		}
		black_cnt += d

		if color[fa[x]] == 0 {
			faw += d
		}

		if black_child_cnt[x] == 2 {
			sum_two += d
			xor_two ^= x
		}
		if black_child_cnt[x] > 2 {
			sum_more += d
		}
		faw -= d * black_child_cnt[x]

		if color[x] == 1 {
			if color[fa[x]] == 1 && black_child_cnt[fa[x]] == 2 {
				sum_two--
				sum_more++
				xor_two ^= fa[x]
			}
			black_child_cnt[fa[x]]++

			if color[fa[x]] == 1 && black_child_cnt[fa[x]] == 2 {
				sum_two++
				xor_two ^= fa[x]
			}
		} else {
			if color[fa[x]] == 1 && black_child_cnt[fa[x]] == 2 {
				sum_two--
				xor_two ^= fa[x]
			}
			black_child_cnt[fa[x]]--
			if color[fa[x]] == 1 && black_child_cnt[fa[x]] == 2 {
				sum_two++
				sum_more--
				xor_two ^= fa[x]
			}
		}
	}

	check := func() bool {
		if black_cnt == 0 {
			return false
		}
		if sum_more > 0 || sum_two > 1 {
			return false
		}
		if faw > 1 {
			return false
		}
		if sum_two == 1 && color[fa[xor_two]] == 1 {
			return false
		}
		return true
	}

	ans := make([]bool, len(queries))

	for i, u := range queries {
		flip(u)
		ans[i] = check()
	}

	return ans
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
