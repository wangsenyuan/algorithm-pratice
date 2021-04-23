package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)

	var buf bytes.Buffer

	A := readNNums(reader, n)
	B := readNNums(reader, n)

	solver := NewSolver(A, B)

	for q > 0 {
		q--
		var pos int
		var t int
		s, _ := reader.ReadBytes('\n')
		pos = readInt(s, 0, &t) + 1
		if t == 1 {
			var z, y int
			pos = readInt(s, pos, &z) + 1
			readInt(s, pos, &y)
			solver.UpdatePers(z, y)
			continue
		}
		if t == 2 {
			var l, r, x int
			pos = readInt(s, pos, &l) + 1
			pos = readInt(s, pos, &r) + 1
			readInt(s, pos, &x)
			solver.Update(l, r, x)
			continue
		}
		var l, r int
		pos = readInt(s, pos, &l) + 1
		readInt(s, pos, &r)
		ans := solver.Get(l, r)
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

type Solver struct {
	a *Treap
	b *Treap
}

func NewSolver(A, B []int) Solver {
	// build two treaps from A & B
	rand.Seed(42)

	a := NewTreap(A)
	b := NewTreap(B)

	return Solver{&a, &b}
}

func (solver *Solver) UpdatePers(pos int, val int) {
	pos--
	solver.b.Update(pos, val)
}

func (solver *Solver) Get(l, r int) int64 {
	l--
	r--
	ans := solver.a.GetSumBefore(r)
	if l > 0 {
		ans -= solver.a.GetSumBefore(l - 1)
	}
	return ans
}

func (solver *Solver) Update(l int, r int, beg int) {
	l--
	r--
	beg--
	end := beg + r - l
	var lt, rt *Item
	split2(solver.b.root, beg-1, &lt, &rt)
	var at, bt *Item
	split2(rt, end, &at, &bt)
	// at is root
	var i, j *Item
	split(solver.a.root, l-1, &i, &j)
	var u, v *Item
	split(j, r, &u, &v)

	var ii *Item
	merge(&ii, i, at)
	merge(&ii, ii, v)
	solver.a.root = ii
}

type Treap struct {
	root *Item
}

func (t *Treap) Update(pos int, val int) {
	update(t.root, pos, int64(val))
}

func (t *Treap) GetSumBefore(key int) int64 {
	return getSumBefore(t.root, key)
}

func NewTreap(arr []int) Treap {
	var root *Item
	for i := 0; i < len(arr); i++ {
		insert(&root, NewItem(i, int64(arr[i])))
	}
	return Treap{root}
}

type Item struct {
	key   int
	value int64
	sum   int64
	prior int64
	left  *Item
	right *Item
}

func NewItem(key int, value int64) *Item {
	item := new(Item)
	item.key = key
	item.value = value
	item.sum = value
	item.prior = rand.Int63()
	return item
}

func (t *Item) clone() *Item {
	item := NewItem(t.key, t.value)
	item.prior = t.prior
	item.sum = t.sum
	item.left = t.left
	item.right = t.right
	return item
}

func (t *Item) GetSum() int64 {
	if t == nil {
		return 0
	}
	return t.sum
}

func (t *Item) UpdateSum() {
	if t == nil {
		return
	}
	t.sum = t.value + t.left.GetSum() + t.right.GetSum()
}

func update(t *Item, key int, value int64) {
	if t == nil {
		return
	}
	if t.key == key {
		t.value = value
	} else if key < t.key {
		update(t.left, key, value)
	} else {
		update(t.right, key, value)
	}
	t.UpdateSum()
}

func getSumBefore(t *Item, key int) int64 {
	if t == nil {
		return 0
	}
	if key < t.key {
		return getSumBefore(t.left, key)
	}
	// k >= t.key
	res := t.left.GetSum() + t.value
	res += getSumBefore(t.right, key)
	return res
}

func split(t *Item, key int, left **Item, right **Item) {
	if t == nil {
		*left = nil
		*right = nil
	} else if key < t.key {
		//tt := t.clone()
		split(t.left, key, left, &t.left)
		*right = t
	} else {
		split(t.right, key, &t.right, right)
		*left = t
	}
	t.UpdateSum()
}

func split2(t *Item, key int, left **Item, right **Item) {
	if t == nil {
		*left = nil
		*right = nil
		return
	}
	t = t.clone()
	if key < t.key {
		split2(t.left, key, left, &t.left)
		*right = t
	} else {
		split2(t.right, key, &t.right, right)
		*left = t
	}
	t.UpdateSum()
}

func insert(t **Item, it *Item) {
	if *t == nil {
		*t = it
	} else if it.prior > (*t).prior {
		split(*t, it.key, &it.left, &it.right)
		*t = it
	} else if it.key < (*t).key {
		insert(&(*t).left, it)
	} else {
		insert(&(*t).right, it)
	}
	(*t).UpdateSum()
}

func merge(t **Item, l, r *Item) {
	if l == nil {
		*t = r
	} else if r == nil {
		*t = l
	} else if l.prior > r.prior {
		merge(&l.right, l.right, r)
		*t = l
	} else {
		merge(&r.left, l, r.left)
		*t = r
	}
	(*t).UpdateSum()
}

func union(l, r *Item) *Item {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.prior < r.prior {
		l, r = r, l
	}
	var lt, rt *Item
	split(r, l.key, &lt, &rt)
	l.left = union(l.left, lt)
	l.right = union(l.right, rt)
	l.UpdateSum()
	return l
}

func heapify(t *Item) {
	if t == nil {
		return
	}
	max := t
	if t.left != nil && t.left.prior > max.prior {
		max = t.left
	}
	if t.right != nil && t.right.prior > max.prior {
		max = t.right
	}
	if max != t {
		t.prior, max.prior = max.prior, t.prior
		heapify(max)
	}
}

type Solver1 struct {
	cur      int
	n        int
	root     []int
	tl       []int
	tr       []int
	sz       int
	smPers   []int64
	sm       []int64
	lazy     []int
	lazyRoot []int
}

func NewSolver1(A []int, B []int) Solver1 {
	n := len(A)
	root := make([]int, 4*n*20)
	tl := make([]int, 4*n*20)
	tr := make([]int, 4*n*20)
	smPers := make([]int64, 80*n)
	sm := make([]int64, 4*n)
	lazy := make([]int, 4*n)
	lazyRoot := make([]int, 4*n)

	var sz int
	var loop1 func(b, e int) int
	loop1 = func(b, e int) int {
		id := sz
		sz++
		if e-b == 1 {
			smPers[id] = int64(B[b])
			return id
		}
		mid := (b + e) >> 1
		tl[id] = loop1(b, mid)
		tr[id] = loop1(mid, e)
		smPers[id] = smPers[tl[id]] + smPers[tr[id]]
		return id
	}

	root[0] = loop1(0, n)

	var loop2 func(v int, b, e int)

	loop2 = func(v int, b, e int) {
		lazy[v] = -1
		if e-b == 1 {
			sm[v] = int64(A[b])
			return
		}
		mid := (b + e) >> 1
		loop2(v<<1, b, mid)
		loop2((v<<1)^1, mid, e)
		sm[v] = sm[v<<1] + sm[(v<<1)^1]
	}

	loop2(1, 0, n)
	return Solver1{0, n, root, tl, tr, sz, smPers, sm, lazy, lazyRoot}
}

func (solver *Solver1) UpdatePers(pos int, val int) {
	pos--
	var loop func(id int, b int, e int) int
	loop = func(id int, b int, e int) int {
		newId := solver.sz
		solver.sz++
		if e-b == 1 {
			solver.smPers[newId] = int64(val)
			return newId
		}
		mid := (b + e) >> 1
		solver.tl[newId] = solver.tl[id]
		solver.tr[newId] = solver.tr[id]
		if pos < mid {
			solver.tl[newId] = loop(solver.tl[id], b, mid)
		} else {
			solver.tr[newId] = loop(solver.tr[id], mid, e)
		}
		solver.smPers[newId] = solver.smPers[solver.tl[newId]] + solver.smPers[solver.tr[newId]]
		return newId
	}
	solver.root[solver.cur+1] = loop(solver.root[solver.cur], 0, solver.n)
	solver.cur++
}

func (solver *Solver1) GetPers(id, l, r int) int64 {
	var loop func(v int, b int, e int) int64
	loop = func(id int, b int, e int) int64 {
		if l <= b && e <= r {
			return solver.smPers[id]
		}
		if r <= b || e <= l {
			return 0
		}
		mid := (b + e) >> 1
		x := loop(solver.tl[id], b, mid)
		y := loop(solver.tr[id], mid, e)
		return x + y
	}

	return loop(id, 0, solver.n)
}

func (solver *Solver1) pushDown(v, b, e, mid int) {
	if solver.lazy[v] < 0 {
		return
	}
	solver.lazyRoot[v<<1] = solver.lazyRoot[v]
	solver.lazyRoot[(v<<1)^1] = solver.lazyRoot[v]

	solver.lazy[v<<1] = solver.lazy[v]
	solver.sm[v<<1] = solver.GetPers(solver.lazyRoot[v], solver.lazy[v], solver.lazy[v]+(mid-b))

	solver.lazy[(v<<1)^1] = solver.lazy[v] + mid - b
	solver.sm[(v<<1)^1] = solver.GetPers(solver.lazyRoot[v], solver.lazy[(v<<1)^1], solver.lazy[(v<<1)^1]+e-mid)

	solver.lazy[v] = -1
}

func (solver *Solver1) Update(l int, r int, beg int) {
	l--
	beg--
	var loop func(v int, b int, e int, beg int)
	loop = func(v int, b int, e int, beg int) {
		if l <= b && e <= r {
			beg += b - l
			solver.sm[v] = solver.GetPers(solver.root[solver.cur], beg, beg+(e-b))
			solver.lazy[v] = beg
			solver.lazyRoot[v] = solver.root[solver.cur]
			return
		}
		if r <= b || e <= l {
			return
		}
		mid := (b + e) >> 1
		solver.pushDown(v, b, e, mid)
		loop(v<<1, b, mid, beg)
		loop((v<<1)^1, mid, e, beg)
		solver.sm[v] = solver.sm[v<<1] + solver.sm[(v<<1)^1]
	}
	loop(1, 0, solver.n, beg)
}

func (solver *Solver1) Get(l, r int) int64 {
	l--
	var loop func(v int, b, e int) int64
	loop = func(v int, b, e int) int64 {
		if l <= b && e <= r {
			return solver.sm[v]
		}
		if r <= b || e <= l {
			return 0
		}
		mid := (b + e) >> 1
		solver.pushDown(v, b, e, mid)
		x := loop(v<<1, b, mid)
		y := loop((v<<1)^1, mid, e)
		return x + y
	}
	return loop(1, 0, solver.n)
}
