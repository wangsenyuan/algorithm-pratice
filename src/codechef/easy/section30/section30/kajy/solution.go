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

	A := readNNums(reader, n)
	B := readNNums(reader, n)
	C := readNNums(reader, n)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(A, B, C, Q)
	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(A []int, B []int, C []int, Q [][]int) []int64 {
	n := len(A)
	P := make([]int64, n+2)

	for i := 1; i <= n; i++ {
		P[i] = P[i-1] + int64(A[i-1])
	}

	M := make([]int64, n+2)
	for i := 1; i <= n; i++ {
		M[i] = M[i-1] + int64(B[i-1])
	}
	S := make([]int64, n+3)

	for i := n; i > 0; i-- {
		S[i] = S[i+1] + int64(C[i-1])
	}

	st := NewSegTree(P, M, S, n)

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		t, p, v := cur[0], cur[1], cur[2]
		p--
		if t == 1 {
			st.Update(0, p+1, n, int64(v-A[p]))
			A[p] = v
		} else if t == 2 {
			st.Update(0, p+1, n, int64(B[p]-v))
			st.Update(1, p+1, n, int64(v-B[p]))
			B[p] = v
		} else {
			st.Update(1, 1, p, int64(v-C[p]))
			C[p] = v
		}
		ans[i] = st.nodes[1].ans
	}

	return ans
}

type Node struct {
	ans    int64
	max_d  int64
	max_e  int64
	lazy_d int64
	lazy_e int64
}

const N = 100010

func (this *Node) combine(left, right *Node) {
	this.ans = max(left.ans, right.ans)
	this.ans = max(this.ans, left.max_d+right.max_e)
	this.max_d = max(left.max_d, right.max_d)
	this.max_e = max(left.max_e, right.max_e)
}

type SegTree struct {
	nodes []*Node
	sz    int
}

func NewSegTree(P, M, S []int64, n int) *SegTree {
	sz := n
	nodes := make([]*Node, 4*(sz+2))
	var build func(id int, s int, e int)

	build = func(id int, s int, e int) {
		nodes[id] = new(Node)

		if s == e {
			nodes[id].ans = 0
			nodes[id].max_d = P[s] - M[s]
			nodes[id].max_e = S[s+1] + M[s]
			return
		}
		mid := (s + e) >> 1
		build(id*2, s, mid)
		build(id*2+1, mid+1, e)
		nodes[id].combine(nodes[id*2], nodes[id*2+1])
	}

	build(1, 0, n)

	return &SegTree{nodes, sz}
}

func (st *SegTree) Update(t int, l int, r int, v int64) {
	if l > r {
		return
	}
	var loop func(id int, s int, e int)

	loop = func(id int, s int, e int) {
		if e < l || r < s {
			return
		}
		if l <= s && e <= r {
			if t == 0 {
				st.nodes[id].max_d += v
				st.nodes[id].lazy_d += v
			} else {
				st.nodes[id].max_e += v
				st.nodes[id].lazy_e += v
			}
			st.nodes[id].ans += v
			return
		}
		st.push(id)
		mid := (s + e) >> 1
		loop(id*2, s, mid)
		loop(id*2+1, mid+1, e)

		st.nodes[id].combine(st.nodes[id*2], st.nodes[id*2+1])
	}

	loop(1, 0, st.sz)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func (st *SegTree) push(id int) {
	root := st.nodes[id]
	left := st.nodes[2*id]
	right := st.nodes[2*id+1]
	x := root.lazy_d
	y := root.lazy_e

	left.ans += x + y
	right.ans += x + y

	left.max_d += x
	left.lazy_d += x
	right.max_d += x
	right.lazy_d += x

	left.max_e += y
	left.lazy_e += y
	right.max_e += y
	right.lazy_e += y

	root.lazy_d = 0
	root.lazy_e = 0
}
