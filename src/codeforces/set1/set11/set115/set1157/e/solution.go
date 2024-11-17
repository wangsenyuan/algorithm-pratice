package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	res := solve(a, b)
	s := fmt.Sprintf("%v", res)
	return s[1 : len(s)-1]
}

func solve(a []int, b []int) []int {
	n := len(a)
	cnt := make([]int, n)
	for _, x := range b {
		cnt[x]++
	}

	set := NewSegTree(n)

	for i := 0; i < n; i++ {
		if cnt[i] > 0 {
			set.Update(i, i)
			cnt[i]--
		}
	}

	c := make([]int, n)

	for i := 0; i < n; i++ {
		v := a[i]
		// (v + b[i]) % n= 0, 1, 2,
		// n - v, n
		x := set.Get(n-v, n)
		if x == inf {
			x = set.Get(0, n-v)
		}
		c[i] = (v + x) % n
		cnt[x]--
		if cnt[x] < 0 {
			set.Update(x, inf)
		}
	}

	return c
}

const inf = 1 << 30

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
