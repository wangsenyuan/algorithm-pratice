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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		L := make([][]int, n)
		for i := 0; i < n; i++ {
			L[i] = readNNums(reader, 2)
		}
		res := solve(n, L)
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

const MAX_X = 200010

func solve(n int, segs [][]int) int {
	ps := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		seg := segs[i]
		u, v := seg[0], seg[1]
		ps = append(ps, u)
		ps = append(ps, v)
	}
	sort.Ints(ps)
	ii := make(map[int]int)
	var m int
	for i := 1; i <= len(ps); i++ {
		if i == len(ps) || ps[i] > ps[i-1] {
			ii[ps[i-1]] = m
			m++
		}
	}

	rg := make([][]int, m)

	for i := 0; i < m; i++ {
		rg[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		cur := segs[i]
		u, v := cur[0], cur[1]
		j, k := ii[u], ii[v]
		rg[j] = append(rg[j], k)
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
		sort.Ints(rg[i])
	}

	var count func(l, r int) int

	count = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] >= 0 {
			return dp[l][r]
		}
		v := rg[l]
		j := sort.SearchInts(v, r)
		var add int
		if j < len(v) && v[j] == r {
			add++
		}
		dp[l][r] = count(l+1, r)
		for i := 0; i < len(v); i++ {
			if v[i] >= r {
				break
			}
			dp[l][r] = max(dp[l][r], count(l, v[i])+count(v[i]+1, r))
		}
		dp[l][r] += add

		return dp[l][r]
	}

	return count(0, m-1)
}

func solve1(n int, segs [][]int) int {
	lines := make([]Line, n, n+1)
	for i := 0; i < n; i++ {
		lines[i] = Line{segs[i][0], segs[i][1]}
	}
	lines = append(lines, Line{0, MAX_X})
	sort.Sort(Lines(lines))
	n++

	G := make([][]int, n)

	for i := 0; i < n; i++ {
		G[i] = make([]int, 0, 10)
	}

	check := func(i, j int) bool {
		a, b := lines[i], lines[j]
		return a.l <= b.l && b.r <= a.r
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if check(i, j) {
				G[i] = append(G[i], j)
			}
		}
	}

	dp := make([]int, n)
	var dfs func(u int) int

	dfs = func(u int) int {
		if dp[u] != 0 {
			return dp[u]
		}
		if len(G[u]) == 0 {
			dp[u] = 1
			return 1
		}
		var root *Node
		var best int
		for _, v := range G[u] {
			tmp := dfs(v)
			cur := lines[v]
			x := root.FindMaxBefore(cur.l - 1)
			if x+tmp > best {
				best = x + tmp
			}
			root = root.Insert(cur.r, x+tmp, 0, MAX_X)
		}

		best++

		dp[u] = best

		return best
	}
	res := dfs(0)

	return res - 1
}

type Node struct {
	left, right *Node
	start, end  int
	maxValue    int
}

func NewNode(start, end, val int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.maxValue = val
	return node
}

func (node *Node) Insert(p, v int, start, end int) *Node {
	if node == nil {
		node = NewNode(start, end, v)
	}
	if node.start == node.end {
		node.maxValue = max(node.maxValue, v)
		return node
	}
	mid := (node.start + node.end) / 2
	if p <= mid {
		node.left = node.left.Insert(p, v, start, mid)
	} else {
		node.right = node.right.Insert(p, v, mid+1, end)
	}
	if node.left == nil {
		node.maxValue = node.right.maxValue
	} else if node.right == nil {
		node.maxValue = node.left.maxValue
	} else {
		node.maxValue = max(node.left.maxValue, node.right.maxValue)
	}

	return node
}

func (node *Node) FindMaxBefore(p int) int {
	if node == nil {
		return 0
	}
	if node.start == node.end {
		return node.maxValue
	}
	mid := (node.start + node.end) / 2
	if p <= mid {
		return node.left.FindMaxBefore(p)
	}
	res := node.right.FindMaxBefore(p)
	if node.left != nil {
		return max(node.left.maxValue, res)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Line struct {
	l, r int
}

type Lines []Line

func (lines Lines) Len() int {
	return len(lines)
}

func (lines Lines) Less(i, j int) bool {
	return lines[i].l < lines[j].l || lines[i].l == lines[j].l && lines[i].r > lines[j].r
}

func (lines Lines) Swap(i, j int) {
	lines[i], lines[j] = lines[j], lines[i]
}
