package p1232

import "sort"

func checkStraightLine(coordinates [][]int) bool {
	sort.Sort(Pairs(coordinates))

	a, b := coordinates[0][0], coordinates[0][1]
	c, d := coordinates[1][0], coordinates[1][1]

	g := gcd(c-a, d-b)

	dx := (c - a) / g
	dy := (d - b) / g

	for i := 2; i < len(coordinates); i++ {
		du := coordinates[i][0] - coordinates[i-1][0]
		dv := coordinates[i][1] - coordinates[i-1][1]
		g = gcd(du, dv)
		du /= g
		dv /= g
		if du != dx || dv != dy {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pairs [][]int

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i][0] < this[j][0] || (this[i][0] == this[j][0] && this[i][1] < this[j][1])
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
