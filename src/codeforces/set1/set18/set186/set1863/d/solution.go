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
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)
		}
		res := solve(n, m, grid)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(res[i])
				buf.WriteByte('\n')
			}
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

func solve(n int, m int, grid []string) []string {

	// do have result
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			buf[i][j] = '.'
		}
	}

	for i := 0; i < n; i++ {
		var cnt int
		for j := 0; j < m; j++ {
			if grid[i][j] == 'U' {
				if cnt%2 == 0 {
					buf[i][j] = 'B'
					buf[i+1][j] = 'W'
				} else {
					buf[i][j] = 'W'
					buf[i+1][j] = 'B'
				}
				cnt++
			}
		}
		if cnt&1 == 1 {
			return nil
		}
	}

	for j := 0; j < m; j++ {
		var cnt int
		for i := 0; i < n; i++ {
			if grid[i][j] == 'L' {
				if cnt%2 == 0 {
					buf[i][j] = 'B'
					buf[i][j+1] = 'W'
				} else {
					buf[i][j] = 'W'
					buf[i][j+1] = 'B'
				}
				cnt++
			}
		}
		if cnt&1 == 1 {
			return nil
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
