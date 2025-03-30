package p3501

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)
	var sum int
	for i := range n {
		sum += int(s[i] - '0')
	}

	type pair struct {
		first  int
		second int
	}
	at := make([][]pair, n)
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		at[r] = append(at[r], pair{l, i})
	}
	ans := make([]int, len(queries))
	tr := NewSegTree(n)

	stack := make([]pair, n)
	var top int
	for i := 0; i < n; i++ {
		if i == 0 || s[i] != s[i-1] {
			stack[top] = pair{i, i}
			top++
		} else {
			stack[top-1].second = i
		}
		if s[i] == '0' {
			if top >= 3 {
				tr.Update(stack[top-3].first, i, 1)
			} else {
				tr.Update(stack[top-1].first, i, 1)
			}
		}
		for _, cur := range at[i] {
			l, id := cur.first, cur.second
			ans[id] = sum
			if top == 1 || l >= stack[top-2].first || top >= 3 && l >= stack[top-3].first && s[i] == '1' {
				// l和i在相邻的两个区间内
				continue
			}
			ans[id] += tr.Get(l, stack[top-1].first-1)
		}
	}

	return ans
}

type SegTree struct {
	arr  []int
	lazy []int
	sz   int
}

func NewSegTree(n int) *SegTree {
	return &SegTree{
		arr:  make([]int, n*4),
		lazy: make([]int, n*4),
		sz:   n,
	}
}

func (tr *SegTree) update(i int, v int) {
	tr.arr[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(i*2+1, tr.lazy[i])
		tr.update(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
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
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		tr.arr[i] = max(tr.arr[i*2+1], tr.arr[i*2+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		return max(loop(i*2+1, l, mid), loop(i*2+2, mid+1, r))
	}
	return loop(0, 0, tr.sz-1)
}
