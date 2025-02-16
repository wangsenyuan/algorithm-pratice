package p3454

import (
	"sort"
)

type event struct {
	y, kind int
	x1, x2  int
}

type Node struct {
	l, r   int
	cover  int
	length float64
}

type SegmentTreeRectangle struct {
	data []int
	n    int
	tree []Node
}

func NewSegmentTreeRectangle(arr []int) *SegmentTreeRectangle {
	n := len(arr) - 1
	tree := make([]Node, n*4+10)
	st := &SegmentTreeRectangle{data: arr, n: n, tree: tree}
	st.build(1, 0, n-1)
	return st
}

func (st *SegmentTreeRectangle) Update(id, l, r, v int) {
	node := &st.tree[id]
	if node.r < l || node.l > r {
		return
	}
	if l <= node.l && node.r <= r {
		node.cover += v
	} else {
		st.Update(id*2, l, r, v)
		st.Update(id*2+1, l, r, v)
	}

	if node.cover > 0 {
		node.length = float64(st.data[node.r+1] - st.data[node.l])
	} else {
		if node.l == node.r {
			node.length = 0
		} else {
			node.length = st.tree[id*2].length + st.tree[id*2+1].length
		}
	}
}

func (st *SegmentTreeRectangle) Query() float64 {
	return st.tree[1].length
}

func (st *SegmentTreeRectangle) build(id, l, r int) {
	st.tree[id].l = l
	st.tree[id].r = r
	if l == r {
		return
	}
	mid := (l + r) / 2
	st.build(id*2, l, mid)
	st.build(id*2+1, mid+1, r)
}

func separateSquares(squares [][]int) float64 {
	events := make([]event, 0, len(squares)*2)
	xSet := map[int]struct{}{}
	for _, sq := range squares {
		x, y, l := sq[0], sq[1], sq[2]
		events = append(events, event{y, 1, x, x + l})
		events = append(events, event{y + l, -1, x, x + l})
		xSet[x] = struct{}{}
		xSet[x+l] = struct{}{}
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].y == events[j].y {
			return events[i].kind < events[j].kind
		}
		return events[i].y < events[j].y
	})

	allX := make([]int, 0, len(xSet))
	for v := range xSet {
		allX = append(allX, v)
	}
	sort.Ints(allX)

	tree := NewSegmentTreeRectangle(allX)
	allSum := 0.0
	preY := events[0].y
	for _, e := range events {
		curY := e.y
		xLen := tree.Query()
		allSum += xLen * float64(curY-preY)
		ql := sort.SearchInts(allX, e.x1)
		qr := sort.SearchInts(allX, e.x2) - 1
		tree.Update(1, ql, qr, e.kind)
		preY = curY
	}

	half := allSum / 2.0
	tree = NewSegmentTreeRectangle(allX)
	preY = events[0].y
	curSum := 0.0
	for _, e := range events {
		curY := e.y
		xLen := tree.Query()
		sum := xLen * float64(curY-preY)
		if curSum+sum >= half {
			diff := half - curSum
			remain := 0.0
			if xLen != 0 {
				remain = diff / xLen
			}
			return float64(preY) + remain
		}
		curSum += sum
		pos1 := sort.SearchInts(allX, e.x1)
		pos2 := sort.SearchInts(allX, e.x2) - 1
		tree.Update(1, pos1, pos2, e.kind)
		preY = curY
	}

	return float64(preY)
}
