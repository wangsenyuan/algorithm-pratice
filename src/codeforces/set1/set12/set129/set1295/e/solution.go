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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	p := readNNums(reader, n)
	a := readNNums(reader, n)
	return solve(p, a)
}

const inf = 1 << 60

func solve(p []int, a []int) int {
	n := len(p)
	pos := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		p[i]--
		pos[p[i]] = i
		b[i] = a[i]
		if i > 0 {
			b[i] += b[i-1]
		}
	}
	tr := NewTree(b)
	res := tr.Get(0, n-2)

	for i := 0; i < n; i++ {
		j := pos[i]
		tr.Update(j, n-1, -a[j])
		tr.Update(0, j-1, a[j])
		res = min(res, tr.Get(0, n-2))
	}

	return res
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = min(val[i*2+1], val[i*2+2])
	}
	build(0, 0, n-1)
	return &Tree{val, lazy, n}
}

func (tr *Tree) update(i int, v int) {
	tr.val[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = min(tr.val[2*i+1], tr.val[2*i+2])
}

// add v to segment L...R
func (tr *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
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
		tr.pull(i)
	}
	loop(0, 0, tr.sz-1)
}

// get min value in the rage L...R
func (tr *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return inf
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		a := loop(2*i+1, l, mid)
		b := loop(2*i+2, mid+1, r)
		return min(a, b)
	}
	return loop(0, 0, tr.sz-1)
}

func solve1(p []int, a []int) int {
	n := len(p)
	tr1 := NewSegTree(n)
	for i := 0; i < n; i++ {
		p[i]--
		tr1.Update(p[i], a[i])
	}

	tr2 := NewSegTree(n)

	count := func(x int) int {
		return tr2.Get(x+1, n) + tr1.Get(0, x)
	}

	search := func() int {
		l, r := 0, n
		for r-l >= 3 {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			a := count(m1)
			b := count(m2)
			if a < b {
				r = m2
			} else {
				l = m1
			}
		}
		res := inf
		for i := l; i < r; i++ {
			res = min(res, count(i))
		}
		return res
	}

	best := inf

	for i := 0; i+1 < n; i++ {
		tr2.Update(p[i], a[i])
		tr1.Update(p[i], -a[i])
		// 这个地方不对
		// 如果前面有x个，那么就是把所有 大于 x的移动到后面，且把后面<= x 的部分移动到前面
		// 这个x不一定是 i+1, 可以是0，1，2， 。。。n
		best = min(best, search())
	}

	return best
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (tr *SegTree) Update(p int, v int) {
	p += tr.sz
	tr.arr[p] += v
	for p > 0 {
		tr.arr[p>>1] = tr.arr[p] + tr.arr[p^1]
		p >>= 1
	}
}

func (tr *SegTree) Get(l int, r int) int {
	var res int
	l += tr.sz
	r += tr.sz
	for l < r {
		if l&1 == 1 {
			res += tr.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tr.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
