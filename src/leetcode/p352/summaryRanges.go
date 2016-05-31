package p352

import "fmt"

//SummaryRanges is used to get the disjoint set
type SummaryRanges struct {
	intervals []Interval
}

// Constructor is used Initialize your data structure here
func Constructor() SummaryRanges {
	return SummaryRanges{intervals: make([]Interval, 0, 0)}
}

//Addnum add one more element
func (sr *SummaryRanges) Addnum(val int) {
	i, j := 0, len(sr.intervals)-1
	//fmt.Printf("will process %d ==> %v\n", val, sr.intervals)
	for i <= j {
		mid := (i + j) / 2
		a := sr.intervals[mid]
		if val+1 < a.Start {
			j = mid - 1
		} else if val > a.End+1 {
			i = mid + 1
		} else {
			if val >= a.Start && val <= a.End {
				return
			}

			if val+1 == a.Start {
				a.Start = val
				sr.intervals = mergeInterval(sr.intervals, mid-1, mid)
				return
			}

			if val == a.End+1 {
				a.End = val
				sr.intervals = mergeInterval(sr.intervals, mid, mid+1)
				return
			}
		}
	}
	b := Interval{Start: val, End: val}

	//fmt.Printf("after %d is %v\n", j+1, sr.intervals[j+1:])

	sr.intervals = insertIntoSlice(sr.intervals, b, j+1)
}

func insertIntoSlice(slice []Interval, x Interval, j int) []Interval {
	//fmt.Printf("will insert %v into %v at %d\n", x, slice, j)
	xs := slice[0:j]
	//fmt.Printf("before %d is %v\n", j, xs)
	ys := slice[j:]
	//fmt.Printf("after %d is %v\n", j, ys)
	zs := make([]Interval, len(xs), len(slice)+1)
	copy(zs, xs)

	zs = append(zs, x)
	zs = append(zs, ys...)

	return zs
}

func mergeInterval(intervals []Interval, i, j int) []Interval {
	if i < 0 || j > len(intervals) {
		return intervals
	}

	a, b := intervals[i], intervals[j]

	if a.End+1 < b.Start {
		return intervals
	}

	c := Interval{Start: a.Start, End: b.End}
	xs := intervals[0:i]
	xs = append(xs, c)
	return append(xs, intervals[j+1:]...)
}

//Getintervals get the interverl result
func (sr *SummaryRanges) Getintervals() []Interval {
	return sr.intervals
}

//Interval is from leetcode
type Interval struct {
	Start int
	End   int
}

func (x Interval) String() string {
	return fmt.Sprintf("[%d, %d]", x.Start, x.End)
}
