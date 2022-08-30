package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		dominoes := make([][]int, n)
		for i := 0; i < n; i++ {
			dominoes[i] = readNNums(reader, 2)
		}
		res := solve(dominoes)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(dominoes [][]int) bool {
	n := len(dominoes)
	// 每个节点最多两条边
	g := NewGraph(n+1, 2*n)

	cnt := make([]int, n+1)

	pos := make([]int, n+1)

	for i := 0; i < n; i++ {
		domino := dominoes[i]
		if domino[0] == domino[1] {
			return false
		}
		a, b := domino[0], domino[1]
		if cnt[a] > 1 || cnt[b] > 1 {
			return false
		}
		cnt[a]++
		cnt[b]++

		if pos[a] > 0 {
			g.AddEdge(i+1, pos[a])
			g.AddEdge(pos[a], i+1)
		}
		if pos[b] > 0 {
			g.AddEdge(i+1, pos[b])
			g.AddEdge(pos[b], i+1)
		}
		pos[a] = i + 1
		pos[b] = i + 1
	}

	color := make([]int, n+1)

	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if color[u] == c {
			return true
		}

		if color[u] != 0 {
			return false
		}

		color[u] = c
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, -c) {
				return false
			}
		}
		return true
	}

	for u := 1; u <= n; u++ {
		if color[u] == 0 {
			if !dfs(u, 1) {
				return false
			}
		}
	}
	return true
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
