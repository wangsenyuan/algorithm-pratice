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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		k := readNum(reader)
		Q := make([]int, k)
		for i := 0; i < k; i++ {
			Q[i] = readNum(reader)
		}
		res := solve(n, E, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
		}
	}

	fmt.Print(buf.String())
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

func solve(n int, E [][]int, Q []int) [][]int {
	// 对于x, 找到u,v, x位于 simple-path u -> v
	// 且 and(path) > 0, and maximize count of nodes in path
	// dp[u][i] 表示在以u为根节点的子树中，经过u的path weight and sum [i] = 1的最长的路径
	// fp[u][i] = 从u开始到某个子节点，and sum[i] = 1的最长路径
	// ans[x] = dp[x][?]
	// 然后开始往上延伸，
	// 对于每一位，需要查找该位最高到某个节点p，
	// 但是没有办法从x一直往上找，这样潜在的是O[n]
	// 可以按位去生成树 （森林）
	// 然后计算每个点在这棵树中，最远的点（这个应该是可以在O(n)时间内算出的),然后对于x，就是在这30颗树中的最大值

	ans := make([][][]int, H)
	for i := 0; i < H; i++ {
		ans[i] = handleTree(n, E, i)
	}

	query := func(x int) []int {
		x--
		var best int
		for i := 0; i < H; i++ {
			best = max(best, ans[i][x][0])
		}
		if best <= 1 {
			// no answer
			return []int{-1, -1}
		}

		for i := 0; i < H; i++ {
			if ans[i][x][0] == best {
				u, v := ans[i][x][1], ans[i][x][2]
				return []int{u + 1, v + 1}
			}
		}
		return []int{-1, -1}
	}

	res := make([][]int, len(Q))

	for i, x := range Q {
		res[i] = query(x)
	}

	return res
}

const H = 30

func handleTree(n int, E [][]int, bit int) [][]int {
	deg := make([]int, n)

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if (w>>bit)&1 == 1 {
			deg[u]++
			deg[v]++
			g.AddEdge(u, v, w)
			g.AddEdge(v, u, w)
		}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = []int{-1, -1, -1}
	}

	height := make([]Pair, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		height[u] = Pair{0, u}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if height[v].first > height[u].first {
					height[u] = height[v]
				}
			}
		}
		height[u] = Pair{height[u].first + 1, height[u].second}
	}

	var dfs2 func(p int, u int, cur Pair)
	dfs2 = func(p int, u int, cur Pair) {
		first, second := -1, -1

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if first < 0 || height[first].first <= height[v].first {
					second = first
					first = v
				} else if second < 0 || height[second].first <= height[v].first {
					second = v
				}
			}
		}

		ans[u] = []int{cur.first + height[u].first, cur.second, height[u].second}

		if first < 0 {
			return
		}

		if second < 0 {
			dfs2(u, first, cur.AddFirst(1))
			return
		}
		// 或者是两条子边的长度 + 自己
		if height[first].first+height[second].first+1 > ans[u][0] {
			ans[u] = []int{height[first].first + height[second].first + 1, height[first].second, height[second].second}
		}
		// second >= 0
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if v == first {
				dfs2(u, v, cur.Max(height[second]).AddFirst(1))
			} else {
				dfs2(u, v, cur.Max(height[first]).AddFirst(1))
			}
		}
	}

	for i := 0; i < n; i++ {
		if ans[i][0] < 0 && deg[i] == 1 {
			dfs(-1, i)
			dfs2(-1, i, Pair{0, i})
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Max(that Pair) Pair {
	if this.first > that.first {
		return this
	}
	return that
}

func (this Pair) AddFirst(v int) Pair {
	return Pair{this.first + v, this.second}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
