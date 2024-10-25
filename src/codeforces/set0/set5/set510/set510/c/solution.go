package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	words := make([]string, n)

	for i := range n {
		words[i] = readString(reader)
	}

	res := solve(words)

	fmt.Println(res)
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

func solve(words []string) string {
	g := NewGraph(26, 26*26)

	deg := make([]int, 26)

	n := len(words)

	for i := 0; i+1 < n; i++ {
		a := words[i]
		b := words[i+1]
		var k int
		for k < len(a) && k < len(b) && a[k] == b[k] {
			k++
		}
		if k == len(a) {
			// a是b的前缀
			continue
		}
		if k == len(b) {
			return "Impossible"
		}
		x := int(a[k] - 'a')
		y := int(b[k] - 'a')
		g.AddEdge(x, y)
		deg[y]++
	}

	vis := make([]bool, 26)
	var buf []byte

	que := make([]int, 26)
	var head, tail int

	add := func(i int) {
		buf = append(buf, byte(i+'a'))
		vis[i] = true
		que[head] = i
		head++
	}

	for i := 0; i < 26; i++ {
		if deg[i] == 0 {
			add(i)
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				add(v)
			}
		}
	}

	if len(buf) != 26 {
		// 可能存在 deg > 0 的环
		return "Impossible"
	}

	return string(buf)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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
