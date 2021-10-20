package lcp42

import "sort"

func circleGame(toys [][]int, circles [][]int, r int) int {

	sort.Slice(circles, func(i, j int) bool {
		return circles[i][0] < circles[j][0] || circles[i][0] == circles[j][0] && circles[i][1] < circles[j][1]
	})

	type Line struct {
		x  int
		ys []int
	}

	var xs []Line

	var prev_y = -1
	for i := 0; i < len(circles); i++ {
		if len(xs) == 0 || circles[i][0] > xs[len(xs)-1].x {
			x := circles[i][0]
			var ys []int
			xs = append(xs, Line{x, ys})
			prev_y = -1
		}

		if circles[i][1] > prev_y {
			j := len(xs)
			xs[j-1].ys = append(xs[j-1].ys, circles[i][1])
			prev_y = circles[i][1]
		}
	}

	var res int

	for _, toy := range toys {
		x, y, rt := toy[0], toy[1], toy[2]
		if rt > r {
			continue
		}
		x0 := x + rt - r
		i := sort.Search(len(xs), func(i int) bool {
			return xs[i].x >= x0
		})
		for i < len(xs) && xs[i].x-r <= x-rt {
			ys := xs[i].ys
			j := sort.Search(len(ys), func(j int) bool {
				return ys[j] >= y
			})
			if j < len(ys) {
				cx, cy := xs[i].x, ys[j]
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= (r-rt)*(r-rt) {
					res++
					break
				}
			}

			if j > 0 {
				// ys[j-1] - r <= y - rt
				j--
				cx, cy := xs[i].x, ys[j]
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= (r-rt)*(r-rt) {
					res++
					break
				}
			}
			i++
		}
	}

	return res
}
