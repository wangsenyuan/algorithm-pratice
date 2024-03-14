package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := 1

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		b := readNNums(reader, n)
		res := solve(b, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(b []int, k int) int {
	n := len(b)

	sum := NewSegTree(n)
	cnt := NewSegTree(n)
	var res int
	for i := n - 1; i >= 0; i-- {
		l := max(i-k+1, 0)
		x := cnt.query(l, i+1)
		s := sum.query(l, i+1)
		tmp := x*i + x - s
		if tmp >= b[i] {
			continue
		}
		// tmp < b[i]
		d := i - l + 1
		c := (b[i] - tmp + d - 1) / d
		res += c
		cnt.update(l, c)
		sum.update(l, c*l)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) update(p int, v int) {
	for i := p + t.n; i > 0; i >>= 1 {
		t.arr[i] += v
	}
}

func (t *SegTree) query(l int, r int) int {
	l += t.n
	r += t.n
	var res int
	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
