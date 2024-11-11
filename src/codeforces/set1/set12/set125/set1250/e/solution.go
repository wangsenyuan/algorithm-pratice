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
		n, m, k := readThreeNums(reader)
		arr := make([]string, n)
		for i := range n {
			arr[i] = readString(reader)[:m]
		}
		cnt, res := solve(arr, k)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))
		if cnt < 0 {
			continue
		}
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
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

func solve(necklaces []string, k int) (int, []int) {
	n := len(necklaces)
	g := NewGraph(n, n*n)

	connect := func(u int, v int, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := check(necklaces[i], necklaces[j], k)
			y := check(necklaces[i], reverse(necklaces[j]), k)

			if !x && !y {
				return -1, nil
			}
			if x && y {
				// 都可以，没有限制
				continue
			}
			if x {
				connect(i, j, 0)
			} else {
				connect(i, j, 1)
			}
		}
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var ans []int

	var dfs func(u int, c int) bool

	cnt := make([]int, 2)
	var comp []int

	dfs = func(u int, c int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}

		color[u] = c
		cnt[c]++
		comp = append(comp, u)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if !dfs(v, w^c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if color[i] >= 0 {
			continue
		}
		clear(cnt)
		comp = comp[:0]
		if !dfs(i, 0) {
			return -1, nil
		}
		add := 0
		if cnt[1] < cnt[0] {
			add = 1
		}
		for _, u := range comp {
			if color[u] == add {
				ans = append(ans, u+1)
			}
		}
	}

	return len(ans), ans
}

func check(a string, b string, k int) bool {
	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			cnt++
		}
	}
	return cnt >= k
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
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
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.val = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
