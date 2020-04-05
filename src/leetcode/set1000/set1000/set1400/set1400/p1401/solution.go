package p1401

import "math"

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {

	R := float64(radius)

	if distance(x_center, y_center, x1, y1) <= R {
		return true
	}

	if distance(x_center, y_center, x2, y2) <= R {
		return true
	}

	if distance(x_center, y_center, x1, y2) <= R {
		return true
	}

	if distance(x_center, y_center, x2, y1) <= R {
		return true
	}

	if x1 <= x_center && x_center <= x2 && y1 <= y_center && y_center <= y2 {
		return true
	}

	if y1 <= y_center && y_center <= y2 {
		x0 := float64(x1+x2) / 2
		dx := math.Abs(x0 - float64(x_center))
		if x_center <= x1 && dx <= R+x0-float64(x1) {
			return true
		}
		if x_center >= x2 && dx <= R+float64(x2)-x0 {
			return true
		}
	}

	if x1 <= x_center && x_center <= x2 {
		y0 := float64(y1+y2) / 2
		dy := math.Abs(y0 - float64(y_center))

		if y_center <= y1 && dy <= R+y0-float64(y1) {
			return true
		}

		if y_center >= y2 && dy <= R+float64(y2)-y0 {
			return true
		}
	}

	return false
}

func distance(a, b, c, d int) float64 {
	dx := float64(c - a)
	dy := float64(d - b)

	return math.Sqrt(dx*dx + dy*dy)

}
