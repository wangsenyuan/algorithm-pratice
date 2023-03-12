package p2586

import "sort"

const MAX_X = 2010

func findMinimumTime(tasks [][]int) int {
	n := len(tasks)
	ts := make([]*Task, n)

	for i := 0; i < n; i++ {
		ts[i] = &Task{tasks[i][0], tasks[i][1], tasks[i][2]}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Less(ts[j])
	})

	t := NewTree(MAX_X)
	var res int
	for i := 0; i < n; i++ {
		u := ts[i]
		a := u.start
		b := u.end
		x := t.Get(a, b)
		tmp := u.duration - min(x, u.duration)
		res += tmp
		for j := b; tmp > 0; j-- {
			if t.Get(j, j) == 0 {
				t.Set(j)
				tmp--
			}
		}
	}

	return res
}

type Tree struct {
	val []int
	sz  int
}

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	return &Tree{arr, n}
}

func (t *Tree) Set(p int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.val[i] = 1
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i, l, mid)
		} else {
			loop(2*i+1, mid+1, r)
		}
		t.val[i] = t.val[2*i] + t.val[2*i+1]
	}
	loop(1, 0, t.sz-1)
}

func (t *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return t.val[i]
		}
		mid := (l + r) / 2
		a := loop(2*i, l, mid)
		b := loop(2*i+1, mid+1, r)
		return a + b
	}

	return loop(1, 0, t.sz-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Task struct {
	start    int
	end      int
	duration int
}

func (this *Task) Less(that *Task) bool {
	return this.end < that.end
}
