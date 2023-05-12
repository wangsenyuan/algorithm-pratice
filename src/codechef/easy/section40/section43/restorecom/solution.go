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
		F := make([][]int, n)
		for i := 0; i < n; i++ {
			var k int
			s, _ := reader.ReadBytes('\n')
			pos := readInt(s, 0, &k)
			F[i] = make([]int, k)
			for j := 0; j < k; j++ {
				pos = readInt(s, pos+1, &F[i][j])
			}
		}
		res := solve(n, F)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, F [][]int) int {
	id := n
	ii := make(map[int]int)
	var edges int
	for _, cur := range F {
		edges += len(cur)
		for _, x := range cur {
			if ii[x] == 0 {
				ii[x] = id
				id++
			}
		}
	}
	N := id

	if edges < N-1 {
		return -1
	}

	g := NewGraph(N, edges*2)
	for i, cur := range F {
		for _, x := range cur {
			j := ii[x]
			g.AddEdge(i, j)
			g.AddEdge(j, i)
		}
	}
	mark := make([]bool, N)
	var comps int

	que := make([]int, N)

	for i := 0; i < N; i++ {
		if mark[i] {
			continue
		}
		comps++
		mark[i] = true
		var front, end int
		que[end] = i
		end++
		for front < end {
			u := que[front]
			front++
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				if !mark[v] {
					mark[v] = true
					que[end] = v
					end++
				}
			}
		}
	}
	return comps - 1
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
