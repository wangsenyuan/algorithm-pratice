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
	n, k := readTwoNums(reader)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 2)
	}
	res := solve(k, points)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(k int, points [][]int) []int {
	n := len(points)

	C := make([]int64, 0, n*n)

	D := make([][]int64, n)

	for i := 0; i < n; i++ {
		D[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			D[i][j] = distance(points[i], points[j])
			if i != j {
				C = append(C, D[i][j])
			}
		}
	}

	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	var p int
	for i := 1; i <= len(C); i++ {
		if i == len(C) || C[i] > C[i-1] {
			C[p] = C[i-1]
			p++
		}
	}

	g := NewGraph(n, n*n+1)

	flag := make([]int, n)

	var remove func(u int, k int) bool

	remove = func(u int, k int) bool {
		if u == n {
			return true
		}
		if flag[u] > 0 {
			// u is already deleted
			return remove(u+1, k)
		}
		// try to keep u
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if flag[v] == 0 {
				cnt++
			}
			flag[v]++
		}
		if k >= cnt && remove(u+1, k-cnt) {
			return true
		}

		// restore
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			flag[g.to[i]]--
		}
		if cnt > 1 && k > 0 {
			// remove u
			flag[u]++
			if remove(u+1, k-1) {
				return true
			}
		}

		flag[u] = 0
		return false
	}

	check := func(x int64) bool {
		g.cur = 0
		for i := 0; i < n; i++ {
			g.nodes[i] = 0
			flag[i] = 0
		}

		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				if D[i][j] >= x {
					g.AddEdge(i, j)
					g.AddEdge(j, i)
				}
			}
		}

		return remove(0, k)
	}

	l, r := 1, p
	var pick int64
	for l <= r {
		mid := (l + r) >> 1
		if check(C[mid-1]) {
			pick = C[mid-1]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	check(pick)

	ans := make([]int, 0, k)

	for i := 0; i < n && len(ans) < k; i++ {
		if flag[i] > 0 {
			ans = append(ans, i+1)
		}
	}

	for i := 0; i < n && len(ans) < k; i++ {
		if flag[i] == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func distance(a, b []int) int64 {
	dx := int64(a[0] - b[0])
	dy := int64(a[1] - b[1])
	return dx*dx + dy*dy
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
