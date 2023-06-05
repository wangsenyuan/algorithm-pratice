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
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(A, Q)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const INF = 1 << 30

func fillArray(n int, iv int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = iv
	}
	return res
}

func solve(A []int, Q [][]int) []int {
	n := len(A)

	x := NewTree(fillArray(n-1, 0), 0, max)
	y := NewTree(fillArray(n-1, INF), INF, min)

	update := func(i int) {
		a, b := 0, INF
		if A[i] != A[i+1] {
			if A[i] < A[i+1] {
				b = (A[i] + A[i+1]) / 2
			} else {
				a = (A[i] + A[i+1] + 1) / 2
			}
		}
		x.Update(i, a)
		y.Update(i, b)
	}

	for i := 0; i+1 < n; i++ {
		update(i)
	}

	ans := make([]int, 0, len(Q))

	for _, cur := range Q {
		if cur[0] == 1 {
			i, v := cur[1], cur[2]
			i--
			A[i] = v
			if i > 0 {
				update(i - 1)
			}
			if i+1 < n {
				update(i)
			}

		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			if l == r {
				ans = append(ans, 0)
				continue
			}
			a := x.Get(l, r)
			b := y.Get(l, r)
			if a <= b {
				ans = append(ans, a)
			} else {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}

type Tree struct {
	arr []int
	sz  int
	f   func(a, b int) int
	iv  int
}

func NewTree(arr []int, iv int, f func(int, int) int) *Tree {
	n := len(arr)
	rra := make([]int, 2*n)
	copy(rra[n:], arr)
	for i := n - 1; i > 0; i-- {
		rra[i] = f(rra[i*2], rra[i*2+1])
	}
	return &Tree{rra, n, f, iv}
}

func (t *Tree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *Tree) Get(l int, r int) int {
	l += t.sz
	r += t.sz
	res := t.iv
	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
