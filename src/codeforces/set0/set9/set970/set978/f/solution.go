package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	pairs := make([][]int, m)
	for i := range m {
		pairs[i] = readNNums(reader, 2)
	}
	return solve(a, pairs)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, pairs [][]int) []int {
	n := len(a)

	users := make([]pair, n)
	for i := 0; i < n; i++ {
		users[i] = pair{a[i], i}
	}

	slices.SortFunc(users, func(x, y pair) int {
		return x.first - y.first
	})

	ans := make([]int, n)

	marked := make([]bool, n)

	g := NewGraph(n, 2*len(pairs))

	for _, cur := range pairs {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	for i := 0; i < n; {
		j := i
		for i < n && users[i].first == users[j].first {
			u := users[i].second
			// 前j个比他高分的人
			ans[u] = j
			for k := g.nodes[u]; k > 0; k = g.next[k] {
				v := g.to[k]
				if marked[v] {
					ans[u]--
				}
			}
			i++
		}

		for j < i {
			marked[users[j].second] = true
			j++
		}
	}

	return ans
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
