package p973

import "sort"

func kClosest(points [][]int, K int) [][]int {
	n := len(points)
	ps := make([]P, n)
	for i := 0; i < n; i++ {
		p := P{points[i][0], points[i][1], distance(points[i][0], points[i][1])}
		ps[i] = p
	}
	sort.Sort(PS(ps))

	res := make([][]int, K)
	for i := 0; i < K && i < n; i++ {
		res[i] = []int{ps[i].x, ps[i].y}
	}
	return res
}

func distance(a, b int) int64 {
	A := int64(a)
	B := int64(b)
	return A*A + B*B
}

type P struct {
	x    int
	y    int
	dist int64
}

type PS []P

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].dist < ps[j].dist
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
