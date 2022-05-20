package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	A := readNNums(reader, n)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, A, Q)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
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

func solve(n int, A []int, Q [][]int) []int {

	right := NewSegTree(n+1, n, min)

	R := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			R[i] = right.Get(A[i], A[i+1]+1)
		}
		right.Update(A[i+1], i+1)
	}

	left := NewSegTree(n+1, -1, max)

	L := make([]int, n)

	for i := 0; i+1 < n; i++ {
		if A[i] < A[i+1] {
			L[i] = left.Get(A[i], A[i+1]+1)
		} else {
			L[i] = n - 1
		}
		left.Update(A[i], i)
	}

	evt := make([][]Pair, n+1)
	que := make([][]Pair, n+1)

	for i := 0; i <= n; i++ {
		evt[i] = make([]Pair, 0, 1)
		que[i] = make([]Pair, 0, 1)
	}

	for i := 0; i < n-1; i++ {
		if L[i] < i && i+1 < R[i] {
			evt[L[i]+1] = append(evt[L[i]+1], Pair{i + 1, 1})
			evt[L[i]+1] = append(evt[L[i]+1], Pair{R[i], -1})
			evt[i+1] = append(evt[i+1], Pair{i + 1, -1})
			evt[i+1] = append(evt[i+1], Pair{R[i], 1})
		}
	}

	ans := make([]int, len(Q))

	for i := 0; i < len(Q); i++ {
		l, r := Q[i][0], Q[i][1]
		l--
		r--
		ans[i] = r - l + 1
		que[l] = append(que[l], Pair{r, i})
	}

	rect := NewSegTree(n+1, 0, plus)

	for i := 0; i <= n; i++ {
		for _, e := range evt[i] {
			r, v := e.first, e.second
			rect.Update(r, v)
		}
		for _, q := range que[i] {
			r, j := q.first, q.second
			ans[j] -= rect.Get(0, r+1)
		}
	}

	return ans
}

type Pair struct {
	first, second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func plus(a, b int) int {
	return a + b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = seg.op(v, seg.arr[p])
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
