package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		cnt, res, _ := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n%d %d\n", cnt, res[0], res[1]))
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (cnt int, res []int, points [][]int) {
	n := readNum(reader)
	points = make([][]int, n)
	for i := range n {
		points[i] = readNNums(reader, 2)
	}
	cnt, res = solve(points)
	return
}

const inf = 1e9

func solve(points [][]int) (int, []int) {
	n := len(points)
	// 先压缩
	xs := make(map[int]int)
	ys := make(map[int]int)
	xs[inf]++
	ys[inf]++
	for _, p := range points {
		xs[p[0]]++
		ys[p[1]]++
	}
	a := compact(xs)
	b := compact(ys)

	w := len(xs)
	h := len(ys)

	bars := make([][]int, w)

	for _, p := range points {
		x := xs[p[0]]
		y := ys[p[1]]
		bars[x] = append(bars[x], y)
	}

	for i := range w {
		sort.Ints(bars[i])
	}

	L := NewSegTree(h)
	R := NewSegTree(h)

	get := func(tr *SegTree, cnt int) (int, int) {
		if tr.val[0]*2 < cnt {
			return -1, -1
		}
		// tr.val[0] >= cnt
		l := tr.LowerBound(cnt)
		r := tr.UpperBound(cnt)
		if l+1 <= r {
			return l + 1, r
		}
		return -1, -1
	}

	check := func(exp int) []int {
		L.Reset()
		R.Reset()

		// 全部先加到右上
		for i := range w {
			for _, j := range bars[i] {
				R.Update(j, 1)
			}
		}

		for i := range w {
			l1, r1 := get(L, exp)
			l2, r2 := get(R, exp)
			if l1 >= 0 && l2 >= 0 && l1 <= r2 && l2 <= r1 {
				return []int{i, min(r1, r2)}
			}
			for _, j := range bars[i] {
				R.Update(j, -1)
				L.Update(j, 1)
			}
		}
		return nil
	}

	l, r := 0, n/4+1

	for l < r {
		mid := (l + r) / 2
		ans := check(mid)
		if len(ans) > 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}

	res := check(r - 1)

	return r - 1, []int{a[res[0]], b[res[1]]}
}

func compact(xs map[int]int) []int {
	var x []int
	for k := range xs {
		x = append(x, k)
	}
	sort.Ints(x)
	for i, k := range x {
		xs[k] = i
	}
	return x
}

type SegTree struct {
	val []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	return &SegTree{arr, n}
}

func (tr *SegTree) Reset() {
	clear(tr.val)
}

func (tr *SegTree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = tr.val[2*i+1] + tr.val[2*i+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) LowerBound(v int) int {
	var loop func(i int, l int, r int, v int) int

	loop = func(i int, l int, r int, v int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.val[2*i+1] >= v {
			return loop(2*i+1, l, mid, v)
		}
		return loop(2*i+2, mid+1, r, v-tr.val[2*i+1])
	}

	return loop(0, 0, tr.sz-1, v)
}

func (tr *SegTree) UpperBound(v int) int {
	var loop func(i int, l int, r int, v int) int

	loop = func(i int, l int, r int, v int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.val[2*i+2] >= v {
			return loop(2*i+2, mid+1, r, v)
		}
		return loop(2*i+1, l, mid, v-tr.val[2*i+2])
	}

	return loop(0, 0, tr.sz-1, v)
}
