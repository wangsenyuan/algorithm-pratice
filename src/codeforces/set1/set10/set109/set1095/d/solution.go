package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, 2)
	}

	res := solve(n, a)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, a [][]int) []int {
	// a[i][0] and a[i][1] 有可能是顺序反的
	// let a[i][0] = j, a[i][1] = k
	// 假设pos[i], pos[j], pos[k] 表示三个的位置
	// 那么 pos[i] => pos[j] and pos[k]
	// 且 pos[j] and pos[k] 肯定也是有关系的， 而且答案存在的情况下，它们的位置是b固定的
	if n == 3 {
		// 答案存在时，这个位置肯定满足要求
		return []int{1, 2, 3}
	}
	// directed graph
	g := NewGraph(n+1, 2*n)
	deg := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var pt1 bool
		for _, j := range a[i-1] {
			if j > 1 {
				g.AddEdge(i, j)
				deg[j]++
			} else {
				pt1 = true
			}
		}

		if pt1 {
			// 它是1的前置节点
			k := a[i-1][0] ^ a[i-1][1] ^ 1
			if a[0][0] == k || a[0][1] == k {
				deg[k]--
			}
		}
	}

	que := make([]int, n)
	var front, tail int
	que[front] = 1
	front++

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	return que
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, 3+e)
	to := make([]int, 3+e)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
