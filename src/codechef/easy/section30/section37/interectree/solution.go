package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(arr []int, k int) int {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", k))
		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		marked := solve(n, edges, ask)
		fmt.Printf("! %d\n", marked)
	}
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

const DI = 20

func solve(n int, edges [][]int, ask func([]int, int) int) int {
	// 二分法，检查的东西是啥？
	// 首先marked肯定在1.。。n个节点中
	// 假设第一步得到的结果是x， 这样的结果有多少个呢？
	// 这个可以通过dfs计算出来，
	// 假设 [x, y]是最后剩余的两个节点，先将x和y中间的edge删掉，
	// 然后将x中另外一个方向的（不是连接到y的）删掉一些edge，使得其最远距离小于原始结果
	g := NewGraph(n, 2*n)

	for i, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v, i+1)
		g.AddEdge(v, u, i+1)
	}

	rem := make([]int, n)

	far_dist := ask(rem, 0)

	var arr []int

	for i := 0; i < n; i++ {
		cur_height := g.height_from(i)
		if cur_height-1 == far_dist {
			arr = append(arr, i)
		}
	}

	l, r := 0, len(arr)-1
	for l < r {
		var k int
		// 将arr二分，找到一些边，使得HH[u] < far_dist
		// 这里有两个问题
		// 第一个是一条边可能被重复删除，这个好办
		// 第二个是删除前面的边，有可能让后面的的节点也变成独立的节点
		// 比如 1 - 2 - 3， 删除2的边，会让1和3也变成独立的点
		m := (l + r) / 2
		mem := make(map[int]bool)
		for i := l; i <= m; i++ {
			u := arr[i]
			for j := g.node[u]; j > 0; j = g.next[j] {
				if !mem[g.label[j]] {
					rem[k] = g.label[j]
					k++
					mem[g.label[j]] = true
				}
			}
		}

		far_dist = ask(rem, k)

		if far_dist == 0 {
			r = m
		} else {
			l = m + 1
		}
	}

	return arr[r] + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node  []int
	next  []int
	to    []int
	label []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.label = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, l int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.label[g.cur] = l
}

func (g *Graph) height_from(root int) int {

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		var height int
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				height = max(height, dfs(u, v))
			}
		}
		height++
		return height
	}

	return dfs(-1, root)
}
