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
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		ok, res := solve(n, edges)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteString("\n")
		}
	}

	fmt.Print(buf.String())
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

func solve(n int, edges [][]int) (bool, []int) {
	if n%3 != 0 {
		return false, nil
	}

	g := NewGraph(n, 2*n)

	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	var res []int

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		cnt := []int{0, 0}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v)
				if tmp < 0 {
					// unable to process
					return -1
				}
				if tmp == 0 {
					// v is already good
					res = append(res, g.val[i]+1)
					continue
				}
				cnt[tmp-1]++
			}
		}
		if cnt[1] > 1 {
			// too many children need this one
			return -1
		}
		if cnt[1] == 1 {
			// only it need
			if cnt[0] > 0 {
				// conflicts
				return -1
			}
			// good to use u
			return 0
		}
		// cnt[1] = 0
		if cnt[0] == 0 {
			// u is left
			return 1
		}
		if cnt[0] == 1 {
			return 2
		}
		if cnt[0] == 2 {
			// pair u up
			return 0
		}
		return -1
	}

	if dfs(-1, 0) < 0 {
		return false, nil
	}

	return true, res
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
