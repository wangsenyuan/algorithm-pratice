package main

import "fmt"

func main() {
	tree := []int{2, 2}
	squirrel := []int{4, 4}
	nuts := [][]int{{3, 0}, {2, 5}}
	fmt.Println(minDistance(5, 7, tree, squirrel, nuts))
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	dd := width + height
	dist := 0

	a, b := tree[0], tree[1]
	c, d := squirrel[0], squirrel[1]

	j := -1
	for i := 0; i < len(nuts); i++ {
		e, f := nuts[i][0], nuts[i][1]
		d1 := distance(a, b, e, f)
		d2 := distance(c, d, e, f)

		if d2-d1 < dd {
			dd = d2 - d1
			j = i
		}

		dist += 2 * d1
	}

	e, f := nuts[j][0], nuts[j][1]
	dist = dist - distance(a, b, e, f) + distance(c, d, e, f)
	return dist
}

func distance(a, b, c, d int) int {
	return abs(c-a) + abs(d-b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
