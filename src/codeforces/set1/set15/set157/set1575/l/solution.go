package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {

	n := len(a)

	b := make([]pair, n)

	for i := 0; i < n; i++ {
		b[i] = pair{i + 1 - a[i], i}
	}

	slices.SortFunc(b, func(x, y pair) int {
		if x.first != y.first {
			return x.first - y.first
		}
		return x.second - y.second
	})

	set := NewSegTree(n)

	for _, cur := range b {
		if cur.first < 0 {
			continue
		}
		// i - a[i] >= j - a[j] 这个是成立的
		i := min(cur.second, a[cur.second]-1)
		v := set.Get(0, i)
		v++
		set.Update(i, v)
	}

	return set.Get(0, n)
}

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)

	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = max(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	var res int
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = max(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
