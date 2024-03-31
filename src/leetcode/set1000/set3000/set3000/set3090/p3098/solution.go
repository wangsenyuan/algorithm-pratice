package p3098

import (
	"slices"
)

func getLines(points [][]int) []Line {
	n := len(points)

	distance := func(i int, j int) int {
		dx := points[i][0] - points[j][0]
		dy := points[i][1] - points[j][1]
		return abs(dx) + abs(dy)
	}

	if n == 2 {
		return []Line{{0, 1, distance(0, 1)}}
	}

	slices.SortFunc(points, func(a []int, b []int) int {
		if a[0] < b[0] || a[0] == b[0] && a[1] < b[1] {
			return -1
		}
		return 1
	})

	findHull := func(sign int) []int {
		stack := make([]int, n)
		var top int
		stack[top] = 0
		top++
		for i := 1; i < n; i++ {
			for top >= 2 && sign*cross(points[stack[top-2]], points[stack[top-1]], points[i]) >= 0 {
				top--
			}
			stack[top] = i
			top++
		}
		return stack[:top]
	}

	upper := findHull(1)
	lower := findHull(-1)

	outer := make([]int, len(upper)+len(lower)-1)
	copy(outer, upper)
	j := len(upper)
	for i := len(lower) - 1; i > 0; i-- {
		outer[j] = lower[i]
		j++
	}

	var arr []Line
	m := len(outer)
	for l, r := 0, 0; l < m; l++ {
		var cur int
		for {
			cur = distance(outer[r], outer[l])
			nr := (r + 1) % m
			nxt := distance(outer[nr], outer[l])
			if nxt < cur || nxt == 0 {
				break
			}
			r = nr
		}

		arr = append(arr, Line{outer[l], outer[r], cur})
	}

	slices.SortFunc(arr, func(a, b Line) int {
		return b.d - a.d
	})

	return arr
}

func minimumDistance(points [][]int) int {
	n := len(points)
	arr := getLines(points)
	x, y := arr[0].x, arr[0].y
	pts := copyPoints(points)
	if x != n-1 {
		pts[x], pts[n-1] = pts[n-1], pts[x]
	}
	arr = getLines(pts[:n-1])
	ans := arr[0].d
	pts = copyPoints(points)
	if y != n-1 {
		pts[y], pts[n-1] = pts[n-1], pts[y]
	}
	arr = getLines(pts[:n-1])
	ans = min(ans, arr[0].d)
	return ans
}

func copyPoints(ps [][]int) [][]int {
	res := make([][]int, len(ps))
	for i := 0; i < len(ps); i++ {
		res[i] = make([]int, 2)
		copy(res[i], ps[i])
	}
	return res
}

type Line struct {
	x int
	y int
	d int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Point []int

func (this Point) Sub(that Point) Point {
	return Point{this[0] - that[0], this[1] - that[1]}
}

func cross(a, b, c Point) int {
	ab := b.Sub(a)
	bc := c.Sub(b)
	x := ab[0] * bc[1]
	y := ab[1] * bc[0]
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}
