package p3453

import "math"

func separateSquares(squares [][]int) float64 {

	get := func(limit float64) float64 {
		var sum float64

		for _, cur := range squares {
			_, y, s := cur[0], cur[1], cur[2]
			if float64(y) >= limit {
				continue
			}
			y1 := math.Min(float64(y+s), limit)
			dy := y1 - float64(y)
			sum += float64(s) * dy
		}
		return sum
	}

	sum := get(1 << 50)

	var l, r float64 = 0, 1 << 50

	for range 100 {
		mid := (l + r) / 2
		tmp := get(mid)
		if 2*tmp >= sum {
			r = mid
		} else {
			l = mid
		}
	}
	return (l + r) / 2
}
