package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	ops := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k)
		ops[i] = make([]int, 2*k)
		for j := 0; j < 2*k; j += 2 {
			pos = readInt(s, pos+1, &ops[i][j])
			pos = readInt(s, pos+1, &ops[i][j+1])
		}
	}
	res := solve(ops)
	fmt.Println(res)
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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	if b < 0 {
		b += MOD
	}
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(ops [][]int) int {
	n := len(ops)
	sum := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	var root func(u int) int
	root = func(u int) int {
		if arr[u] != u {
			p := root(arr[u])
			sum[u] = modAdd(sum[u], sum[arr[u]])
			arr[u] = p
		}
		return arr[u]
	}
	var res int
	for u := 0; u < n; u++ {
		op := ops[u]
		for i := 0; i < len(op); i += 2 {
			v := op[i] - 1
			r := root(v)
			x := modAdd(op[i+1], sum[v])
			sum[r] = x
			arr[r] = u
			res = modAdd(res, x)
		}
	}
	return res
}

func solve1(ops [][]int) int {
	// root(i)
	//root is easy to use some kind fo union-find
	// but depth is kind of hard
	// because the ops later will change the depth
	n := len(ops)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	deg := make([]int, n)
	var find func(x int) int

	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}
	g := NewGraph(n, n)
	// 分成两步来进行， 先构建树，然后计算euler order，然后 range update
	for u := 0; u < n; u++ {
		op := ops[u]
		for i := 0; i < len(op); i += 2 {
			v := op[i] - 1
			r := find(v)
			deg[r]++
			g.AddEdge(u, r)
			arr[r] = u
		}
	}
	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = -1
	}
	out := make([]int, n)
	var dfs func(u int)
	var time int
	dfs = func(u int) {
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
		}
		out[u] = time
	}

	for u := 0; u < n; u++ {
		if deg[u] == 0 {
			dfs(u)
		}
	}

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	sum := make([]int, 2*n)

	set := func(l, r int, v int) {
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				sum[l] = modAdd(sum[l], v)
				l++
			}
			if r&1 == 1 {
				r--
				sum[r] = modAdd(sum[r], v)
			}
			l >>= 1
			r >>= 1
		}
	}

	get := func(p int) int {
		p += n
		var res int
		for p > 0 {
			res = modAdd(res, sum[p])
			p >>= 1
		}
		return res
	}
	var res int

	for u := 0; u < n; u++ {
		op := ops[u]
		for i := 0; i < len(op); i += 2 {
			v := op[i] - 1
			x := op[i+1]
			d := get(in[v])
			x = modAdd(x, d)
			r := find(v)
			set(in[r], out[r], x)
			arr[r] = u
			res = modAdd(res, x)
		}
	}

	return res
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
