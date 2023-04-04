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
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = readNNums(reader, m)
	}
	res := solve(P)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, row := range res {
			for i := 0; i < 3; i++ {
				buf.WriteString(fmt.Sprintf("%d ", row[i]))
			}
			buf.WriteByte('\n')
		}
		fmt.Print(buf.String())
	}
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

func solve(P [][]int) [][]int {
	n := len(P)
	m := len(P[0])
	vis := make([][]bool, n)
	special := make([][]bool, n)
	for i := 0; i < n; i++ {
		special[i] = make([]bool, m)
		vis[i] = make([]bool, m)
	}

	que := make([]int, n*m)

	var front, end int

	var res [][]int

	findColor := func(x int, y int) int {
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				if !special[x+i][y+j] {
					return P[x+i][y+j]
				}
			}
		}
		return -1
	}

	pushBack := func(i int, j int) {
		if vis[i][j] {
			return
		}
		vis[i][j] = true

		c := findColor(i, j)

		res = append(res, []int{i + 1, j + 1, c})

		u := i*m + j
		que[end] = u
		end++
	}

	popFront := func() (int, int) {
		u := que[front]
		front++
		return u / m, u % m
	}

	for i := 0; i+1 < n; i++ {
		for j := 0; j+1 < m; j++ {
			if P[i][j] == P[i+1][j] && P[i][j] == P[i][j+1] && P[i][j] == P[i+1][j+1] {
				pushBack(i, j)
			}
		}
	}

	canColor := func(x, y int) bool {
		c := findColor(x, y)
		if c < 0 {
			return false
		}
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				if !special[x+i][y+j] && P[x+i][y+j] != c {
					return false
				}
			}
		}
		return true
	}

	if front == end {
		return nil
	}

	for front < end {
		x, y := popFront()
		// x, y
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				special[i+x][j+y] = true
			}
		}

		// 9 positions
		// x - 1, y - 1, x - 1, y, x - 1, y + 1
		// x, y - 1, x, y, x, y + 1
		// x + 1, y - 1,
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				u, v := x+dx, y+dy
				if u >= 0 && u < n-1 && v >= 0 && v < m-1 && !vis[u][v] && canColor(u, v) {
					pushBack(u, v)
				}
			}
		}
	}

	// 所有的都需要变成special

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !special[i][j] {
				return nil
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
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
