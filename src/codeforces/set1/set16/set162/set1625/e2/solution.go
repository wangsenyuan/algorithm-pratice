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
	s := readString(reader)[:n]

	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}

	res := solve(s, Q)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(s string, Q [][]int) []int64 {
	s = "(" + s + ")"
	n := len(s)
	stack := make([]int, n)
	var top int

	pair := make([]int, n)

	pair[0] = n - 1
	pair[n-1] = 0

	for i := 1; i+1 < n; i++ {
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			if top > 0 {
				j := stack[top-1]
				pair[i] = j
				pair[j] = i
				top--
			}
		}
	}

	g := NewGraph(n, 2*n)
	in := make([]int, n)
	out := make([]int, n)
	var build func(u int, l int, r int)
	var time int

	tot := 1

	L := make([]int, n)
	R := make([]int, n)
	id := make([]int, n)
	sz := make([]int64, n)
	FA := make([]int, n)
	dp := make([]int64, n)
	build = func(u int, l int, r int) {
		in[u] = time
		time++
		l++
		for l < r {
			for pair[l] == 0 && l < r {
				// invalid (
				l++
			}
			if l >= r {
				break
			}
			tot++
			g.AddEdge(u, tot)
			L[tot] = l
			R[tot] = pair[l]
			id[l] = tot
			id[pair[l]] = tot
			l = pair[l] + 1
			sz[u]++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			FA[v] = u
			build(v, L[v], R[v])
		}
		dp[u] = 1 + sz[u]*(sz[u]-1)/2
		out[u] = time
	}

	build(1, 0, n-1)

	A := NewBit[int64](tot + 1)
	B := NewBit[int](tot + 1)

	for i := 1; i <= tot; i++ {
		A.Add(in[i], dp[i])
		B.Add(i, 1)
	}

	var ans []int64

	for _, cur := range Q {
		t, l, r := cur[0], cur[1], cur[2]
		if t == 2 {
			u := id[l]
			v := id[r]
			if u > v {
				u, v = v, u
			}
			ln := B.GetRange(u, v)
			tmp := int64(ln) * int64(ln-1) / 2

			if in[u] < out[v] {
				tmp += A.GetRange(in[u], out[v]-1)
			} else {
				tmp += A.GetRange(in[v], out[u]-1)
			}

			ans = append(ans, tmp)
		} else {
			u := id[l]
			A.Add(in[u], -1)
			B.Add(u, -1)
			sz[FA[u]]--
			A.Add(in[FA[u]], -sz[FA[u]])
		}
	}

	return ans
}

type Num interface {
	int | int64
}

type Bit[T Num] struct {
	arr []T
}

func NewBit[T Num](sz int) *Bit[T] {
	arr := make([]T, sz+1)
	return &Bit[T]{arr}
}

func (b *Bit[T]) Add(p int, v T) {
	p++
	for p < len(b.arr) {
		b.arr[p] += v
		p += p & -p
	}
}

func (b *Bit[T]) Get(p int) T {
	p++
	var res T
	for p > 0 {
		res += b.arr[p]
		p -= p & -p
	}
	return res
}

func (b *Bit[T]) GetRange(l int, r int) T {
	return b.Get(r) - b.Get(l-1)
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
