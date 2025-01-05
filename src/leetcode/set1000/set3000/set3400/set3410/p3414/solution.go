package p3414

import "sort"

func maximumWeight(intervals [][]int) []int {
	pos := make(map[int]int)
	for _, cur := range intervals {
		pos[cur[0]]++
		pos[cur[1]]++
	}
	var arr []int
	for k := range pos {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i, x := range arr {
		pos[x] = i
	}
	m := len(arr)
	tr := NewTree(m)

	n := len(intervals)
	fp := make([]pair, m)
	from := make([][]pair, 4)

	for d := 0; d < 4; d++ {
		for i := range m {
			fp[i] = pair{0, -1}
		}
		from[d] = make([]pair, n)
		for i, cur := range intervals {
			l, r, w := cur[0], cur[1], cur[2]
			l = pos[l]
			r = pos[r]
			tmp := tr.Get(0, l)
			w += tmp.first
			from[d][i] = pair{w, tmp.second}
			fp[r] = max_of(fp[r], pair{w, i})
		}

		for i := 0; i < m; i++ {
			tr.Update(i, fp[i])
		}
	}

	get := func(u int) (res int, tmp []int) {
		d := 3
		res = from[3][u].first
		for u >= 0 {
			tmp = append(tmp, u)
			u = from[d][u].second
			d--
		}
		sort.Ints(tmp)
		return
	}

	var ans []int
	var best int
	for i := 0; i < n; i++ {
		cur, tmp := get(i)
		if cur < best {
			continue
		}
		if best < cur {
			ans = tmp
		} else {
			ans = comp(ans, tmp)
		}
		best = cur
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func comp(a, b []int) []int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return a
		} else if a[i] > b[i] {
			return b
		}
	}
	if len(a) < len(b) {
		return a
	}
	return b
}

type pair struct {
	first  int
	second int
}

func max_of(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

type Tree struct {
	arr []pair
	sz  int
}

func NewTree(n int) *Tree {
	arr := make([]pair, 2*n)
	for i := 0; i < 2*n; i++ {
		arr[i] = pair{0, -1}
	}
	return &Tree{arr, n}
}

func (tr *Tree) Update(p int, v pair) {
	p += tr.sz
	tr.arr[p] = v
	for p > 1 {
		tr.arr[p>>1] = max_of(tr.arr[p], tr.arr[p^1])
		p >>= 1
	}
}

func (tr *Tree) Get(l int, r int) pair {
	l += tr.sz
	r += tr.sz
	res := pair{0, -1}
	for l < r {
		if l&1 == 1 {
			res = max_of(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_of(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
