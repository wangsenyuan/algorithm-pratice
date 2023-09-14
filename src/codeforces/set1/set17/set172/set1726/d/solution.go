package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(n int, E [][]int) string {
	g := NewGraph(n, len(E)*2)
	for i, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	vis := make([]bool, n)

	track := make([]int, n)

	buf := make([]byte, len(E))
	for i := 0; i < len(E); i++ {
		buf[i] = '0'
	}

	var dfs func(u int)

	dep := make([]int, n)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				buf[g.val[i]] = '1'
				track[v] = u
				dep[v] = dep[u] + 1
				dfs(v)
			}
		}
	}

	dfs(0)

	blue := make(map[int]int)
	for i, e := range E {
		if buf[i] == '0' {
			u, v := e[0]-1, e[1]-1
			blue[u]++
			blue[v]++
		}
	}

	if len(blue) != 3 {
		return string(buf)
	}
	arr := make([][]int, 0, 3)
	for v, d := range blue {
		arr = append(arr, []int{v, d})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	if arr[0][1] < arr[2][1] || arr[0][1] != 2 {
		return string(buf)
	}
	// need to improve
	for i := 0; i < 3; i++ {
		v := arr[i][0]
		arr[i][1] = dep[v]
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	u := arr[0][0]

	var first, second int
	for i := g.nodes[u]; i > 0; i = g.next[i] {
		v := g.to[i]
		j := g.val[i]
		if buf[j] == '0' {
			first = j
		} else if v == track[u] {
			second = j
		}
	}
	buf[first], buf[second] = buf[second], buf[first]
	return string(buf)
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
