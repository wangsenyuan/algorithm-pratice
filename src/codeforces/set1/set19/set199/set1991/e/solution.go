package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--

		n, m := readTwoNums(reader)
		edges := make([][]int, m)

		for i := range m {
			edges[i] = readNNums(reader, 2)
		}

		play := func(role string, x int, y int, cnt int) (int, int) {
			if cnt == 0 {
				fmt.Println(role)
			}
			if role == "Alice" {
				fmt.Printf("%d %d\n", x, y)
				// alice 的情况下，juedge始终会有返回
				return readTwoNums(reader)
			}
			// role = Bob
			if cnt > 0 {
				fmt.Printf("%d %d\n", x, y)
			}
			if cnt < n {
				return readTwoNums(reader)
			}
			return -1, -1
		}

		solve(n, edges, play)
	}

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

func solve(n int, edges [][]int, play func(string, int, int, int) (int, int)) {
	g := NewGraph(n, 2*len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var dfs func(p int, u int, c int) bool

	dfs = func(p int, u int, c int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				ok := dfs(u, v, c^1)
				if !ok {
					return false
				}
			}
		}
		return true
	}

	bi := dfs(-1, 0, 0)

	if !bi {
		// 这种情况下，1， 2肯定能让alice获胜
		for i := 0; i < n; i++ {
			play("Alice", 1, 2, i)
		}
		return
	}
	// BOB这个怎么处理呢？
	arr := make([][]int, 2)
	for i := 0; i < n; i++ {
		arr[color[i]] = append(arr[color[i]], i+1)
	}
	id, c := -1, -1
	for i := 0; i <= n; i++ {
		a, b := play("Bob", id, c, i)
		if i < n {

			if (a == 1 || b == 1) && len(arr[0]) > 0 {
				id, c = arr[0][0], 1
				arr[0] = arr[0][1:]
				continue
			}
			// 要么没有颜色1 (2, 3)
			if (a == 2 || b == 2) && len(arr[1]) > 0 {
				id, c = arr[1][0], 2
				arr[1] = arr[1][1:]
				continue
			}
			// 第三种情况， 要么 arr[0] 用完了，要么arr[1]用完了 (如果没有用完，a, b 肯定会有1或者2)
			// 假设 arr[0]用完了
			// 一定会有颜色3吗？假设是(1, 2)那么上面肯定有一个条件是满足的，所以不会到这里
			// 所以肯定有颜色3
			if len(arr[0]) > 0 {
				id, c = arr[0][0], 3
				arr[0] = arr[0][1:]
			} else {
				id, c = arr[1][0], 3
				arr[1] = arr[1][1:]
			}
		}
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
