package p2280

import (
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	n := len(stockPrices)

	if n == 1 {
		return 0
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})


	P := make([]Point, n-1)

	for i := 1; i < n; i++ {
		P[i-1] = NewPoint(stockPrices[i]).Sub(NewPoint(stockPrices[i-1]))
	}

	res := 1

	for i := 1; i < n-1; i++ {
		if P[i-1].Cross(P[i]) != 0 {
			res++
		}
	}

	return res
}

type Point struct {
	x int64
	y int64
}

func NewPoint(coord []int) Point {
	x, y := coord[0], coord[1]
	return Point{int64(x), int64(y)}
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y}
}

func (this Point) Cross(that Point) int64 {
	return this.x*that.y - this.y*that.x
}
