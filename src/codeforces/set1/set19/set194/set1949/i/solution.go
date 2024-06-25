package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	disks := make([][]int, n)
	for i := 0; i < n; i++ {
		disks[i] = readNNums(reader, 3)
	}
	res := solve(disks)
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

func solve(disks [][]int) bool {
	n := len(disks)

	var edge_cnt int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if tangent(disks[i], disks[j]) {
				edge_cnt++
			}
		}
	}

	g := NewGraph(n, edge_cnt*2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if tangent(disks[i], disks[j]) {
				g.AddEdge(i, j)
				g.AddEdge(j, i)
			}
		}
	}

	color := make([]int, n)

	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var dfs func(u int, c int, cnt []int) bool
	// 知道了，不能中途退出，否则， 那些没有被访问到的点，会被重新跑一遍，进而得到错误的答案
	dfs = func(u int, c int, cnt []int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		cnt[c]++
		res := true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, 1^c, cnt) {
				res = false
			}
		}
		return res
	}

	for i := 0; i < n; i++ {
		if color[i] < 0 {
			cnt := make([]int, 2)
			if dfs(i, 0, cnt) && cnt[0] != cnt[1] {
				return true
			}
		}
	}

	return false
}

func tangent(a, b []int) bool {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	dr := b[2] + a[2]
	return dx*dx+dy*dy == dr*dr
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
