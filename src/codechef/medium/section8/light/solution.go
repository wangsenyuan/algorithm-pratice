package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var N int
	var L, R float64
	fmt.Scanf("%d %f %f", &N, &L, &R)
	lights := make([][]float64, N)
	for i := 0; i < N; i++ {
		lights[i] = make([]float64, 3)
		fmt.Scanf("%f %f %f", &lights[i][0], &lights[i][1], &lights[i][2])
	}

	res := solve(N, L, R, lights)

	fmt.Printf("%.10f\n", res)
}

func solve(N int, L, R float64, lights [][]float64) float64 {

	arr := make([]Pair, 2*N+2)

	check := func(h float64) bool {
		var n int

		for i := 0; i < N; i++ {
			x, y, z := lights[i][0], lights[i][1], lights[i][2]
			z = z * math.Pi / 180
			if y < h {
				continue
			}
			dx := (y - h) * math.Tan(z)
			arr[n] = Pair{x - dx, 1}
			n++
			arr[n] = Pair{x + dx, -1}
			n++
		}
		arr[n] = Pair{L, 0}
		n++
		arr[n] = Pair{R, 0}
		n++
		sort.Sort(Pairs(arr[:n]))
		var cnt int
		var la float64 = float64(math.MinInt32)

		for i := 0; i < n; i++ {
			x := arr[i].first
			if x > L && x <= R && la < x && cnt <= 0 {
				return false
			}
			la = x
			cnt += arr[i].second
		}
		return true
	}

	var left float64
	var right float64 = 1001

	for i := 0; i < 64; i++ {
		mid := (left + right) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return (left + right) / 2
}

type Pair struct {
	first  float64
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
