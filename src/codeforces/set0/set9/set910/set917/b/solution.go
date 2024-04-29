package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	edges := make([][]int, m)
	letters := make([]byte, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var pos int
		edges[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			pos = readInt(s, pos, &edges[i][j]) + 1
		}
		letters[i] = s[pos]
	}

	res := solve(n, edges, letters)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(res[i])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(n int, edges [][]int, letters []byte) []string {

	g := NewGraph(n, len(edges))

	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, int(letters[i]-'a')+1)
	}

	dp := make([][][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([][]int, 27)
			for k := 0; k < 27; k++ {
				dp[i][j][k] = make([]int, 2)
			}
		}
	}

	// 可以用 dfs吗？

	var dfs func(a, b, x, d int)

	dfs = func(a, b, x, d int) {
		// x = 0, 表示一开始的位置就在(a, b)上
		if dp[a][b][x][d] != 0 {
			// already processed
			return
		}
		if a == b {
			// 另外一方可以跟着当前的路径
			if d == 0 {
				dp[a][b][x][d] = -1
			} else {
				dp[a][b][x][d] = 1
			}
			return
		}
		u := a
		// 当前如果由alice move， 假设 bob 能获胜
		// 然后看alice是否有获胜的机会
		dp[a][b][x][d] = -1
		if d == 1 {
			// 轮到bob了
			u = b
			dp[a][b][x][d] = 1
		}
		// 需要知道下次move的状态
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if w >= x {
				if d == 0 {
					dfs(v, b, w, 1)
					if dp[v][b][w][1] == 1 {
						// 到下一个位置后，alice可以获胜
						dp[a][b][x][0] = 1
					}
				} else {
					dfs(a, v, w, 0)
					if dp[a][v][w][0] == -1 {
						dp[a][b][x][1] = -1
					}
				}
			}
		}
	}
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			buf[i][j] = 'A'
			dfs(i, j, 0, 0)
			if dp[i][j][0][0] == -1 {
				buf[i][j] = 'B'
			}
		}
	}

	res := make([]string, n)

	for i := 0; i < n; i++ {
		res[i] = string(buf[i])
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
