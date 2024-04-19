package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	m := readNum(reader)

	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var tp int
		pos := readInt(s, 0, &tp)
		if tp == 1 {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 2)
		}

		queries[i][0] = tp
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos+1, &queries[i][j])
		}
	}

	res := solve(a, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	// tr.Get(i), 表示i后面未满的杯子
	tr := NewSegTree(n, n, min)

	for i := 0; i < n; i++ {
		tr.Update(i, i)
	}
	vol := make([]int, n)

	fill := func(p int, x int) {
		p = tr.Get(p, n)
		for p < n && x > 0 {
			tmp := min(a[p]-vol[p], x)
			vol[p] += tmp
			x -= tmp
			if vol[p] == a[p] {
				// cup p full
				tr.Update(p, n)
			}
			p = tr.Get(p, n)
		}
	}

	var ans []int
	for _, cur := range queries {
		if cur[0] == 1 {
			p, x := cur[1], cur[2]
			fill(p-1, x)
		} else {
			p := cur[1]
			ans = append(ans, vol[p-1])
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
	seg.arr[p] = v
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
