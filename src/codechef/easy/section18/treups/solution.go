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
		n, q := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, q)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
func solve(n int, E [][]int, q int) int {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var V []int
	var dfs func(p, u, level int)
	dfs = func(p, u, level int) {
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				dfs(u, v, level^1)
			}
		}

		if cnt != 1 && level == 0 {
			V = append(V, abs(cnt-1))
		}
	}
	dfs(-1, 0, 0)
	var sum int
	for _, v := range V {
		sum += v
	}
	if q == 1 {
		return sum
	}
	// knapsack with compressed form
	freq := make(map[int]int)
	for _, v := range V {
		freq[v]++
	}
	m := len(freq)
	a := make([]int, 0, m)
	b := make([]int, 0, m)
	for k, v := range freq {
		a = append(a, k)
		b = append(b, v)
	}

	// dp[i][j] = 第i项时，获得和为j，所需要的 a[i] 的 数量 (<= b[i])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, sum+1)
		for j := 1; j <= sum; j++ {
			dp[i][j] = -1
		}
		//dp[i][0] = 0
	}

	for j := a[0]; j <= sum; j++ {
		if j%a[0] == 0 && j/a[0] <= b[0] {
			dp[0][j] = j / a[0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j <= sum; j++ {
			if dp[i-1][j] >= 0 {
				//前i-1项已经获得了sum j
				dp[i][j] = 0
				continue
			}
			if j >= a[i] {
				if dp[i][j-a[i]] >= 0 && dp[i][j-a[i]] < b[i] {
					dp[i][j] = dp[i][j-a[i]] + 1
				}
			}
		}
	}

	ans := sum

	for j := 0; j <= sum; j++ {
		if dp[m-1][j] >= 0 {
			ans = min(ans, abs(sum-2*j))
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
