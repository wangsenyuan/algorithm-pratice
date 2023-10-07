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
		b := make([]string, n)
		for i := 0; i < n; i++ {
			b[i] = readString(reader)
		}
		res := solve(b)
		for _, row := range res {
			buf.WriteString(fmt.Sprintf("%d", len(row)))
			for i := 0; i < len(row); i++ {
				buf.WriteString(fmt.Sprintf(" %d", row[i]))
			}
			buf.WriteByte('\n')
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

func solve(b []string) [][]int {
	// b[i][j] = 1, a[i] 是 a[j] 的一个子集
	// 那这样子，可以搞出一个图出来
	// 然后topo排序
	n := len(b)
	g := NewGraph(n, n*n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] == '1' {
				g.AddEdge(j, i)
			}
		}
	}
	vis := make([]bool, n)

	var ord []int
	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		ord = append(ord, u)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}

	//reverse(ord)

	it := 1

	ans := make([][]int, n)

	cnt := make([]int, n+1)
	for _, u := range ord {
		var tmp []int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			// v is subset of u
			// 还要去重
			for _, x := range ans[v] {
				cnt[x]++
				if cnt[x] == 1 {
					tmp = append(tmp, x)
				}
			}
		}
		tmp = append(tmp, it)
		for _, x := range tmp {
			cnt[x] = 0
		}
		ans[u] = tmp
		it++
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
