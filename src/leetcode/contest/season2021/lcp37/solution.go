package lcp37

import (
	"math"
	"sort"
)

const eps = 1e-8

func minRecSize(lines [][]int) float64 {
	ll := make([]Line, len(lines))

	for i, l := range lines {
		ll[i] = Line{float64(l[0]), float64(l[1])}
	}

	sort.Slice(ll, func(i, j int) bool {
		return ll[i].first < ll[j].first || ll[i].first == ll[j].first && ll[i].second > ll[j].second
	})
	xmin := -get(ll, math.Max)

	sort.Slice(ll, func(i, j int) bool {
		return ll[i].first > ll[j].first || ll[i].first == ll[j].first && ll[i].second > ll[j].second
	})
	xmax := -get(ll, math.Min)

	for i, l := range ll {
		ll[i] = Line{1 / l.first, -l.second / l.first}
	}

	sort.Slice(ll, func(i, j int) bool {
		return ll[i].first < ll[j].first || ll[i].first == ll[j].first && ll[i].second > ll[j].second
	})

	ymin := -get(ll, math.Max)

	sort.Slice(ll, func(i, j int) bool {
		return ll[i].first > ll[j].first || ll[i].first == ll[j].first && ll[i].second > ll[j].second
	})
	ymax := -get(ll, math.Min)

	return (xmax - xmin) * (ymax - ymin)
}

type Line struct {
	first, second float64
}

func get(lines []Line, fn func(float64, float64) float64) float64 {
	n := len(lines)
	var best float64
	var valid bool
	for i := 1; i < n; i++ {
		if math.Abs(lines[i].first-lines[i-1].first) < eps {
			continue
		}
		slop := (lines[i].second - lines[i-1].second) / (lines[i].first - lines[i-1].first)
		if !valid {
			valid = true
			best = slop
		} else {
			best = fn(best, slop)
		}
	}
	if valid {
		return best
	}
	return 0
}
