package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m := readNum(reader)
	arrs := make([][]int, m)
	for i := 0; i < m; i++ {
		n := readNum(reader)
		arrs[i] = readNNums(reader, n)
	}

	ok, res := solve(arrs)

	if !ok {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer

	buf.WriteString("YES\n")

	for _, r := range res {
		buf.WriteString(r)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(arrs [][]int) (bool, []string) {
	cnt := make(map[int]int)
	id := make(map[int]int)
	m := len(arrs)
	var cur int
	var edgeCount int
	for _, arr := range arrs {
		for _, x := range arr {
			cnt[x]++
			if _, ok := id[x]; !ok {
				id[x] = cur
				cur++
			}
			edgeCount++
		}
	}

	for _, v := range cnt {
		if v&1 == 1 {
			return false, nil
		}
	}
	n := m + len(id)
	g := NewGraph(n, 2*edgeCount)

	for i, arr := range arrs {
		for k, x := range arr {
			j := id[x]
			g.AddEdge(i, j+m, k)
			g.AddEdge(j+m, i, k)
		}
	}

	var dfs func(v int)

	ans := make([][]byte, m)

	for i := 0; i < m; i++ {
		ans[i] = make([]byte, len(arrs[i]))
		for j := 0; j < len(ans[i]); j++ {
			ans[i][j] = 'L'
		}
	}

	pos := make([]int, n)

	for i := 0; i < n; i++ {
		pos[i] = g.nodes[i]
	}

	used := make([]bool, g.cur+1)

	dfs = func(v int) {
		for pos[v] > 0 {
			i := pos[v]
			pos[v] = g.next[i]
			if i^1 <= g.cur && used[i^1] {
				continue
			}
			used[i] = true
			k := g.val[i]
			if v < m {
				ans[v][k] = 'R'
			}
			u := g.to[i]
			dfs(u)
		}
	}

	for i := 0; i < n; i++ {
		dfs(i)
	}

	res := make([]string, m)
	for i := 0; i < m; i++ {
		res[i] = string(ans[i])
	}

	return true, res
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
	next := make([]int, e+2)
	to := make([]int, e+2)
	val := make([]int, e+2)
	return &Graph{nodes, next, to, val, 1}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
