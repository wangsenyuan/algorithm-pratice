package p2250

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {
	all_x := make(map[int]int)
	for _, rec := range rectangles {
		all_x[rec[0]]++
	}

	for _, point := range points {
		all_x[point[0]]++
	}

	ii := make([]int, 0, len(all_x))

	for x := range all_x {
		ii = append(ii, x)
	}

	sort.Ints(ii)

	for i := 0; i < len(ii); i++ {
		all_x[ii[i]] = i
	}

	n := len(all_x)

	arr := make([]int, n+1)

	set := func(p int) {
		p++
		for p <= n {
			arr[p]++
			p += p & -p
		}
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += arr[p]
			p -= p & -p
		}
		return res
	}

	P := make([]Point, len(rectangles)+len(points))

	for i := 0; i < len(rectangles); i++ {
		P[i] = Point{rectangles[i][0], rectangles[i][1], -1}
	}

	for i := 0; i < len(points); i++ {
		P[i+len(rectangles)] = Point{points[i][0], points[i][1], i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].Less(P[j])
	})

	ans := make([]int, len(points))

	for i := 0; i < len(P); i++ {
		p := P[i]
		x := all_x[p.x]
		if p.index >= 0 {
			// already sorted by Y, only need to check X
			ans[p.index] = get(n-1) - get(x-1)
		} else {
			set(x)
		}
	}
	return ans
}

type Point struct {
	x, y  int
	index int
}

func (this Point) Less(that Point) bool {
	return this.y > that.y || this.y == that.y && this.index < that.index
}
