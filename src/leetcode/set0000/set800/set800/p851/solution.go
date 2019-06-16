package p851

import "sort"

const MOD = 1000000007

func rectangleArea(rectangles [][]int) int {
	n := len(rectangles)
	events := make([]Event, 2*n)

	ys := make([]int, 2*n)
	for i := 0; i < n; i++ {
		x0, y0, x1, y1 := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		//enter
		events[2*i] = Event{x0, y0, y1, 1}
		//exit
		events[2*i+1] = Event{x1, y0, y1, -1}
		ys[2*i] = y0
		ys[2*i+1] = y1
	}

	sort.Sort(Events(events))
	sort.Ints(ys)

	// keep unique y
	var yl int
	for i := 1; i <= 2*n; i++ {
		if i < 2*n && ys[i] == ys[i-1] {
			continue
		}
		ys[yl] = ys[i-1]
		yl++
	}

	yi := make(map[int]int)
	for i := 0; i < yl; i++ {
		yi[ys[i]] = i
	}
	st := new(SegTree)
	st.ys = ys[:yl]
	st.start = 0
	st.end = yl - 1

	var ans uint64
	var cur_y_sum uint64
	var cur_x = uint64(events[0].x)

	for _, event := range events {
		x, tpe, y0, y1 := event.x, event.tpe, event.y0, event.y1
		ans += (cur_y_sum * (uint64(x) - cur_x)) % MOD
		ans %= MOD
		cur_y_sum = uint64(st.Update(yi[y0], yi[y1], tpe))
		cur_x = uint64(x)
	}

	return int(ans)
}

type Event struct {
	x, y0, y1 int
	tpe       int
}

type Events []Event

func (this Events) Len() int {
	return len(this)
}

func (this Events) Less(i, j int) bool {
	return this[i].x < this[j].x
}

func (this Events) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type SegTree struct {
	ys           []int
	start, end   int
	count, total int
	left, right  *SegTree
}

func (st *SegTree) getRangeMid() int {
	return (st.start + st.end) >> 1
}

func (st *SegTree) getLeft() *SegTree {
	if st.left == nil {
		left := new(SegTree)
		left.start = st.start
		left.end = st.getRangeMid()
		left.ys = st.ys
		st.left = left
	}
	return st.left
}

func (st *SegTree) getRight() *SegTree {
	if st.right == nil {
		right := new(SegTree)
		right.start = st.getRangeMid()
		right.end = st.end
		right.ys = st.ys
		st.right = right
	}
	return st.right
}

func (st *SegTree) Update(bot, top int, val int) int {
	var loop func(root *SegTree, start, end int) int
	loop = func(root *SegTree, start, end int) int {
		if end <= start {
			return 0
		}
		if root.start == start && root.end == end {
			root.count += val
		} else {
			loop(root.getLeft(), start, min(root.getRangeMid(), end))
			loop(root.getRight(), max(root.getRangeMid(), start), end)
		}

		var total int
		if root.count > 0 {
			total = root.ys[root.end] - root.ys[root.start]
		} else {
			total = root.getLeft().total + root.getRight().total
		}
		root.total = total

		return total;
	}

	return loop(st, bot, top)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rectangleArea1(rectangles [][]int) int {
	n := len(rectangles)
	xs := make([]int, 2*n)
	for i, j := 0, 0; i < n; i++ {
		rect := rectangles[i]
		xs[j] = rect[0]
		j++
		xs[j] = rect[2]
		j++
	}

	sort.Ints(xs)

	rs := make(Rects, n)
	for i := 0; i < n; i++ {
		rs[i] = Rect{rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]}
	}

	sort.Sort(rs)

	overlap := func(r Rect, left, right int) bool {
		return r.x0 < right && left < r.x1
	}

	area := func(left, right int) uint64 {
		j := -1
		for i := 0; i < n; i++ {
			r := rs[i]
			if overlap(r, left, right) {
				j = i
				break
			}
		}
		// j is the first overlapped rect
		if j < 0 {
			return 0
		}

		y0 := rs[j].y0
		y1 := rs[j].y1

		var res uint64
		for i := j + 1; i < n; i++ {
			r := rs[i]
			if !overlap(r, left, right) {
				continue
			}
			if r.y0 <= y1 {
				y1 = max(y1, r.y1)
				continue
			}
			res += uint64(y1 - y0)
			res %= MOD
			y0 = r.y0
			y1 = r.y1
		}
		res += uint64(y1 - y0)
		res %= MOD
		return res
	}

	var ans uint64

	for i := 1; i < 2*n; i++ {
		left, right := xs[i-1], xs[i]
		ans += (area(left, right) * uint64(right-left)) % MOD
		ans %= MOD
	}

	return int(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Rect struct {
	x0, y0, x1, y1 int
}

type Rects []Rect

func (rs Rects) Len() int {
	return len(rs)
}

func (rs Rects) Less(i, j int) bool {
	// from bottom to top
	return rs[i].y0 < rs[j].y0 || (rs[i].y0 == rs[j].y0 && rs[i].y1 < rs[j].y1)
}

func (rs Rects) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
