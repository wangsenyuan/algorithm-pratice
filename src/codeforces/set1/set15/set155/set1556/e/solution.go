package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, b, queries)
}

func solve(a []int, b []int, queries [][]int) []int {
	n := len(a)
	pa := make([]int, n+1)
	pb := make([]int, n+1)
	next := make([]int, n)

	for i := range n {
		next[i] = n
		pa[i+1] = pa[i] + a[i]
		pb[i+1] = pb[i] + b[i]
	}

	tr := NewSegTree(n)

	at := make([][]int, n)
	for i, cur := range queries {
		l := cur[0] - 1
		at[l] = append(at[l], i)
	}

	ans := make([]int, len(queries))

	for i := n - 1; i >= 0; i-- {
		diff := b[i] - a[i]
		tr.Update(i, n-1, diff)
		if tr.val[0] < 0 {
			// some place sum_a[j] > sum_b[j]
			next[i] = tr.GetMinPosition()
		}

		for _, j := range at[i] {
			// l == i
			r := queries[j][1] - 1
			if pb[r+1]-pb[i] != pa[r+1]-pa[i] || next[i] < r {
				ans[j] = -1
				continue
			}
			ans[j] = tr.GetMaxVal(i, r)
		}
	}

	return ans
}

type SegTree struct {
	val  []int
	mx   []int
	lazy []int
	sz   int
}

func NewSegTree(n int) *SegTree {
	val := make([]int, 4*n)
	mx := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &SegTree{val, mx, lazy, n}
}

func (tr *SegTree) update(i int, v int) {
	tr.val[i] += v
	tr.mx[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.val[i] = min(tr.val[2*i+1], tr.val[2*i+2])
		tr.mx[i] = max(tr.mx[2*i+1], tr.mx[2*i+2])
	}

	loop(0, 0, tr.sz-1)
}

type pair struct {
	first  int
	second int
}

func (tr *SegTree) GetMinPosition() int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		tr.push(i)
		mid := (l + r) / 2
		if tr.val[2*i+1] < 0 {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}

	return loop(0, 0, tr.sz-1)
}

func (tr *SegTree) GetMaxVal(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return tr.mx[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		x := loop(2*i+1, l, mid)
		y := loop(2*i+2, mid+1, r)
		return max(x, y)
	}
	return loop(0, 0, tr.sz-1)
}
