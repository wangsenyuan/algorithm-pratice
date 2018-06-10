package p851

import "sort"

const MOD = 1000000007

func rectangleArea(rectangles [][]int) int {
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
