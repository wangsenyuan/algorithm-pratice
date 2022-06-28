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
		S := readNNums(reader, n)
		T := readNNums(reader, n)
		res := solve(n, S, T)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, S []int, T []int) string {
	g := NewGraph(n, 2*n)

	for i := 0; i+1 < n; i += 2 {
		g.AddEdge(S[i], S[i+1])
		g.AddEdge(S[i+1], S[i])
		g.AddEdge(T[i], T[i+1])
		g.AddEdge(T[i+1], T[i])
	}

	assign := make([]int, n)
	for i := 0; i < n; i++ {
		assign[i] = -1
	}
	que := make([]int, n)
	var front, end int
	for i := 0; i < n; i++ {
		if assign[i] >= 0 {
			continue
		}
		assign[i] = 0
		que[end] = i
		end++
		for front < end {
			u := que[front]
			front++
			for j := g.node[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				if assign[v] < 0 {
					assign[v] = 1 - assign[u]
					que[end] = v
					end++
				}
			}
		}
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		if assign[i] == 0 {
			buf[i] = 'A'
		} else {
			buf[i] = 'B'
		}
	}
	return string(buf)
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
