package main

import "sort"

//Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Sort(byRange(intervals))

	last := intervals[0]
	result := make([]Interval, 0, len(intervals))
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval.Start > last.End {
			result = append(result, last)
			last = interval
			continue
		}

		if interval.End > last.End {
			last.End = interval.End
		}
	}
	result = append(result, last)

	return result
}

type byRange []Interval

func (a byRange) Len() int {
	return len(a)
}

func (a byRange) Less(i, j int) bool {
	return a[i].Start < a[j].Start || (a[i].Start == a[j].Start && a[i].End < a[j].End)
}

func (a byRange) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
