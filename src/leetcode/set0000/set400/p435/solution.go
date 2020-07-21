package p435

import "sort"

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(ByEnd(intervals))

	keep := 1
	last := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if last.End <= intervals[i].Start {
			keep++
			last = intervals[i]
		}
	}
	return len(intervals) - keep
}

/**
* sort interval by end
 */
type ByEnd []Interval

//length of the array
func (a ByEnd) Len() int {
	return len(a)
}

func (a ByEnd) Less(i, j int) bool {
	return a[i].End < a[j].End
}

func (a ByEnd) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
