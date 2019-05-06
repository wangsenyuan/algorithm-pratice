package p1037

import "sort"

func isBoomerang(points [][]int) bool {
	sort.Sort(PS(points))
	a := points[0]
	b := points[1]
	c := points[2]

	if samePoint(a, b) || samePoint(b, c) {
		return false
	}

	dx := b[0] - a[0]
	dy := b[1] - a[1]

	du := c[0] - b[0]
	dv := c[1] - b[1]

	// dy / dx == dv / du
	if dx == 0 {
		return du != 0
	}

	return dy*du != dv*dx
}

func samePoint(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

type PS [][]int

func (this PS) Len() int {
	return len(this)
}

func (this PS) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this PS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
