package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
	s := fmt.Sprintf("%v", res)
	s = s[1 : len(s)-1]
	fmt.Println(s)
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

func solve(n int, E [][]int) []int {
	// ans[1] = 1 (它时一个完全图）
	// ans[2] = 如果只有3个节点，且长度为3，那么就分成了2部分
	// 找到最长的路径的两个端点(x, y)
	// 如果k > dist(x, y), 那么所有的点都断开了
	// 需要知道，每一个点离他最远的距离，然后当k > dist 它就孤立了
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([][]Pair, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dist[u] = make([]Pair, 2)
		dist[u][0] = Pair{0, u}
		dist[u][1] = Pair{0, u}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if dist[v][0].first+1 >= dist[u][0].first {
					dist[u][1] = dist[u][0]
					dist[u][0] = Pair{dist[v][0].first + 1, v}
				} else if dist[v][0].first+1 >= dist[u][1].first {
					dist[u][1] = Pair{dist[v][0].first + 1, v}
				}
			}
		}
	}

	dfs(0, 0)

	var dfs2 func(p int, u int, ln int)

	dfs2 = func(p int, u int, ln int) {
		if p >= 0 {
			if ln+1 >= dist[u][0].first {
				dist[u][1] = dist[u][0]
				dist[u][0] = Pair{ln + 1, p}
			} else if ln+1 >= dist[u][1].first {
				dist[u][1] = Pair{ln + 1, p}
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if v == dist[u][0].second {
					dfs2(u, v, dist[u][1].first)
				} else {
					dfs2(u, v, dist[u][0].first)
				}
			}
		}
	}

	dfs2(-1, 0, 0)

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return dist[id[i]][0].first < dist[id[j]][0].first
	})
	ans := make([]int, n)
	ans[0] = 1
	for k, i := 1, 0; k < n; k++ {
		ans[k] = ans[k-1]
		for i < n && dist[id[i]][0].first == k {
			ans[k]++
			i++
		}
		if ans[k] > n {
			ans[k] = n
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
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
