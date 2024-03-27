package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	locked := readNNums(reader, n)
	lines := make([][]int, m)
	for i := 0; i < m; i++ {
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k) + 1
		lines[i] = make([]int, k)
		for j := 0; j < k; j++ {
			pos = readInt(s, pos, &lines[i][j]) + 1
		}
	}
	res := solve(n, locked, m, lines)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(n int, locked []int, m int, lines [][]int) bool {
	// 0 表示 locked， 1 表示unlocked
	var tot int
	for _, ls := range lines {
		tot += len(ls)
	}

	first := make([]int, n)
	for i := 0; i < n; i++ {
		first[i] = -1
	}
	g := NewGraph(m, 2*n)

	for i, ls := range lines {
		for j := 0; j < len(ls); j++ {
			r := ls[j] - 1
			if first[r] < 0 {
				first[r] = i
			} else {
				g.AddEdge(i, first[r], 1^locked[r])
				g.AddEdge(first[r], i, 1^locked[r])
			}
		}
	}
	color := make([]int, m)

	for i := 0; i < m; i++ {
		color[i] = -1
	}

	var dfs func(u int, expect int) bool
	dfs = func(u int, expect int) bool {
		if color[u] >= 0 {
			return color[u] == expect
		}
		color[u] = expect
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, expect^g.val[i]) {
				return false
			}
		}
		return true
	}

	for i := 0; i < m; i++ {
		if color[i] < 0 && !dfs(i, 0) {
			return false
		}
	}

	return true
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
