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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(k, a)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
	}

	fmt.Print(buf.String())
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

func solve(k int, a []int) []int {
	n := len(a)

	pos := make([][]int, k+1)

	for i := 0; i < n; i++ {
		x := a[i]
		if pos[x] == nil {
			pos[x] = make([]int, 0, 1)
		}
		pos[x] = append(pos[x], i+1)
	}

	res := make([]int, k+1)
	l, r := n, 1
	for i := k; i >= 1; i-- {
		if len(pos[i]) == 0 {
			continue
		}

		for _, j := range pos[i] {
			l = min(l, j)
			r = max(r, j)
		}
		res[i] = (r - l + 1) * 2
	}

	return res[1:]

}

func solve1(k int, a []int) []int {
	n := len(a)

	l := make([]int, k+1)
	r := make([]int, k+1)

	for i := 0; i <= k; i++ {
		l[i] = n
		r[i] = 0
	}

	tr := NewSegTree(k, n, min)

	for i := 0; i < n; i++ {
		x := a[i] - 1
		tr.Update(x, i)
		l[x] = min(l[x], tr.Get(x, k)-1)
	}

	tr = NewSegTree(k, 0, max)

	for i := n - 1; i >= 0; i-- {
		x := a[i] - 1
		tr.Update(x, i)
		r[x] = max(r[x], tr.Get(x, k)+1)
	}

	res := make([]int, k)

	for i := 0; i < k; i++ {
		if r[i] == 0 {
			continue
		}
		res[i] = (r[i] - l[i] - 1) * 2
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
	return a + b - max(a, b)
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
