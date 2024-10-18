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
	sets := make([][]int, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		sets[i] = make([]int, k)
		for j := 0; j < k; j++ {
			pos = readInt(s, pos, &sets[i][j]) + 1
		}
	}

	ok, res := solve(m, sets)

	if !ok {
		fmt.Println("No")
		return
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Yes\n%d\n", len(res)))

	if len(res) > 0 {
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(m int, words [][]int) (bool, []int) {
	m2 := m * 2
	var sum int
	for _, w := range words {
		sum += len(w)
		for i := range w {
			w[i]--
		}
	}
	g := NewGraph(m2, sum*5)

	connect := func(u int, v int) {
		g.AddEdge(u, v)
		g.AddEdge(v^1, u^1)
	}

	// 不大写为 2 * num, 大写为 2 * num + 1
	for i := 0; i+1 < len(words); i++ {
		a := words[i]
		b := words[i+1]
		// 只有在保证相等的情况下才能继续
		eq := true
		for j := 0; j < len(a) && j < len(b); j++ {
			if a[j] <= b[j] {
				// 那么大写b[j]的时候，也必须大写a[j]
				connect(2*b[j]+1, 2*a[j]+1)
			} else {
				// 如果 a[j] > b[j]那么，必须大写a[j](才能保证a < b), 且b[j]不能大写
				connect(2*a[j], 2*a[j]+1)
				connect(2*b[j]+1, 2*b[j])
			}
			if a[j] != b[j] {
				eq = false
				break
			}
		}

		if eq && len(a) > len(b) {
			return false, nil
		}
	}

	low := make([]int, m2)
	dsc := make([]int, m2)
	vis := make([]int, m2)
	comp := make([]int, m2)
	var cid int
	stack := make([]int, m2)
	var top int
	var time int

	var dfs func(u int)
	dfs = func(u int) {
		dsc[u] = time
		low[u] = time
		time++
		stack[top] = u
		top++
		vis[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 2 {
				continue
			}
			if vis[v] == 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], dsc[v])
			}
		}

		if low[u] == dsc[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				comp[v] = cid
				if u == v {
					break
				}
			}
			cid++
		}
		vis[u]++
	}

	for i := 0; i < m2; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	var ans []int

	for i := 0; i < m; i++ {
		if comp[2*i] == comp[2*i+1] {
			return false, nil
		}
		if comp[2*i] > comp[2*i+1] {
			// ssc是从树的底层开始形成的
			// c_id越小，越在更底层,
			ans = append(ans, i+1)
		}
	}

	return true, ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
