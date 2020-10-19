package p1623

import "math"

func bestCoordinate(towers [][]int, radius int) []int {

	check := func(u, v int) int {
		var tot int
		for _, tow := range towers {
			x, y, q := tow[0], tow[1], tow[2]
			d := distance(x, y, u, v)

			if d > float64(radius) {
				continue
			}
			tot += int(float64(q) / (1 + d))
		}
		return int(tot)
	}
	var x0, y0 int
	best := check(x0, y0)

	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			tmp := check(x, y)
			if tmp > best {
				best = tmp
				x0, y0 = x, y
			}
		}
	}

	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			tmp := check(x, y)
			if tmp == best {
				return []int{x, y}
			}
		}
	}
	return nil
}

func distance(x0, y0, x1, y1 int) float64 {
	dx := float64(x1 - x0)
	dy := float64(y1 - y0)

	return math.Sqrt(dx*dx + dy*dy)
}
