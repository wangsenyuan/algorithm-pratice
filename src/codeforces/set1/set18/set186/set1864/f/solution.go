package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	qs := make([][]int, m)
	for i := 0; i < m; i++ {
		qs[i] = readNNums(reader, 2)
	}
	return solve(a, qs)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, qs [][]int) []int {
	n := len(a)
	pos := make([][]int, n+1)
	for i, x := range a {
		pos[x] = append(pos[x], i)
	}

	var pts []pair
	t1 := NewSegTree(n, func(a, b int) int {
		return max(a, b)
	})
	distict_cnt := make([]int, n+1)
	for x := 1; x <= n; x++ {
		distict_cnt[x] = distict_cnt[x-1]
		if len(pos[x]) > 0 {
			distict_cnt[x]++
		}
		for i := 0; i+1 < len(pos[x]); i++ {
			lf := pos[x][i]
			rg := pos[x][i+1]
			if lf+1 < rg {
				y := t1.Get(lf+1, rg, 0)
				if y == 0 {
					continue
				}
				// 这里y肯定比x小
				pts = append(pts, pair{n - y, x})
			}
		}
		for _, i := range pos[x] {
			t1.Update(i, x)
		}
	}
	slices.SortFunc(pts, func(x, y pair) int {
		if x.first != y.first {
			return x.first - y.first
		}
		return x.second - y.second
	})
	ans := make([]int, len(qs))

	type item struct {
		id int
		l  int
		r  int
	}
	items := make([]item, len(qs))
	for i, cur := range qs {
		items[i] = item{i, n - cur[0], cur[1]}
		ans[i] = distict_cnt[cur[1]] - distict_cnt[cur[0]-1]
	}

	slices.SortFunc(items, func(x, y item) int {
		if x.l != y.l {
			return x.l - y.l
		}
		return x.r - y.r
	})

	t2 := NewSegTree(n+1, func(a, b int) int {
		return a + b
	})

	var ptr int
	for _, it := range items {
		for ptr < len(pts) && pts[ptr].first <= it.l {
			t2.Update(pts[ptr].second, 1)
			ptr++
		}
		ans[it.id] += t2.Get(0, it.r+1, 0)
	}

	return ans
}

type SegTree struct {
	arr []int
	sz  int
	f   func(int, int) int
}

func NewSegTree(n int, f func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n, f}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = t.f(t.arr[p], v)
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int, iv int) int {
	res := iv
	l += t.sz
	r += t.sz

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
