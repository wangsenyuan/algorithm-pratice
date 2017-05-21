package main

import "fmt"

func main() {

	/*points := []Point{
		Point{1, 2}, Point{2, 2}, Point{4, 2},
	}*/

	points := []Point{
		Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{1, 2},
		Point{2, 2}, Point{3, 2}, Point{3, 1}, Point{3, 0},
		Point{2, 0}, Point{1, 0}, Point{1, 1}, Point{3, 3},
	}
	/*points := []Point{
		Point{1, 1}, Point{2, 2}, Point{2, 0}, Point{2, 4}, Point{3, 3}, Point{4, 2},
	}*/
	res := outerTrees(points)
	fmt.Println(res)
}

/**
 * Definition for a point.
  */
type Point struct {
	X int
	Y int
}

func outerTrees(points []Point) []Point {
	var first = 0
	for i, point := range points {
		if points[first].X > point.X {
			first = i
		}
	}

	answer := make(map[int]bool)
	answer[first] = true
	cur := first
	for {
		next := 0
		for i := 1; i < len(points); i++ {
			if i == cur {
				continue
			}
			cross := crossProduct(points[i], points[cur], points[next])

			if next == cur || cross > 0 || (cross == 0 && dist(points[i], points[cur]) > dist(points[next], points[cur])) {
				next = i
			}
		}
		for i := 0; i < len(points); i++ {
			cross := crossProduct(points[i], points[cur], points[next])
			if i != cur && cross == 0 {
				answer[i] = true
			}
		}
		cur = next
		if cur == first {
			break
		}
	}

	res := make([]Point, len(answer))
	j := 0
	for i, _ := range answer {
		res[j] = points[i]
		j++
	}
	return res
}

func dist(a Point, b Point) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}
func crossProduct(a Point, b Point, c Point) int {
	bax := a.X - b.X
	bay := a.Y - b.Y
	bcx := c.X - b.X
	bcy := c.Y - b.Y
	return bax*bcy - bcx*bay
}
