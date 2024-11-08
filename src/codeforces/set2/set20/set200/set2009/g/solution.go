package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k, m := readThreeNums(reader)
		a := readNNums(reader, n)
		queries := make([][]int, m)
		for i := range m {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, k, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type SegTree struct {
	arr  []int
	lazy []int
	sz   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := 0; i < 4*n; i++ {
		lazy[i] = -100
	}
	sz := n
	return &SegTree{arr, lazy, sz}
}

func (tr *SegTree) update(i int, l int, r int, v int) {
	tr.arr[i] = (r - l + 1) * v
	tr.lazy[i] = v
}

func (tr *SegTree) push(i int, l int, r int) {
	if tr.lazy[i] != -100 && l < r {
		mid := (l + r) / 2
		tr.update(2*i+1, l, mid, tr.lazy[i])
		tr.update(2*i+2, mid+1, r, tr.lazy[i])
		tr.lazy[i] = -100
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int, lo int, hi int)

	loop = func(i int, l int, r int, lo int, hi int) {
		tr.push(i, l, r)
		if lo == l && r == hi {
			tr.update(i, l, r, v)
			return
		}
		mid := (l + r) / 2
		if l <= hi && lo <= mid {
			loop(2*i+1, l, mid, lo, min(mid, hi))
		}
		if mid+1 <= hi && lo <= r {
			loop(2*i+2, mid+1, r, max(mid+1, lo), hi)
		}
		tr.arr[i] = tr.arr[2*i+1] + tr.arr[2*i+2]
	}
	loop(0, 0, tr.sz-1, L, R)
}

func (tr *SegTree) Get(L int, R int) int {
	var loop func(i int, l int, r int, lo int, hi int) int
	loop = func(i int, l int, r int, lo int, hi int) int {
		tr.push(i, l, r)

		if lo == l && r == hi {
			return tr.arr[i]
		}
		mid := (l + r) / 2

		var res int

		if lo <= mid && l <= hi {
			res += loop(2*i+1, l, mid, lo, min(hi, mid))
		}
		if mid+1 <= hi && lo <= r {
			res += loop(2*i+2, mid+1, r, max(lo, mid+1), hi)
		}

		return res
	}

	return loop(0, 0, tr.sz-1, L, R)
}

type MultiSet struct {
	cnt []int
	val []int
	sz  int
}

func NewMultiSet(n int) *MultiSet {
	cnt := make([]int, 4*n)
	val := make([]int, 4*n)
	sz := n
	return &MultiSet{cnt, val, sz}
}

func (tr *MultiSet) Add(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i] += v
			if tr.cnt[i] > 0 {
				tr.val[i] = l
			} else {
				tr.val[i] = 0
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = max(tr.val[2*i+1], tr.val[2*i+2])
	}

	loop(0, 0, tr.sz-1)
}

func solve(a []int, k int, queries [][]int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] -= i
	}

	set := NewMultiSet(n + 1)
	freq := make(map[int]int)

	add := func(num int) {
		if freq[num] > 0 {
			set.Add(freq[num], -1)
		}
		freq[num]++
		set.Add(freq[num], 1)
	}

	rem := func(num int) {
		set.Add(freq[num], -1)
		freq[num]--
		if freq[num] > 0 {
			set.Add(freq[num], 1)
		}
	}

	mn := make([]int, n)

	for i := 0; i < n; i++ {
		if i >= k {
			rem(a[i-k])
		}
		add(a[i])
		if i >= k-1 {
			mn[i-k+1] = k - set.val[0]
		}
	}

	type question struct {
		id int
		l  int
		r  int
	}

	qs := make([]question, len(queries))

	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		qs[i] = question{i, l, r}
	}

	slices.SortFunc(qs, func(x, y question) int {
		if x.l != y.l {
			return y.l - x.l
		}
		return x.r - y.r
	})

	pos := n - k

	type pair struct {
		first  int
		second int
	}
	stack := make([]pair, n)

	stack[0] = pair{mn[n-k], n - k}
	top := 1

	tr := NewSegTree(n)
	tr.Update(n-k, n-k, mn[n-k])

	res := make([]int, len(qs))

	for _, cur := range qs {
		for pos != cur.l {
			pos--
			next := pos
			for top > 0 && stack[top-1].first > mn[pos] {
				next = stack[top-1].second
				top--
			}
			stack[top] = pair{mn[pos], next}
			top++
			tr.Update(pos, next, mn[pos])
		}
		res[cur.id] = tr.Get(cur.l, cur.r-k+1)
	}

	return res
}

func solve1(a []int, k int, queries [][]int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] += n - 2 - i
	}
	mn := make([]int, n-k+1)
	cnt := make([]int, n*2-1)
	h := lazyHeap{todo: make([]int, n+1)}
	for r, v := range a {
		if cnt[v] > 0 {
			h.del(cnt[v])
		}
		cnt[v]++
		h.push(cnt[v])

		l := r + 1 - k
		if l < 0 {
			continue
		}
		mn[l] = k - h.top()

		v = a[l]
		h.del(cnt[v])
		cnt[v]--
		if cnt[v] > 0 {
			h.push(cnt[v])
		}
	}

	type pair struct{ r, i int }
	qs := make([][]pair, n-k+1)
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		qs[l-1] = append(qs[l-1], pair{r - k, i})
	}

	ans := make([]int, len(queries))
	type data struct{ r, v, s int }
	st := []data{{len(mn), -1, 0}}
	for i := len(mn) - 1; i >= 0; i-- {
		v := mn[i]
		r := i
		for st[len(st)-1].v >= v {
			r = st[len(st)-1].r
			st = st[:len(st)-1]
		}
		st = append(st, data{r, v, st[len(st)-1].s + (r-i+1)*v})
		for _, p := range qs[i] {
			j := sort.Search(len(st), func(i int) bool { return st[i].r < p.r }) - 1
			ans[p.i] = st[len(st)-1].s - st[j-1].s - (st[j].r-p.r)*st[j].v
		}
	}

	return ans
}

type lazyHeap struct {
	sort.IntSlice
	todo []int
}

func (h lazyHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *lazyHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *lazyHeap) del(v int)         { h.todo[v]++ }
func (h *lazyHeap) push(v int) {
	if h.todo[v] > 0 {
		h.todo[v]--
	} else {
		heap.Push(h, v)
	}
}
func (h *lazyHeap) top() int {
	for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
		h.todo[h.IntSlice[0]]--
		heap.Pop(h)
	}
	return h.IntSlice[0]
}
